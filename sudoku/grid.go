package sudoku

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const gridSize int = 9

type collIdx int

const (
	collRow collIdx = 0
	collCol collIdx = 1
	collBlk collIdx = 2
)

// gridCoord allows us to use the following code:
//  for ri, rn := range gridCoord
// this allows us to use ri for a zero-based value (0 to 8)
// or rn for a 1-based value (1 to 9) as required.
//
// I actually started using 1 to 9, but then halfway through switched
// to 0 to 8. At this point, unsure which I need, but this seemed
// a neet way to allow either until I'm sure we don't need rn/cn
var gridCoord = [gridSize]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Grid is the entire game board
type Grid struct {
	grid       [gridSize][gridSize]*cell
	collection [3][gridSize]cellCollection
	emptyCells int
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := newCell(g)
			g.grid[ri][ci] = c
			bi := calcBlkIdx(ri, ci)
			for vi := val1 - 1; vi < val9; vi++ {
				g.collection[collRow][ri][vi] = c
				g.collection[collCol][ci][vi] = c
				g.collection[collBlk][bi][vi] = c
			}
		}
	}
	return g
}

func calcBlkIdx(ri, ci int) int {
	r := ri / 3
	c := ci / 3
	return c*3 + r
}

// Import from specified .sp file
func (g *Grid) Import(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ri := 0
	for scanner.Scan() {
		line := scanner.Text()
		if ok, _ := regexp.MatchString(`^[*1-9]{9}$`, line); ok {
			r := []rune(line)
			for ci := range gridCoord {
				rv := r[ci]
				if rv == '*' {
					continue
				}
				vi, _ := strconv.Atoi(string(rv))
				err := g.SetValue(ri, ci, value(vi))
				if err != nil {
					return err
				}
			}
			ri++
		}
	}
	return nil
}

// SetValue sets the value of the specified cell. This includes
// recalculating all valid possible values appropriately
func (g *Grid) SetValue(ri, ci int, v value) error {
	fmt.Printf("Setting (%d,%d) to %d\n", ri, ci, v)
	c := g.grid[ri][ci]
	if c.canSet(v) {
		c.setValue(v)
		bi := calcBlkIdx(ri, ci)
		g.collection[collRow][ri].notPossible(v)
		g.collection[collCol][ci].notPossible(v)
		g.collection[collBlk][bi].notPossible(v)
		return nil
	}
	return fmt.Errorf("cannot set cell (%d,%d) to %d", ri, ci, v)
}

// Display handles the grid output
func (g *Grid) Display() {
	for ri := range gridCoord {
		for ci := range gridCoord {
			fmt.Printf("[ %s ]", g.grid[ri][ci].val) // Stringify to return " " for empty
		}
		fmt.Println()
	}
}

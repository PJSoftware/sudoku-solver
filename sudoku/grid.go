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
	collRow collIdx = iota
	collCol
	collBlk
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
	cells      [gridSize][gridSize]*cell
	collection [3][gridSize]cellCollection
	emptyCells int
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := newCell()
			g.emptyCells++
			g.cells[ri][ci] = c
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
	c := g.cells[ri][ci]
	if c.canSet(v) {
		c.setValue(v)
		g.emptyCells--
		c.status = original
		g.updateCollections(ri, ci, v)
		return nil
	}
	return fmt.Errorf("cannot set cell (%d,%d) to %d", ri, ci, v)
}

func (g *Grid) updateCollections(ri, ci int, v value) {
	bi := calcBlkIdx(ri, ci)
	fmt.Printf("Setting row %d, col %d, block %d value %d\n", ri, ci, bi, v)
	g.collection[collRow][ri].notPossible(v)
	g.collection[collCol][ci].notPossible(v)
	g.collection[collBlk][bi].notPossible(v)
}

// Display handles the grid output
func (g *Grid) Display() {
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.cells[ri][ci]
			fmt.Printf("[ %s ]", c.val) // Stringify to return " " for empty
		}
		fmt.Println()
	}
}

// Solve the grid
func (g *Grid) Solve() {
	for g.emptyCells > 0 {
		fmt.Println("Solver running:")
		ec := g.emptyCells
		g.solveFirstPass()
		if ec == g.emptyCells { // if emptyCells has not changed our solver is stuck
			fmt.Printf("Solver is stuck with %d empty cells remaining\n", ec)
			return
		}
	}
}

func (g *Grid) solveFirstPass() {
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.cells[ri][ci]
			fmt.Printf("Cell (%d,%d) [%s] has %d possible values: %v\n", ri, ci, c.val, c.pCount, c.possible)
			if c.pCount == 1 {
				for v, p := range c.possible {
					if p {
						fmt.Printf("S01: Setting (%d,%d) to %d\n", ri, ci, v)
						c.setValue(v)
						g.emptyCells--
						g.updateCollections(ri, ci, v)
					}
				}
			}
		}
	}
}

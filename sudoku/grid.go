package sudoku

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const gridSize int = 9

// gridCoord allows us to use the following code:
//  for ri, rn := range gridCoord
// this allows us to use ri for a zero-based value (0 to 8)
// or rn for a 1-based value (1 to 9) as required.
var gridCoord = [gridSize]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Grid is the entire game board
type Grid struct {
	cell [gridSize][gridSize]*cell
	cc   *cellCollections
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	g.cc = new(cellCollections)
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := newCell(ri, ci, g)
			g.cell[ri][ci] = c
		}
	}
	return g
}

func (g *Grid) returnCell(ri, ci int) *cell {
	return g.cell[ri][ci]
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
				vr, err := strconv.Atoi(string(rv))
				if err != nil {
					return fmt.Errorf("Error converting '%s' to value: %v", string(rv), err)
				}
				c := g.returnCell(ri, ci)
				c.val = value(vr)
			}
			ri++
		}
	}
	return nil
}

// Display handles the grid output
func (g *Grid) Display() {
	for ri := range gridCoord {
		if ri%3 == 0 {
			drawHoriz()
		}
		for ci := range gridCoord {
			if ci%3 == 0 {
				drawVert()
			}
			c := g.returnCell(ri, ci)
			fmt.Printf(" %s ", c.val)
		}
		drawEndOfRow()
	}
	drawHoriz()
}

func (g *Grid) emptyCells() int {
	ecc := 0
	for ri := range gridCoord {
		for ci := range gridCoord {
			if g.cell[ri][ci].val == empty {
				ecc++
			}
		}
	}
	return ecc
}

func drawHoriz() {
	fmt.Println("+---------+---------+---------+")
}

func drawVert() {
	fmt.Print("|")
}

func drawEndOfRow() {
	drawVert()
	fmt.Println()
}

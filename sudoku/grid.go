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
	cells      [gridSize][gridSize]*cell
	emptyCells int
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := newCell(ri, ci)
			g.cells[ri][ci] = c
			g.emptyCells++
		}
	}
	return g
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
	if c.setValue(v) {
		g.emptyCells--
		return nil
	}
	return fmt.Errorf("cannot set cell (%d,%d) to %d", ri, ci, v)
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
			c := g.cells[ri][ci]
			fmt.Printf(" %s ", c.val) // Stringify to return " " for empty
		}
		drawEndofRow()
	}
	drawHoriz()
}

// Solve the grid
func (g *Grid) Solve() {
	pass := 1
	for g.emptyCells > 0 {
		fmt.Printf("Solver running; pass %d: ", pass)
		numSolved := g.solveFirstPass()
		fmt.Printf("%d cells solved\n", numSolved)
		if numSolved == 0 {
			fmt.Printf("Solver is stuck with %d empty cells remaining\n", g.emptyCells)
			return
		}
		pass++
	}
}

func (g *Grid) solveFirstPass() int {
	nowEmpty := g.emptyCells
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.cells[ri][ci]
			if c.pCount() == 1 {
				for vi, p := range c.possible {
					val := values[vi]
					if p {
						c.setValue(val)
						g.emptyCells--
					}
				}
			}
		}
	}
	return nowEmpty - g.emptyCells
}

func drawHoriz() {
	fmt.Println("+---------+---------+---------+")
}

func drawVert() {
	fmt.Print("|")
}

func drawEndofRow() {
	drawVert()
	fmt.Println()
}

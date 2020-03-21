package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const gridSize int = 9

type gridDisplay int

const (
	showValues gridDisplay = iota
	showPCount
)

// gridCoord allows us to use the following code:
//  for ri, rn := range gridCoord
// this allows us to use ri for a zero-based value (0 to 8)
// or rn for a 1-based value (1 to 9) as required.
var gridCoord = [gridSize]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Grid is the entire game board
type Grid struct {
	cell        [gridSize][gridSize]*cell
	showWorking bool
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := newCell(ri, ci)
			g.cell[ri][ci] = c
		}
	}
	return g
}

func (g *Grid) returnCell(ri, ci int) *cell {
	return g.cell[ri][ci]
}

// ShowWorking sets whether the solver should explain its thinking
func (g *Grid) ShowWorking(sw bool) {
	g.showWorking = sw
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
				err = c.setValue(value(vr))
				if err == nil {
					c.status = cellOriginal
				} else {
					return err
				}
			}
			ri++
		}
	}
	return nil
}

// Display handles the grid output
func (g *Grid) Display(displayType ...gridDisplay) {
	dt := showValues
	if len(displayType) > 0 {
		dt = displayType[0]
	}

	for ri := range gridCoord {
		if ri%3 == 0 {
			drawHoriz()
		}
		for ci := range gridCoord {
			if ci%3 == 0 {
				drawVert()
			}
			c := g.returnCell(ri, ci)
			switch dt {
			case showValues:
				if g.showWorking {
					switch c.status {
					case cellNew:
						fmt.Printf("<%s>", c.val)
						c.status = cellSolved
					case cellSolved:
						fmt.Printf("-%s-", c.val)
					default:
						fmt.Printf(" %s ", c.val)
					}
				} else {
					fmt.Printf(" %s ", c.val)
				}
			case showPCount:
				pc, opv := c.pCount()
				if pc == 1 {
					fmt.Printf("<%s>", opv)
				} else {
					fmt.Printf(" %d ", pc)
				}
			}
		}
		drawEndofRow()
	}
	drawHoriz()
}

type solver struct {
	name   string
	solver func(*Grid) (int, error)
}

var solvers = []solver{
	solver{"OnlyPossibleValue", (*Grid).solveUseOPV},
	solver{"OPVByBlock", (*Grid).solveOPVbyBlock},
}

// Solve the grid
func (g *Grid) Solve() {
	pass := 1

	if g.showWorking {
		g.Display(showPCount)
		displayCollections()
	}

	for g.emptyCells() > 0 {
		fmt.Printf("Cells remaining: %d; Solver running; pass %d:\n", g.emptyCells(), pass)
		numSolved := 0
		for _, sv := range solvers {
			g.working(fmt.Sprintf("Running '%s' solver:", sv.name))
			ns, err := sv.solver(g)
			if err != nil {
				log.Fatalf("Error in %s: %v", sv.name, err)
			} else {
				numSolved += ns
			}
		}
		if g.showWorking {
			g.Display()
		}
		fmt.Printf("Pass %d: %d cells solved\n", pass, numSolved)
		if numSolved == 0 {
			fmt.Printf("Solver is stuck with %d empty cells remaining\n", g.emptyCells())
			return
		}
		pass++
	}
}

func (g *Grid) solveUseOPV() (int, error) {
	nowEmpty := g.emptyCells()
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.returnCell(ri, ci)
			pc, opv := c.pCount()
			if pc == 1 {
				g.working(fmt.Sprintf("  Cell (%d,%d) set to OPV: %s", ri, ci, opv))
				err := c.setValue(opv)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return nowEmpty - g.emptyCells(), nil
}

func (g *Grid) solveOPVbyBlock() (int, error) {
	nowEmpty := g.emptyCells()
	for bi := range gridCoord {
		for vi, val := range values {
			cc := blkColl[bi]
			pc := 0
			var cp *cell
			for _, c := range cc {
				if c.val != empty {
					continue
				}
				if c.possible[vi] {
					pc++
					cp = c
				}
			}
			if pc == 1 {
				g.working(fmt.Sprintf("  Cell (%d, %d) set to %s by block examination", cp.ri, cp.ci, val))
				err := cp.setValue(val)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return nowEmpty - g.emptyCells(), nil
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

func (g *Grid) working(msg string) {
	if g.showWorking {
		fmt.Printf("  %s\n", msg)
	}
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

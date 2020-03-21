package sudoku

import (
	"fmt"
	"log"
)

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

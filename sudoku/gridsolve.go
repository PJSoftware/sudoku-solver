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
	solver{"ExtendPossVal", (*Grid).solveExtendPossVal},
}

// Solve the grid
func (g *Grid) Solve() (int, int) {
	pass := 1

	if g.showWorking {
		g.Display(showPCount)
		g.cc.displayCollections()
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
		if numSolved == 0 {
			ec := g.emptyCells()
			fmt.Printf("Solver is stuck with %d empty cells remaining\n", ec)
			return ec, pass
		}
		pass++
	}
	return 0, pass - 1
}

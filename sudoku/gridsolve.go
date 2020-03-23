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
func (g *Grid) Solve() int {
	pass := 1

	if g.showWorking {
		g.Display(showPCount)
		gc.displayCollections()
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
			return ec
		}
		pass++
	}
	return 0
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
			cc := gc.blkColl[bi]
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

func (g *Grid) solveExtendPossVal() (int, error) {
	ext := 0
	for bi := range gridCoord {
		ec := gc.blkColl[bi].emptyCount()
		if ec >= 2 && ec <= 3 {
			rows := make(map[int]bool)
			cols := make(map[int]bool)
			for i := range gridCoord {
				c := gc.blkColl[bi][i]
				if c.val == empty {
					rows[c.ri] = true
					cols[c.ci] = true
				}
			}
			if len(rows) == 1 || len(cols) == 1 {
				unused := gc.blkColl[bi].unusedValues()

				var cc collection
				var fixed int
				ignore := make(map[int]bool)

				if len(rows) == 1 {
					for i := range rows {
						fixed = i
					}
					cc = gc.rowColl[fixed]
					for i := range cols {
						ignore[i] = true
					}
				} else { // len(cols) == 1
					for i := range cols {
						fixed = i
					}
					cc = gc.colColl[fixed]
					for i := range rows {
						ignore[i] = true
					}
				}

				for _, val := range unused {
					for i := range gridCoord {
						c := cc[i]
						skip := false
						if len(rows) == 1 {
							if _, ok := ignore[c.ci]; ok {
								skip = true
							}
						} else {
							if _, ok := ignore[c.ri]; ok {
								skip = true
							}
						}
						if skip {
							continue
						}

						vi := int(val) - 1
						if c.val == empty && c.possible[vi] {
							g.working(fmt.Sprintf("Empty Cell(%d,%d), value %s set to not possible", c.ri, c.ci, val))
							c.possible[vi] = false
							ext++
						}
					}
				}
			}
		}
	}
	return ext, nil
}

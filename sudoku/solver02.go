package sudoku

import "fmt"

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
				g.working(fmt.Sprintf("  Cell(%d, %d) set to %s by block examination", cp.ri, cp.ci, val))
				err := cp.setValue(val)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return nowEmpty - g.emptyCells(), nil
}

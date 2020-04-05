package sudoku

import "fmt"

func (g *Grid) solveUseOPV() (int, error) {
	nowEmpty := g.emptyCells()
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.returnCell(ri, ci)
			pc, opv := c.pCount()
			if pc == 1 {
				g.working(fmt.Sprintf("  Cell(%d,%d) set to OPV: %s", ri, ci, opv))
				err := c.setValue(opv)
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return nowEmpty - g.emptyCells(), nil
}

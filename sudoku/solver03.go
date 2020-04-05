package sudoku

import "fmt"

// solveExtendPossVal (solver 3) examines blocks containing only 2 or 3
// empty cells which are in a line (row or column); it considers whether
// any of the possible values are disallowed because setting them would
// prevent all valid moves in neighbouring blocks. In such a case, it
// sets the Possible values of the cell appropriately. This does not
// directly set the value of a cell, but may enable further progress.
func (g *Grid) solveExtendPossVal() (int, error) {
	ext := 0
	for bi := range gridCoord {
		ec := g.cc.blkColl[bi].emptyCount()
		if ec >= 2 && ec <= 3 {
			rows := make(map[int]bool)
			cols := make(map[int]bool)
			for i := range gridCoord {
				c := g.cc.blkColl[bi][i]
				if c.val == empty {
					rows[c.ri] = true
					cols[c.ci] = true
				}
			}
			if len(rows) == 1 || len(cols) == 1 {
				unused := g.cc.blkColl[bi].unusedValues()

				var cc collection
				var fixed int
				ignore := make(map[int]bool)

				if len(rows) == 1 {
					for i := range rows {
						fixed = i
					}
					cc = g.cc.rowColl[fixed]
					for i := range cols {
						ignore[i] = true
					}
				} else { // len(cols) == 1
					for i := range cols {
						fixed = i
					}
					cc = g.cc.colColl[fixed]
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
							g.working(fmt.Sprintf("  Empty Cell(%d,%d), value %s set to not possible", c.ri, c.ci, val))
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

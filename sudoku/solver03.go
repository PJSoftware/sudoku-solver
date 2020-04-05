package sudoku

import "fmt"

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

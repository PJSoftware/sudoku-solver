package sudoku

// Cell is a single square containing values
type cell struct {
	val        value
	possible   [gridSize]bool
	ri, ci, bi int
	opv        value // only possible value
}

func newCell(ri, ci int) *cell {
	c := new(cell)
	c.val = empty
	c.opv = empty
	bi := calcBlkIdx(ri, ci)
	c.ri = ri
	c.ci = ci
	c.bi = bi

	for vi := 0; vi < 9; vi++ {
		c.possible[vi] = true
	}

	rowColl[ri] = append(rowColl[ri], c)
	colColl[ci] = append(colColl[ci], c)
	blkColl[bi] = append(blkColl[bi], c)

	return c
}

func (c *cell) setValue(v value) bool {
	vi := int(v - 1)
	if c.possible[vi] {
		c.val = v
		for vi := range values {
			c.possible[vi] = false
		}

		for i := 0; i < 9; i++ {
			rowColl[c.ri][i].possible[vi] = false
			colColl[c.ci][i].possible[vi] = false
			blkColl[c.bi][i].possible[vi] = false
		}
		return true
	}

	return false
}

// pCount returns number of possible values for cell
// if only one possible value, return it too
func (c *cell) pCount() (int, value) {
	if c.opv != empty {
		return 1, c.opv
	}

	pc := 0
	opv := empty
	for vi, p := range values {
		if c.possible[vi] {
			opv = p
			pc++
		}
	}
	if pc != 1 {
		opv = empty
	}
	c.opv = opv
	return pc, opv
}

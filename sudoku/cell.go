package sudoku

import "fmt"

type cellStatus int

const (
	cellEmpty cellStatus = iota
	cellOriginal
	cellSolved
	cellNew
)

// Cell is a single square containing values
type cell struct {
	val        value
	possible   [gridSize]bool
	ri, ci, bi int
	opv        value // only possible value
	status     cellStatus
	parent     *Grid
}

func newCell(ri, ci int, g *Grid) *cell {
	c := new(cell)
	c.val = empty
	c.opv = empty
	c.status = cellEmpty
	c.parent = g
	bi := calcBlkIdx(ri, ci)
	c.ri = ri
	c.ci = ci
	c.bi = bi

	for vi := 0; vi < 9; vi++ {
		c.possible[vi] = true
	}

	g.cc.rowColl[ri] = append(g.cc.rowColl[ri], c)
	g.cc.colColl[ci] = append(g.cc.colColl[ci], c)
	g.cc.blkColl[bi] = append(g.cc.blkColl[bi], c)

	return c
}

func (c *cell) setValue(v value) error {
	if c.val != empty {
		return fmt.Errorf("Cannot set Cell(%d,%d) to %s; cell already set to %s", c.ri, c.ci, v, c.val)
	}

	vi := int(v - 1)
	if !c.possible[vi] {
		return fmt.Errorf("Cannot set Cell(%d,%d) to impossible value %s", c.ri, c.ci, v)
	}

	c.val = v
	c.opv = empty
	c.status = cellNew

	for i := range gridCoord {
		c.parent.cc.rowColl[c.ri][i].possible[vi] = false
		c.parent.cc.colColl[c.ci][i].possible[vi] = false
		c.parent.cc.blkColl[c.bi][i].possible[vi] = false
	}

	for vi := range values {
		c.possible[vi] = false
	}
	return nil
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

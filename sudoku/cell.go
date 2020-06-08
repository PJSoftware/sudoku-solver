package sudoku

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
	ri, ci, bi int
	parent     *Grid
}

func newCell(ri, ci int, g *Grid) *cell {
	c := new(cell)
	c.val = empty
	c.parent = g
	bi := calcBlkIdx(ri, ci)
	c.ri = ri
	c.ci = ci
	c.bi = bi

	g.cc.rowColl[ri] = append(g.cc.rowColl[ri], c)
	g.cc.colColl[ci] = append(g.cc.colColl[ci], c)
	g.cc.blkColl[bi] = append(g.cc.blkColl[bi], c)

	return c
}

func (c *cell) isPossible(val value) bool {
	isUsed := valueIn(c.parent.cc.rowColl[c.ri], val) ||
		valueIn(c.parent.cc.colColl[c.ci], val) ||
		valueIn(c.parent.cc.blkColl[c.bi], val)
	return !isUsed
}

func valueIn(coll collection, val value) bool {
	for _, c := range coll {
		if val == c.val {
			return true
		}
	}
	return false
}

package sudoku

// Cell is a single square containing values
type cell struct {
	val      value
	possible map[value]bool
}

func newCell() *cell {
	c := new(cell)
	c.val = empty
	c.possible = make(map[value]bool)
	c.setAllPossible(true)
	return c
}

func (c *cell) setValue(v value) {
	c.val = v
	c.setAllPossible(false)
	c.possible[v] = true
}

func (c *cell) setAllPossible(isPossible bool) {
	for v := val1; v <= val9; v++ {
		c.possible[v] = isPossible
	}

}

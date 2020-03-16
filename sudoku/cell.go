package sudoku

// Cell is a single square containing values
type cell struct {
	val         value
	possible    map[value]bool
	parentBlock *block
}

type cellRow [9]cell
type cellCol [9]cell

func newCell(b *block) *cell {
	c := new(cell)
	c.parentBlock = b
	c.val = empty
	c.parentBlock.parentGrid.emptyCells++
	c.possible = make(map[value]bool)
	c.setAllPossible(true)
	return c
}

func (c *cell) setValue(v value) {
	c.val = v
	c.setAllPossible(false)
}

func (c *cell) setAllPossible(isPossible bool) {
	for v := val1; v <= val9; v++ {
		c.possible[v] = isPossible
	}

}

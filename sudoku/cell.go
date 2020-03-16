package sudoku

// Cell is a single square containing values
type cell struct {
	val      value
	possible map[value]bool
}

type cellCollection [gridSize]*cell

func (cc cellCollection) notPossible(v value) {
	for i := 0; i < gridSize; i++ {
		cc[i].possible[v] = false
	}
}

func newCell(g *Grid) *cell {
	c := new(cell)
	c.val = empty
	g.emptyCells++
	c.possible = make(map[value]bool)
	c.setAllPossible(true)
	return c
}

func (c *cell) canSet(v value) bool {
	return c.possible[v]
}

func (c *cell) setValue(v value) bool {
	c.val = v
	c.setAllPossible(false)
	return true
}

func (c *cell) setAllPossible(isPossible bool) {
	for v := val1; v <= val9; v++ {
		c.possible[v] = isPossible
	}

}

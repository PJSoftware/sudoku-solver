package sudoku

// Cell is a single square containing values
type cell struct {
	val      value
	possible map[value]bool
	status   cellStatus
	pCount   int
}

type cellStatus int

const (
	original cellStatus = iota + 1
	calculated
)

type cellCollection [gridSize]*cell

func (cc cellCollection) notPossible(v value) {
	for i := range gridCoord {
		c := cc[i]
		c.setPossible(v, false)
	}
}

func newCell() *cell {
	c := new(cell)
	c.val = empty
	c.possible = make(map[value]bool)
	c.pCount = 0
	c.setAllPossible(true)
	return c
}

func (c *cell) canSet(v value) bool {
	return c.possible[v]
}

func (c *cell) setValue(v value) bool {
	c.val = v
	c.status = calculated
	c.setAllPossible(false)
	return true
}

func (c *cell) setAllPossible(possibility bool) {
	for v := val1; v <= val9; v++ {
		c.setPossible(v, possibility)
	}
}

func (c *cell) setPossible(v value, possibility bool) {
	if c.possible[v] == possibility {
		return
	}
	c.possible[v] = possibility
	if possibility {
		c.pCount++
	} else {
		c.pCount--
	}
}

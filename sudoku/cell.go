package sudoku

// Cell is a single square containing values
type Cell struct {
	value    Value
	possible map[Value]bool
}

// NewCell returns a new, empty cell
func NewCell() *Cell {
	c := new(Cell)
	c.value = empty
	c.possible = make(map[Value]bool)
	// start with all values being possible
	for v := val1; v <= val9; v++ {
		c.possible[v] = true
	}
	return c
}

// SetValue sets a cell's value (and resets possibles)
// TODO: The deeper I get into this code, the more I think all interfaces
// except Grid should remain unexported
func (c *Cell) SetValue(v Value) {
	c.value = v
	for lv := val1; lv <= val9; lv++ {
		c.possible[lv] = false
	}
	c.possible[v] = true
}

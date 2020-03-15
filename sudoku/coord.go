package sudoku

type coord struct {
	row, col int
	bR, bC   int
	gR, gC   int
}

var cache = [9][2]int{
	{0, 0}, {0, 1}, {0, 2},
	{1, 0}, {1, 1}, {1, 2},
	{2, 0}, {2, 1}, {2, 2},
}

func newCoord(row, col int) *coord {
	coord := new(coord)
	if isValid(row) && isValid(col) {
		coord.row = row
		coord.col = col
		coord.convert()
		return coord
	}
	return nil
}

// convertCoord converts a coordinate in the range 1 to 9,
// into a GridCoord, BlockCoord combo (each in the range 0-2)
func (c *coord) convert() {
	c.gR, c.bR = cache[c.row-1][0], cache[c.row-1][1]
	c.gC, c.bC = cache[c.col-1][0], cache[c.col-1][1]
}

func isValid(c int) bool {
	return c >= 1 && c <= 9
}

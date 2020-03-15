package sudoku

// Block is the subset of a grid. A block is 3x3 Cells.
// Each Cell of a Block must contain a unique Value
type block struct {
	blk [3][3]cell
}

func newBlock() *block {
	b := new(block)
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			b.blk[x][y] = *newCell()
		}
	}
	return b
}

package sudoku

// Block is the subset of a grid. A block is 3x3 Cells.
// Each Cell of a Block must contain a unique Value
type Block struct {
	block [3][3]Cell
}

// NewBlock returns a new, empty block
func NewBlock() *Block {
	b := new(Block)
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			b.block[x][y] = *NewCell()
		}
	}
	return b
}

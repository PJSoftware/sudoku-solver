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

func (b *block) updatePossibility(rc coord, v value) {
	for r := 0; r <= 2; r++ {
		for c := 0; c <= 2; c++ {
			if !(r == rc.bR && c == rc.bC) {
				b.blk[r][c].possible[v] = false
			}
		}
	}
}

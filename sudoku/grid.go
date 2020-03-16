package sudoku

import (
	"bufio"
	"os"
	"regexp"
)

// Grid is the entire game board, made of 3x3 Blocks.
// In addition to each Block containing 1 of each possible Values,
// each row and column in a Block must contain 1 of each Value.
type Grid struct {
	grid       [3][3]block
	emptyCells int
	rows       [9]cellRow
	cols       [9]cellCol
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			g.grid[x][y] = *newBlock(g)
		}
	}
	return g
}

// Import from specified .sp file
func (g *Grid) Import(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 1
	for scanner.Scan() {
		line := scanner.Text()
		if ok, _ := regexp.MatchString(`^[*1-9]{9}$`, line); ok {
			r := []rune(line)
			for col := 1; col <= 9; col++ {
				rv := r[col-1]
				if rv == '*' {
					break
				}
				g.SetValue(row, col, value(rv))
			}
		}
	}
	return nil
}

// SetValue sets the value of the specified cell. This includes
// recalculating all valid possible values appropriately
func (g *Grid) SetValue(row, col int, v value) {
	coord := newCoord(row, col)
	b := g.block(*coord)
	b.updatePossibility(*coord, v)
	c := g.cell(*coord)
	c.setValue(v)
}

// Block takes a row and col (in the range 1 to 9) and returns
// a pointer to the correct Block.
func (g *Grid) block(rc coord) *block {
	b := g.grid[rc.gR][rc.gC]
	return &b
}

// Cell takes a row and col (in the range 1 to 9) and returns
// a pointer to the correct Cell
func (g *Grid) cell(rc coord) *cell {
	c := g.grid[rc.gR][rc.gC].blk[rc.bR][rc.bC]
	return &c
}

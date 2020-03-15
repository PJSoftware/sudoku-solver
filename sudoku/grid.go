package sudoku

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// Grid is the entire game board, made of 3x3 Blocks.
// In addition to each Block containing 1 of each possible Values,
// each row and column in a Block must contain 1 of each Value.
type Grid struct {
	grid [3][3]Block
}

// NewGrid returns a new, empty grid
func NewGrid() *Grid {
	g := new(Grid)
	for x := 0; x <= 2; x++ {
		for y := 0; y <= 2; y++ {
			g.grid[x][y] = *NewBlock()
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
				g.SetValue(row, col, Value(rv))
			}
		}
	}
	return nil
}

// SetValue sets the value of the specified cell. This includes
// recalculating all valid possible values appropriately
func (g *Grid) SetValue(row, col int, v Value) {
	c := g.Cell(row, col)
	c.SetValue(v)
}

// Block takes a row and col (in the range 1 to 9) and returns
// a pointer to the correct Block.
//
// Possibly should not be visible outside this package?
func (g *Grid) Block(row, col int) *Block {
	gR, _, err1 := convertCoord(row)
	gC, _, err2 := convertCoord(col)
	if err1 == nil && err2 == nil {
		b := g.grid[gR][gC]
		return &b
	}
	log.Fatalf("Coordinates (%d, %d) not valid", row, col)
	return nil
}

// Cell takes a row and col (in the range 1 to 9) and returns
// a pointer to the correct Cell
func (g *Grid) Cell(row, col int) *Cell {
	gR, bR, err1 := convertCoord(row)
	gC, bC, err2 := convertCoord(col)
	if err1 == nil && err2 == nil {
		c := g.grid[gR][gC].block[bR][bC]
		return &c
	}
	log.Fatalf("Coordinates (%d, %d) not valid", row, col)
	return nil
}

// convertCoord converts a coordinate in the range 1 to 9,
// into a GridCoord, BlockCoord combo (each in the range 0-2)
func convertCoord(coord int) (int, int, error) {
	var cache = [9][2]int{
		{0, 0}, {0, 1}, {0, 2},
		{1, 0}, {1, 1}, {1, 2},
		{2, 0}, {2, 1}, {2, 2},
	}
	if coord >= 1 && coord <= 9 {
		return cache[coord-1][0], cache[coord-1][0], nil
	}
	return -1, -1, fmt.Errorf("coord %d out of expected bounds (1-9)", coord)
}

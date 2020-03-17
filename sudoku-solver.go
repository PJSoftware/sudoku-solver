package main

import (
	"fmt"

	"./sudoku"
)

// 1. Initialise grid
// 2. Solve grid
// 3. Display solution

const defaultPuzzleFile string = "puzzles/easy.sp"

func main() {
	grid := sudoku.NewGrid()
	err := grid.Import(defaultPuzzleFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	grid.Display()

	grid.Solve()
	grid.Display()
}

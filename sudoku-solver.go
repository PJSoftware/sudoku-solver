package main

import (
	"flag"
	"fmt"

	"./sudoku"
)

// 1. Initialise grid
// 2. Solve grid
// 3. Display solution

func main() {
	puzzle := flag.String("puzzle", "easy", "enter name of predefined puzzle to solve (from puzzles folder)")
	flag.Parse()

	grid := sudoku.NewGrid()
	puzzleFile := fmt.Sprintf("puzzles/%s.sp", *puzzle)
	fmt.Printf("Solving '%s'\n", puzzleFile)

	err := grid.Import(puzzleFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	grid.Display()
	grid.Solve()
	grid.Display()
}

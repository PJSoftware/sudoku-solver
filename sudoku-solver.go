package main

import "./sudoku"

// 1. Initialise grid
// 2. Solve grid
// 3. Display solution

const defaultPuzzleFile string = "puzzles/easy.sp"

func main() {
	grid := sudoku.NewGrid()
	grid.Import(defaultPuzzleFile)
	grid.Display()
}

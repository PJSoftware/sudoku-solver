package sudoku

import (
	"fmt"
)

// Solve the grid
func (g *Grid) Solve() int {
	g.solveRecursive()

	ecc := g.emptyCells()
	if ecc > 0 {
		fmt.Printf("Solver is stuck with %d empty cells remaining\n", ecc)
	}
	return ecc
}

// solveRecursive (solver 0) uses a recursive, brute force approach
// per the python script I saw on computerphile
//
// The presenter in that video stated that attempting to solve the game
// in the way a player would solve the game is ... difficult to code.
// I am aware!
func (g *Grid) solveRecursive() {
	for ri := range gridCoord {
		for ci := range gridCoord {
			c := g.returnCell(ri, ci)
			if c.val == empty {
				for _, val := range values {
					if c.isPossible(val) {
						c.val = val
						g.solveRecursive()
						if g.emptyCells() == 0 {
							return
						}
						c.val = empty
					}
				}
				return
			}
		}
	}
}

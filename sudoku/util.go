package sudoku

func calcBlkIdx(ri, ci int) int {
	// r := ri / 3
	// c := ci / 3
	return (ri/3)*3 + (ci / 3)
}

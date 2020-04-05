package sudoku

func calcBlkIdx(ri, ci int) int {
	return (ri/3)*3 + (ci / 3)
}

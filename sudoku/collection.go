package sudoku

import (
	"fmt"
	"sort"
)

type collection []*cell

var rowColl [gridSize]collection
var colColl [gridSize]collection
var blkColl [gridSize]collection

func displayCollections() {
	displayColl("row", rowColl)
	displayColl("column", colColl)
	displayColl("block", blkColl)
}

func displayColl(cName string, cArray [gridSize]collection) {
	fmt.Printf("Displaying %s collections:\n", cName)
	for i := range gridCoord {
		fmt.Printf("  %s %d values: ", cName, i)
		var sv []value
		for ci := range cArray[i] {
			c := cArray[i][ci]
			sv = append(sv, c.val)
		}
		sort.Slice(sv, func(i, j int) bool {
			return sv[i] < sv[j]
		})
		for si := range sv {
			if sv[si] == empty {
				continue
			}
			fmt.Printf("%s", sv[si])
		}
		fmt.Println()
	}
}

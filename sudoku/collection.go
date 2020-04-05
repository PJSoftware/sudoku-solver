package sudoku

import (
	"fmt"
	"sort"
)

type collection []*cell

type cellCollections struct {
	rowColl [gridSize]collection
	colColl [gridSize]collection
	blkColl [gridSize]collection
}

func (gc *cellCollections) displayCollections() {
	displayColl("row", gc.rowColl)
	displayColl("column", gc.colColl)
	displayColl("block", gc.blkColl)
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

func (cc collection) emptyCount() int {
	ec := 0
	for i := range gridCoord {
		if cc[i].val == empty {
			ec++
		}
	}
	return ec
}

func (cc collection) usedValues() []value {
	used := cc.usedMap()
	var rv []value
	for val := range used {
		rv = append(rv, val)
	}
	return rv
}

func (cc collection) unusedValues() []value {
	used := cc.usedMap()
	var unused []value
	for _, val := range values {
		if _, ok := used[val]; !ok {
			unused = append(unused, val)
		}
	}
	return unused
}

func (cc collection) usedMap() map[value]bool {
	used := make(map[value]bool)
	for i := range gridCoord {
		if cc[i].val != empty {
			used[cc[i].val] = true
		}
	}
	return used
}

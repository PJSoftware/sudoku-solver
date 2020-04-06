package sudoku

import (
	"fmt"
	"strings"
)

const s3DEBUG bool = false

type s3Solver struct {
	rows        map[int]bool // determines rows used
	cols        map[int]bool // determines columns used
	ec          int          // empty cells in block
	block       collection   // block we are working with
	coll        collection   // collection to examine
	unused      []value      // values unused in this block
	ignore      map[int]bool // cells in collection to ignore
	ext         int          // how many extensions were processed
	showWorking bool         // Additional working info for s3
}

// solveExtendPossVal (solver 3) examines blocks containing only 2 or 3
// empty cells which are in a line (row or column); it considers whether
// any of the possible values are disallowed because setting them would
// prevent all valid moves in neighbouring blocks. In such a case, it
// sets the Possible values of the cell appropriately. This does not
// directly set the value of a cell, but may enable further progress.
func (g *Grid) solveExtendPossVal() (int, error) {
	s3 := new(s3Solver)
	s3.showWorking = s3DEBUG
	if s3.showWorking {
		g.Display()
	}

	for bi := range gridCoord {
		if s3.worthConsidering(g.cc.blkColl[bi]) {
			s3.examineEmpty()

			if s3.emptyInLine() {
				s3.findUnusedValues()
				s3.chooseCollection(g.cc)

				for _, val := range s3.unused {
					for _, c := range s3.coll {
						s3.extendPossValue(c, val)
					}
				}
			}
		}
	}
	return s3.ext, nil
}

func (s3 *s3Solver) worthConsidering(block collection) bool {
	s3.block = block
	s3.ec = block.emptyCount()
	wc := s3.ec >= 2 && s3.ec <= 3
	s3.working(fmt.Sprintf("emptyCount = %d, worthConsidering = %v", s3.ec, wc), 1)
	return wc
}

func (s3 *s3Solver) examineEmpty() {
	s3.rows = make(map[int]bool)
	s3.cols = make(map[int]bool)
	for _, c := range s3.block {
		if c.val == empty {
			s3.rows[c.ri] = true
			s3.cols[c.ci] = true
		}
	}
}

func (s3 *s3Solver) emptyInLine() bool {
	il := len(s3.rows) == 1 || len(s3.cols) == 1
	s3.working(fmt.Sprintf("empty cells are in a line: %v", il), 2)
	return il
}

func (s3 *s3Solver) findUnusedValues() {
	s3.unused = s3.block.unusedValues()
	s3.working(fmt.Sprintf("Unused values: %v", s3.unused), 2)
}

func (s3 *s3Solver) isRow() bool {
	return len(s3.cols) == 1
}

func (s3 *s3Solver) isCol() bool {
	return len(s3.rows) == 1
}

func (s3 *s3Solver) chooseCollection(cc *cellCollections) {
	var fixed int
	s3.ignore = make(map[int]bool)

	if s3.isCol() {
		for i := range s3.rows {
			fixed = i
		}
		s3.coll = cc.rowColl[fixed]
		for ci := range s3.cols {
			s3.ignore[ci] = true
		}
	} else { // s3.isRow()
		for i := range s3.cols {
			fixed = i
		}
		s3.coll = cc.colColl[fixed]
		for ri := range s3.rows {
			s3.ignore[ri] = true
		}
	}
}

func (s3 *s3Solver) extendPossValue(c *cell, val value) {
	if s3.isCol() {
		s3.working(fmt.Sprintf("Examining column %d", c.ci), 3)
		if _, ok := s3.ignore[c.ci]; ok {
			return
		}
	} else {
		s3.working(fmt.Sprintf("Examining row %d", c.ri), 3)
		if _, ok := s3.ignore[c.ri]; ok {
			return
		}
	}

	s3.working(fmt.Sprintf("Looking at value %s in Cell(%d,%d)", val, c.ri, c.ci), 4)
	vi := int(val) - 1
	if c.val == empty && c.possible[vi] {
		c.parent.working(fmt.Sprintf("Empty Cell(%d,%d), value %s set to not possible", c.ri, c.ci, val), 1)
		c.possible[vi] = false
		s3.ext++
	}

}

func (s3 *s3Solver) working(msg string, indent int) {
	if s3.showWorking {
		is := strings.Repeat(" -", indent)
		fmt.Printf("%s%s\n", is, msg)
	}
}

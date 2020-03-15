package sudoku

import "strconv"

// Value is an enum containing an int from 1 to 9 inclusive
// with an optional 'empty' value (0)
type Value int

const (
	empty Value = 0
	val1  Value = 1
	val2  Value = 2
	val3  Value = 3
	val4  Value = 4
	val5  Value = 5
	val6  Value = 6
	val7  Value = 7
	val8  Value = 8
	val9  Value = 9
)

// String converts a Value to a string
func (v Value) String() string {
	if v == empty {
		return " "
	}
	return strconv.Itoa(int(v))
}

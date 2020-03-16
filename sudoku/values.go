package sudoku

import "strconv"

// value is an enum containing an int from 1 to 9 inclusive
// with an optional 'empty' value (0)
type value int

const (
	empty value = iota
	val1
	val2
	val3
	val4
	val5
	val6
	val7
	val8
	val9
)

// String converts a value to a string
func (v value) String() string {
	if v == empty {
		return " "
	}
	return strconv.Itoa(int(v))
}

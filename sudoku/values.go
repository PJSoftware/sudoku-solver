package sudoku

import "strconv"

// value is an enum containing an int from 1 to 9 inclusive
// with an optional 'empty' value (0)
type value int

const (
	empty value = 0
	val1  value = 1
	val2  value = 2
	val3  value = 3
	val4  value = 4
	val5  value = 5
	val6  value = 6
	val7  value = 7
	val8  value = 8
	val9  value = 9
)

// String converts a value to a string
func (v value) String() string {
	if v == empty {
		return " "
	}
	return strconv.Itoa(int(v))
}

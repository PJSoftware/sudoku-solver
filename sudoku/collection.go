package sudoku

type collection []*cell

var rowColl [gridSize]collection
var colColl [gridSize]collection
var blkColl [gridSize]collection

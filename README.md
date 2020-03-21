# sudoku-solver

## [#100DaysofCode](https://github.com/PJSoftware/100-days-of-code) Project 1: Sudoku Solver

I first had the idea to write my own **Sudoku Solver** probably five years ago, and I dabbled a little with some Perl code, but I never took it very far.

The idea came to me because I was actually attempting to solve an extended grid of five interlinked Sudoku squares, and I wanted my code to be able to solve that kind of complex puzzle--but I could never quite wrap my head around how to model it. Typical for me, I think I was trying to run before I could walk; while I may ultimately attempt the same with this code, I shall first attempt to get a basic one-grid solver working.

## Other Solvers

A quick google search will show you that there are already plenty of Sudoku Solvers out there.

Of particular interest to me is the [Sudoku Wiki Solver](https://www.sudokuwiki.org/sudoku.htm), which appears at first glance to contain plenty of information on solving strategies--no doubt far more than I have considered. For now I shall keep that in reserve until I get stuck, and bumble along with my own naive ideas!

## Requirements

### Puzzle File Format

Obviously one of the first things we will need is the ability to import puzzle data from a file, so our solver has something to work with.

At this stage I'm considering a simple ASCII grid of numbers and spaces; simple, easy to create, easy to read. This may change later, but it will do for now.

### Modelling the Puzzle

My first thought was to model the puzzle as a 3x3 Grid of 3x3 Blocks. However, while this potentially simplified validation of any of the Blocks, it did not help (and possibly even complicated) row and column validation.

Once I realised that row, column, and block could all be handled the same way--as a collection of 9 cells, which could be set once at grid creation and then simply referred to as required--I ditched the 3x3(x3x3) nesting, and changed to a plain 9x9 grid.

### Solving the Puzzle

I am taking the preliminary approach of scanning the grid every time a known value is entered, updating the "possible" values for each empty cell. For simple puzzles this may be enough, but more complex solution strategies will be added as I need them.

## Solver Code

### One Possible Value (OPV)

This was my first thought: simply check which cells only have one possible value (based on the other cells in their row/col/block "neighbourhoods") and set them to that value. This actually worked for some easy puzzles.

### OPV By Block

Where the simple OPV approach failed, the next step was to examine the cells in each block, examining each possible value, and determining which can only be placed in one possible cell. For instance, a block may have four empty cells, and one cell may have possible values of "2", "3", "5" -- but after examining the other cells, none of them can take the "3" so it must logically belong in this cell.

### Extend Possible Values

If a block has only two or three empty cells, all in a row or column, it can be determined that the possible values which can be placed in that row or column will affect the neighbouring blocks' possible values. Identifying such cases may give us enough information to fill in another cell.

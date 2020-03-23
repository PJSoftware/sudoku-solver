package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"./sudoku"
)

func main() {
	puzzle := flag.String("puzzle", "easy", "enter name of predefined puzzle to solve (from puzzles folder)")
	working := flag.Bool("working", false, "show working while solving puzzle")
	all := flag.Bool("all", false, "process all available puzzles")
	flag.Parse()

	puzzles := puzzleList(*all, *puzzle)
	np := len(puzzles)

	if np == 0 {
		fmt.Printf("Unable to open '%s'\n", puzzlePath(*puzzle))
		return
	}

	solved := 0
	for _, p := range puzzles {
		grid := sudoku.NewGrid()
		grid.ShowWorking(*working)

		fmt.Printf("Solving '%s'\n", p)

		err := grid.Import(p)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if np == 1 {
			grid.Display()
		}

		ecc := grid.Solve()
		if ecc == 0 {
			solved++
		}

		if np > 1 && ecc == 0 && !*working {
			fmt.Println("Puzzle solved!")
		} else {
			grid.Display()
		}
		fmt.Println()
	}
	fmt.Printf("%d of %d puzzles solved!\n", solved, len(puzzles))
}

func puzzleList(all bool, puzzle string) []string {
	var list []string
	var rv []string
	if all {
		files, err := ioutil.ReadDir(puzzlePath(""))
		if err == nil {
			for _, file := range files {
				list = append(list, puzzlePath(file.Name()))
			}
		}
	}
	if len(list) == 0 {
		list = append(list, puzzlePath(puzzle))
	}

	for _, f := range list {
		info, err := os.Stat(f)
		if !os.IsNotExist(err) && !info.IsDir() {
			rv = append(rv, f)
		}
	}
	return rv
}

func puzzlePath(fn string) string {
	pDir := "./puzzles"
	if fn == "" {
		return pDir
	}
	if ok, _ := regexp.MatchString("[.]sp$", fn); !ok {
		fn = fn + ".sp"
	}
	fmt.Println(fn)
	return (fmt.Sprintf("%s/%s", pDir, fn))
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/pjsoftware/sudoku-solver/sudoku"
)

type solution struct {
	file   string
	passes int
	grid   *sudoku.Grid
}

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
	var sl []solution

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

		ecc, npr := grid.Solve()
		if ecc == 0 {
			solved++
			sl = append(sl, solution{p, npr, nil})
		} else {
			sl = append(sl, solution{p, npr, grid})
		}

		if ecc == 0 {
			fmt.Println("Puzzle solved!")
		} else {
			fmt.Println("Unable to solve puzzle!")
		}
		if np == 1 && !*working {
			grid.Display()
		}
		fmt.Println()
	}

	for _, sol := range sl {
		if sol.grid == nil {
			fmt.Printf("%s solved in %d passes\n", sol.file, sol.passes)
		}
	}
	fmt.Printf("%d of %d puzzles solved!\n\n", solved, len(puzzles))

	for _, sol := range sl {
		if sol.grid != nil {
			fmt.Printf("%s could not be solved (%d passes)\n", sol.file, sol.passes)
			sol.grid.ShowWorking(false)
			sol.grid.Display()
			fmt.Println()
		}
	}
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
	pDir := "puzzles"
	if fn == "" {
		return pDir
	}
	if ok, _ := regexp.MatchString("[.]sp$", fn); !ok {
		fn = fn + ".sp"
	}
	fmt.Println(fn)
	return (fmt.Sprintf("%s/%s", pDir, fn))
}

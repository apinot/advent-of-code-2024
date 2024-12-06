package main

import (
	"fmt"
	"os"
	"strings"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

func main() {
	input := readInput()
	grid, guard := parseInput(input)

	result := doLogic(grid, guard)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day6.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

const OBSTRUCTIONS_RUNE = '#'
const GUARD_RUNE = '^'

type Dir struct {
	row int
	col int
}

type Pos struct {
	row int
	col int
}

func getNextDir(dir Dir) Dir {
	r := dir.row
	c := dir.col

	// North -> EAST
	if r == -1 && c == 0 {
		return Dir{0, 1}
	}

	// EAST -> SOUTH
	if r == 0 && c == 1 {
		return Dir{1, 0}
	}

	// SOUTH -> WEST
	if r == 1 && c == 0 {
		return Dir{0, -1}
	}

	// DEFAULT TO NORTH
	return Dir{-1, 0}
}

func getNextPosFromDir(pos Pos, dir Dir) Pos {
	nRow := pos.row + dir.row
	nCol := pos.col + dir.col
	return Pos{nRow, nCol}
}

type Guard struct {
	pos Pos
	dir Dir
}

type Grid struct {
	elements  [][]bool
	movements [][]Dir
	nbRows    int
	nbCols    int
}

func isEmptyDir(d Dir) bool {
	return d.col == 0 && d.row == 0
}

func isSameDir(da Dir, db Dir) bool {
	return da.row == db.row && da.col == db.col
}

func parseInput(input string) (Grid, Guard) {
	var e [][]bool
	var m [][]Dir
	var guard Guard

	rows := strings.Split(input, "\n")
	for r := 0; r < len(rows); r++ {
		var eRow []bool
		var mRow []Dir

		cols := []rune(rows[r])
		for c := 0; c < len(cols); c++ {
			v := cols[c]
			if v == OBSTRUCTIONS_RUNE {
				eRow = append(eRow, false)
			} else {
				eRow = append(eRow, true)
			}

			if v == GUARD_RUNE {
				guard = Guard{Pos{r, c}, Dir{-1, 0}}
			}

			mRow = append(mRow, Dir{0, 0})
		}

		e = append(e, eRow)
		m = append(m, mRow)
	}

	return Grid{e, m, len(m), len(m[0])}, guard
}

func doLogic(grid Grid, guard Guard) int {
	movements := 0
	move := true
	for move {
		if isEmptyDir(grid.movements[guard.pos.row][guard.pos.col]) {
			movements++
			grid.movements[guard.pos.row][guard.pos.col] = guard.dir
		}

		nextPos := getNextPosFromDir(guard.pos, guard.dir)

		// Out of grid, so stop
		if !inGrid(grid, nextPos) {
			move = false
			continue
		}

		// Can't move, so rotate to right
		if !canMove(grid, nextPos) {
			guard.dir = getNextDir(guard.dir)
			continue
		}

		// Apply movement
		guard.pos = nextPos
	}

	displayGrid(grid)

	return movements
}

func inGrid(grid Grid, pos Pos) bool {
	row := pos.row
	col := pos.col
	return row >= 0 && row < grid.nbRows && col >= 0 && col < grid.nbCols
}

func canMove(g Grid, p Pos) bool {
	return g.elements[p.row][p.col]
}

func displayGrid(g Grid) {
	for r := 0; r < g.nbRows; r++ {
		str := ""
		for c := 0; c < g.nbCols; c++ {
			if !g.elements[r][c] {
				str += string(OBSTRUCTIONS_RUNE)
				continue
			}
			if !isEmptyDir(g.movements[r][c]) {
				str += "X"
				continue
			}
			str += "."
		}
		fmt.Println(str)
	}
}

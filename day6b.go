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

func cloneGuard(g Guard) Guard {
	return Guard{Pos{g.pos.row, g.pos.col}, Dir{g.dir.row, g.dir.col}}
}

type Grid struct {
	elements  [][]bool
	movements [][]Dir
	nbRows    int
	nbCols    int
}

func cloneGrid(g Grid) Grid {
	var e [][]bool
	var m [][]Dir

	for r := 0; r < g.nbRows; r++ {
		var eRow []bool
		var mRow []Dir

		for c := 0; c < g.nbCols; c++ {
			eRow = append(eRow, g.elements[r][c])
			mRow = append(mRow, Dir{0, 0})
		}

		e = append(e, eRow)
		m = append(m, mRow)
	}

	return Grid{e, m, g.nbRows, g.nbCols}
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
	nbLoops := 0

	for r := 0; r < grid.nbRows; r++ {
		for c := 0; c < grid.nbCols; c++ {
			if !grid.elements[r][c] {
				continue
			}

			tmpGrid := cloneGrid(grid)
			tmpGrid.elements[r][c] = false
			tmpGuard := cloneGuard(guard)

			if checkLoop(tmpGrid, tmpGuard) {
				nbLoops++
			}
		}
	}
	return nbLoops
}

func checkLoop(grid Grid, guard Guard) bool {
	move := true
	nbMove := 0
	for move {
		grid.movements[guard.pos.row][guard.pos.col] = Dir{guard.dir.row, guard.dir.col}

		nextPos := getNextPosFromDir(guard.pos, guard.dir)

		// Out of grid, so stop
		if !inGrid(grid, nextPos) {
			move = false
			continue
		}

		// Can't move, so rotate to right
		if !canMove(grid, nextPos) {
			guard.dir = getNextDir(guard.dir)
			grid.movements[guard.pos.row][guard.pos.col] = Dir{guard.dir.row, guard.dir.col}
			continue
		}

		// Apply movement
		guard.pos = nextPos
		movementPosDir := grid.movements[guard.pos.row][guard.pos.col]
		if !isEmptyDir(movementPosDir) && isSameDir(movementPosDir, guard.dir) {
			return true
		}

		nbMove++
		// If it move far far away, then its a loop ^^
		if nbMove > 1000000 {
			return true
		}
	}

	return false
}

func inGrid(grid Grid, pos Pos) bool {
	row := pos.row
	col := pos.col
	return row >= 0 && row < grid.nbRows && col >= 0 && col < grid.nbCols
}

func canMove(g Grid, p Pos) bool {
	return g.elements[p.row][p.col]
}

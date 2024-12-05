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
	charGrid := parseInput(input)

	result := doLogic(charGrid)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day4.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func parseInput(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
}

func doLogic(grid [][]rune) int {
	found := 0

	for row, cols := range grid {
		for col, _ := range cols {
			found += countInAllDir(grid, row, col)
		}
	}

	return found
}

func countInAllDir(grid [][]rune, row int, col int) int {
	if !checkStringValueInGrid(grid, row, col, "X") {
		return 0
	}

	result := 0

	if checkN(grid, row, col) {
		result++
	}
	if checkNE(grid, row, col) {
		result++
	}
	if checkE(grid, row, col) {
		result++
	}
	if checkSE(grid, row, col) {
		result++
	}
	if checkS(grid, row, col) {
		result++
	}
	if checkSW(grid, row, col) {
		result++
	}
	if checkW(grid, row, col) {
		result++
	}
	if checkNW(grid, row, col) {
		result++
	}
	return result
}

func checkN(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row-i, col, letter) {
			return false
		}
	}
	return true
}

func checkNE(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row-i, col+i, letter) {
			return false
		}
	}
	return true
}

func checkE(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row, col+i, letter) {
			return false
		}
	}
	return true
}

func checkSE(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row+i, col+i, letter) {
			return false
		}
	}
	return true
}

func checkS(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row+i, col, letter) {
			return false
		}
	}
	return true
}

func checkSW(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row+i, col-i, letter) {
			return false
		}
	}
	return true
}

func checkW(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row, col-i, letter) {
			return false
		}
	}
	return true
}

func checkNW(grid [][]rune, row int, col int) bool {
	xmasRunes := xmasAsRunes()
	for i := 1; i < len(xmasRunes); i++ {
		letter := string(xmasRunes[i])
		if !checkStringValueInGrid(grid, row-i, col-i, letter) {
			return false
		}
	}
	return true
}

func xmasAsRunes() []rune {
	return []rune("XMAS")
}

func checkStringValueInGrid(grid [][]rune, row int, col int, s string) bool {
	found := getStringFromGrid(grid, row, col)
	return found == s
}

func getStringFromGrid(grid [][]rune, row int, col int) string {
	if row >= len(grid) || row < 0 || col >= len(grid[row]) || col < 0 {
		return ""
	}

	return string(grid[row][col])
}

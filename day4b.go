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
			if isValidXMAS(grid, row, col) {
				found++
			}
		}
	}

	return found
}

func isValidXMAS(grid [][]rune, row int, col int) bool {
	c := getStringFromGrid(grid, row, col)
	nw := getStringFromGrid(grid, row-1, col-1)
	sw := getStringFromGrid(grid, row+1, col-1)
	ne := getStringFromGrid(grid, row-1, col+1)
	se := getStringFromGrid(grid, row+1, col+1)

	word1 := nw + c + se
	word2 := sw + c + ne

	isValid1 := word1 == "SAM" || word1 == "MAS"
	isValid2 := word2 == "SAM" || word2 == "MAS"

	return isValid1 && isValid2
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

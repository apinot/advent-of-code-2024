package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

func main() {
	input := readInput()
	antennas, frequencies, gi := parseInput(input)

	result := doLogic(antennas, frequencies, gi)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day8.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

type Antenna struct {
	frequency string
	row       int
	col       int
}

type GridInfo struct {
	maxRow int
	maxCol int
}

func parseInput(input string) (map[string][]Antenna, []string, GridInfo) {
	antennasMap := make(map[string][]Antenna)
	var frequencies []string

	rows := strings.Split(input, "\n")

	maxRow := len(rows)
	maxCol := 0
	for row := 0; row < maxRow; row++ {
		cols := []rune(rows[row])
		maxCol = len(cols)
		for col := 0; col < maxCol; col++ {
			value := string(cols[col])
			if value == "." {
				continue
			}

			a := Antenna{
				frequency: value,
				row:       row,
				col:       col,
			}

			antennasMap[value] = append(antennasMap[value], a)
			if !slices.Contains(frequencies, value) {
				frequencies = append(frequencies, value)
			}
		}
	}

	gi := GridInfo{maxRow, maxCol}

	return antennasMap, frequencies, gi
}

type Antinode struct {
	row int
	col int
}

func doLogic(antennas map[string][]Antenna, frequencies []string, gridInfo GridInfo) int {
	antinodes := computeAntinodes(antennas, frequencies, gridInfo)
	return countDinstinct(antinodes)
}

func computeAntinodes(antennas map[string][]Antenna, frequencies []string, gi GridInfo) []Antinode {
	var antinodes []Antinode

	for _, f := range frequencies {
		aOfF := antennas[f]
		nbForF := len(aOfF)

		for i := 0; i < nbForF; i++ {
			for j := 0; j < nbForF; j++ {
				if i == j {
					continue
				}

				aBase := aOfF[i]
				aComp := aOfF[j]

				diffRow := aBase.row - aComp.row
				diffCol := aBase.col - aComp.col

				antinodeRow := aBase.row + diffRow
				antinodeCol := aBase.col + diffCol
				if !isInBound(antinodeRow, antinodeCol, gi) {
					continue
				}
				antinodes = append(antinodes, Antinode{row: antinodeRow, col: antinodeCol})
			}
		}
	}
	return antinodes
}

func countDinstinct(antinodes []Antinode) int {
	mapColsPerRow := make(map[int][]int)

	for _, a := range antinodes {
		r := a.row
		c := a.col

		if slices.Contains(mapColsPerRow[r], c) {
			continue
		}

		mapColsPerRow[r] = append(mapColsPerRow[r], c)
	}

	nb := 0
	for _, v := range mapColsPerRow {
		nb += len(v)
	}
	return nb
}

func isInBound(r int, c int, gi GridInfo) bool {
	return r >= 0 && c >= 0 && r < gi.maxRow && c < gi.maxCol
}

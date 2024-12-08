package main

import (
	"fmt"
	"os"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

func main() {
	input := readInput()
	parsed := parseInput(input)

	result := doLogic(parsed)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day0.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func parseInput(input string) string {
	return input
}

func doLogic(input string) string {
	return input
}

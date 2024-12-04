package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	firstList, secondList := parseInput(input)

	result := doLogic(firstList, secondList)

	fmt.Printf("%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day1.txt")
    if err != nil {
        fmt.Print(err)
    }

    return string(b)
}

func parseInput(input string) ([]int, []int) {
	var firstElements []int
	var secondElements []int

	lines := strings.Split(input, "\r\n")
	for _, line := range lines  {
		elements := strings.Split(line, "   ")
		firstElement, _ :=  strconv.Atoi(elements[0])
		secondElement, _ := strconv.Atoi(elements[1])

		firstElements = append(firstElements, firstElement)
		secondElements = append(secondElements, secondElement)
	}

	return firstElements, secondElements
}

func doLogic(firstList []int, secondList []int) int {
	similarityScore := 0

	for _, element := range firstList {
		appearInSecondList := 0
		for _, comparedElement := range secondList {
			if element == comparedElement {
				appearInSecondList++
			}
		}

		similarityScore += (element * appearInSecondList)
	}

	return similarityScore
}

func intAbs(x int) int {
	if (x >= 0) {
		return x;
	}
	return x * -1;
} 

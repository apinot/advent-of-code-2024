package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	reports := parseInput(input)

	result := doLogic(reports)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day2.txt")
    if err != nil {
        fmt.Print(err)
    }

    return string(b)
}

func parseInput(input string) [][]int {
	var reports [][]int

	lines := strings.Split(input, "\r\n")
	for _, line := range lines  {
		elements := strings.Split(line, " ")

		var report []int
		for _, element := range elements {
			elementAsInt, _ := strconv.Atoi(element)
			report = append(report, elementAsInt)
		}

		reports = append(reports, report)
	}

	return reports
}

func doLogic(reports [][]int) int {
	nbValidReports := 0
	for _, report := range reports {
		if isReportValid(report) {
			nbValidReports++
		}
	}
	return nbValidReports;
}

func isReportValid(report []int) bool {
	diffs := computeDiffs(report);
	return areValid(diffs)
}

func computeDiffs(report []int) []int {
	var diffs []int
	for i := 1; i < len(report) ; i++ {
		diff := report[i - 1] - report[i]
		diffs = append(diffs, diff)
	} 
	return diffs
}

func areValid(diffs []int) bool {
	previousValue := diffs[0]
	if (!isValueValid(previousValue)) {
		return false
	}

	previousIsPositive := previousValue > 0
	for i := 1; i < len(diffs) ; i++ {
		currentValue := diffs[i]
		if !isValueValid(currentValue) {
			return false
		}

		currentIsPositive := currentValue > 0
		if currentIsPositive != previousIsPositive {
			return false
		}

		previousIsPositive = currentIsPositive
	}
	return true
}

func isValueValid(value int) bool {
	return (value >= -3 && value <= -1) || (value >= 1 && value <= 3)
}
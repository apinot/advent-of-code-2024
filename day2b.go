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
	diffs := computeDiffs(report, -1);
	if areValid(diffs) {
		return true
	}

	for i := 0; i < len(report); i++ {
		diffsWithSkipped := computeDiffs(report, i)
		if areValid(diffsWithSkipped) {
			return true
		}
	}

	return false
}

func computeDiffs(report []int, skipIndex int) []int {
	var diffs []int

	var previousValue int
	var i int
	if skipIndex == 0 {
		previousValue = report[1]
		i = 2
	} else {
		previousValue = report[0]
		i = 1
	}

	for ; i < len(report); i++ {
		if i == skipIndex {
			continue
		}

		currentValue :=  report[i]
		
		diff := previousValue - currentValue
		diffs = append(diffs, diff) 
		previousValue = currentValue
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
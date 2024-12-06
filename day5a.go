package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

func main() {
	input := readInput()
	pagesOrders, updates := parseInput(input)

	result := doLogic(pagesOrders, updates)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day5.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

type PageComparison struct {
	before int
	after  int
}

type ManualUpdate struct {
	pages []int
}

func parseInput(input string) ([]PageComparison, []ManualUpdate) {
	var pageComparisons []PageComparison
	var manualUpdates []ManualUpdate

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if strings.Contains(line, "|") {
			pages := strings.Split(line, "|")
			pageBefore, _ := strconv.Atoi(pages[0])
			pageAfter, _ := strconv.Atoi(pages[1])

			pageComparisons = append(pageComparisons, PageComparison{pageBefore, pageAfter})
			continue
		}

		pages := strings.Split(line, ",")
		if len(pages) < 2 {
			continue
		}

		var pagesAsInt []int
		for _, page := range pages {
			pageAsInt, _ := strconv.Atoi(page)
			pagesAsInt = append(pagesAsInt, pageAsInt)
		}

		manualUpdates = append(manualUpdates, ManualUpdate{pagesAsInt})
	}

	return pageComparisons, manualUpdates
}

func doLogic(pageComparisons []PageComparison, manualUpdates []ManualUpdate) int {
	pagesBeforeMap := getPagesBeforeMap(pageComparisons)
	validsManualUpdates := keepValidManualUpdates(manualUpdates, pagesBeforeMap)
	return sumMiddlePages(validsManualUpdates)
}

func getPagesBeforeMap(pageComparisons []PageComparison) map[int][]int {
	pagesMap := make(map[int][]int)
	for _, pc := range pageComparisons {
		before := pc.before
		after := pc.after

		pagesMap[before] = append(pagesMap[before], after)
	}
	return pagesMap
}

func keepValidManualUpdates(manualsUpdates []ManualUpdate, pagesBeforeMaps map[int][]int) []ManualUpdate {
	var valids []ManualUpdate
	for _, mu := range manualsUpdates {
		if isValidManualUpdate(mu, pagesBeforeMaps) {
			valids = append(valids, mu)
		}
	}

	return valids
}

func isValidManualUpdate(manualUpdate ManualUpdate, pagesBeforeMaps map[int][]int) bool {
	pages := manualUpdate.pages

	firstPage := pages[0]
	nextPages := pages[1:]
	return isCorrectlyPlaced(firstPage, nextPages, pagesBeforeMaps)
}

func isCorrectlyPlaced(page int, nextPages []int, pagesBeforeMaps map[int][]int) bool {
	nextPagesCount := len(nextPages)
	if nextPagesCount < 1 {
		return true
	}

	if nextPagesCount == 1 {
		return isBefore(page, nextPages[0], pagesBeforeMaps)
	}

	pageToCompare := nextPages[0]
	if !isBefore(page, pageToCompare, pagesBeforeMaps) {
		return false
	}
	return isCorrectlyPlaced(pageToCompare, nextPages[1:], pagesBeforeMaps)
}

func isBefore(page int, pageRequiredAfter int, pagesBeforeMaps map[int][]int) bool {
	pagesAfterAfter := pagesBeforeMaps[pageRequiredAfter]
	return !slices.Contains(pagesAfterAfter, page)
}

func sumMiddlePages(manualUpdates []ManualUpdate) int {
	total := 0
	for _, mu := range manualUpdates {
		total += getMiddlePage(mu)
	}
	return total
}

func getMiddlePage(mu ManualUpdate) int {
	nbPages := len(mu.pages)
	middlePage := nbPages / 2
	dd(mu.pages[middlePage])
	return mu.pages[middlePage]
}

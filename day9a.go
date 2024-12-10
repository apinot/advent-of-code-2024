package main

import (
	"fmt"
	"os"
	"strconv"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

func main() {
	input := readInput()
	numbers := parseInput(input)
	checksum := doLogic(numbers)

	fmt.Printf("\n\n%d\n", checksum)
}

func readInput() string {
	b, err := os.ReadFile("./day9.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func parseInput(input string) []int {
	nAsRunes := []rune(input)

	var numbers []int
	for _, r := range nAsRunes {
		nAsInt, _ := strconv.Atoi(string(r))
		numbers = append(numbers, nAsInt)
	}
	return numbers
}

func doLogic(input []int) int {
	blocks := getDiskBlocs(input)
	reorganizeBlocks(blocks)
	return getBlocksCheckSum(blocks)
}

func getDiskBlocs(diskMap []int) []string {
	var blocks []string

	fileId := 0
	for index, blockSize := range diskMap {
		blockVal := "."
		if !isFreeSpace(index) {
			blockVal = strconv.Itoa(fileId)
			fileId++
		}

		for b := 0; b < blockSize; b++ {
			blocks = append(blocks, blockVal)
		}
	}
	return blocks
}

func isFreeSpace(mapIndex int) bool {
	return mapIndex%2 == 1
}

func trimEndEmptyBloc(blocks []string) []string {
	lastIndex := len(blocks) - 1
	if blocks[lastIndex] != "." {
		return blocks
	}
	return trimEndEmptyBloc(blocks[0:lastIndex])
}

func reorganizeBlocks(rawBlocks []string) {
	blocks := trimEndEmptyBloc(rawBlocks)

	lastIndex := len(blocks) - 1
	lastBlock := blocks[lastIndex]
	for index, blockId := range blocks {
		if blockId != "." {
			continue
		}

		blocks[index] = lastBlock
		blocks[lastIndex] = "."
		reorganizeBlocks(blocks)
		return
	}
}

func getBlocksCheckSum(blocks []string) int {
	checksum := 0
	for index, blockId := range blocks {
		if blockId == "." {
			continue
		}

		blockValue, _ := strconv.Atoi(blockId)
		checkPos := blockValue * index
		checksum = checksum + checkPos
	}
	return checksum
}

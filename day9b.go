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
	blocks = trimEndEmptyBloc(blocks)
	blocksNb := len(blocks)
	reorganizeBlocks(blocks, blocks[blocksNb-1])
	return getBlocksCheckSum(blocks)
}

func getDiskBlocs(diskMap []int) []string {
	var blocks []string

	fileId := 0
	for index, blockSize := range diskMap {
		blockVal := "."
		if index%2 == 0 {
			blockVal = strconv.Itoa(fileId)
			fileId++
		}

		for b := 0; b < blockSize; b++ {
			blocks = append(blocks, blockVal)
		}
	}
	return blocks
}

func trimEndEmptyBloc(blocks []string) []string {
	lastIndex := len(blocks) - 1
	if blocks[lastIndex] != "." {
		return blocks
	}
	return trimEndEmptyBloc(blocks[0:lastIndex])
}

func reorganizeBlocks(rawBlocks []string, blockId string) {
	blockIdAsInt, _ := strconv.Atoi(blockId)
	if blockIdAsInt < 0 {
		return
	}

	nextBlockId := strconv.Itoa(blockIdAsInt - 1)

	blocks := trimEndEmptyBloc(rawBlocks)

	blockIndex := len(blocks) - 1
	for blockIndex >= 0 {
		if blocks[blockIndex] != blockId {
			blockIndex--
			continue
		}

		blockSpace := getBlockSpaceByEnd(blocks, blockIndex)

		freeIndex := 0
		for freeIndex < blockIndex {
			if blocks[freeIndex] != "." {
				freeIndex++
				continue
			}

			freeSpace := getBlockSpaceByStart(blocks, freeIndex)

			if blockSpace > freeSpace {
				freeIndex += freeSpace
				continue
			}
			for ri := 0; ri < blockSpace; ri++ {
				replaceFreeIndex := freeIndex + ri
				replaceBlockIndex := blockIndex - ri

				blocks[replaceFreeIndex] = blocks[replaceBlockIndex]
				blocks[replaceBlockIndex] = "."
			}

			reorganizeBlocks(blocks, nextBlockId)
			return
		}

		blockIndex -= blockSpace
	}
	reorganizeBlocks(blocks, nextBlockId)
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

func getBlockSpaceByStart(blocks []string, pos int) int {
	blockId := blocks[pos]
	i := pos
	for i < len(blocks) {
		if blockId == blocks[i] {
			i++
			continue
		}
		return i - pos
	}

	return 1
}

func getBlockSpaceByEnd(blocks []string, pos int) int {
	blockId := blocks[pos]
	i := pos
	for i >= 0 {
		if blockId == blocks[i] {
			i--
			continue
		}
		return pos - i
	}

	return 1
}

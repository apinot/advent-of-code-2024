package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	muls := parseInput(input)

	result := doLogic(muls)

	fmt.Printf("\n\n%d\n", result)
}

func readInput() string {
	b, err := os.ReadFile("./day3.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

type Mul struct {
	x int
	y int
}

func parseInput(input string) []Mul {
	var muls []Mul

	r := createMulRegexp()
	matched := r.FindAllString(input, -1)

	for _, mulString := range matched {
		trimmedMulString := strings.Trim(mulString, "mul(")
		trimmedMulString = strings.Trim(trimmedMulString, ")")

		x, y, _ := strings.Cut(trimmedMulString, ",")

		muls = append(muls, createMulFromStrings(x, y))
	}

	return muls
}

func createMulRegexp() *regexp.Regexp {
	var regexpBuilder strings.Builder
	regexpBuilder.WriteString(regexp.QuoteMeta("mul("))
	regexpBuilder.WriteString("\\d+")
	regexpBuilder.WriteString(regexp.QuoteMeta(","))
	regexpBuilder.WriteString("\\d+")
	regexpBuilder.WriteString(regexp.QuoteMeta(")"))

	r, _ := regexp.Compile(regexpBuilder.String())

	return r
}

func createMulFromStrings(x string, y string) Mul {
	xAsInt, _ := strconv.Atoi(x)
	yAsInt, _ := strconv.Atoi(y)
	return Mul{
		x: xAsInt,
		y: yAsInt,
	}
}

func doLogic(muls []Mul) int {
	fmt.Printf("%v", muls)
	result := 0
	for _, mul := range muls {
		result = result + (mul.x * mul.y)
	}
	return result
}

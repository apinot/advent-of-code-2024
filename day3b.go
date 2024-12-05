package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func dd(v any) {
	fmt.Printf("%v\n", v)
}

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
	input = removeDisabledInstructions(input)
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

func removeDisabledInstructions(input string) string {
	dontRegexp := createDontRegexp()
	doRegexp := createDoRegexp()

	disabledAtMatch := dontRegexp.FindStringIndex(input)
	if disabledAtMatch == nil {
		return input
	}
	disabledAt := disabledAtMatch[0]

	inputAfterDont := deleteInString(input, 0, disabledAt)
	enabledAtMatch := doRegexp.FindStringIndex(inputAfterDont)
	if enabledAtMatch == nil {
		return deleteInString(input, disabledAt, -1)
	}
	enabledAt := disabledAt + enabledAtMatch[1]

	return removeDisabledInstructions(deleteInString(input, disabledAt, enabledAt))
}

func deleteInString(s string, start int, end int) string {
	fmt.Printf("%d -> %d \n", start, end)
	r := []rune(s)
	if end < 0 {
		return string(r[0:start])
	}

	return string(append(r[0:start], r[end+1:]...))
}

func createDontRegexp() *regexp.Regexp {
	var regexpBuilder strings.Builder
	regexpBuilder.WriteString(regexp.QuoteMeta("don't()"))

	r, _ := regexp.Compile(regexpBuilder.String())

	return r
}

func createDoRegexp() *regexp.Regexp {
	var regexpBuilder strings.Builder
	regexpBuilder.WriteString(regexp.QuoteMeta("do()"))
	r, _ := regexp.Compile(regexpBuilder.String())

	return r
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
	result := 0
	for _, mul := range muls {
		result = result + (mul.x * mul.y)
	}
	return result
}

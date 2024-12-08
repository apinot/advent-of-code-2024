package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

type Equation struct {
	result  float64
	numbers []float64
}

func readInput() string {
	b, err := os.ReadFile("./day7.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}

func parseInput(input string) []Equation {
	var equations []Equation

	rows := strings.Split(input, "\n")
	for _, r := range rows {
		splitByColon := strings.Split(r, ":")
		totalAsInt, _ := strconv.ParseFloat(splitByColon[0], 64)

		var numbersAsInt []float64
		operands := strings.Split(strings.Trim(splitByColon[1], " "), " ")
		for _, o := range operands {
			nASInt, _ := strconv.ParseFloat(o, 64)
			numbersAsInt = append(numbersAsInt, nASInt)
		}

		equations = append(equations, Equation{totalAsInt, numbersAsInt})
	}
	return equations
}

func doLogic(equations []Equation) int {
	sumsOfValid := 0

	for _, eq := range equations {
		if isValidEquation(eq) {
			sumsOfValid += int(eq.result)
		}
	}
	// i := 1
	// fmt.Printf("%v is %b\n-----------------------\n", equations[i], isValidEquation(equations[i]))

	return sumsOfValid
}

func isValidEquation(eq Equation) bool {
	nbNumbers := len(eq.numbers)
	if nbNumbers == 0 {
		// dd("--------")
		// dd("zero")
		// dd(eq.result)
		// dd("----------")
		return eq.result == 0
	}

	if nbNumbers == 1 {
		// dd("--------")
		// dd("one")
		// dd(eq.result)
		// dd(eq.numbers[0])
		// dd("----------")
		return eq.result == eq.numbers[0]
	}

	if nbNumbers == 2 {
		a := eq.numbers[0]
		b := eq.numbers[1]
		r := eq.result

		// dd("--------")
		// dd("two")
		// dd("r")
		// dd(r)
		// dd("a")
		// dd(a)
		// dd("b")
		// dd(b)
		// dd("a+b and a*b")
		// dd(a + b)
		// dd(a * b)
		// dd("----------")
		return a+b == r || a*b == r
	}

	lastIndex := len(eq.numbers) - 1
	z := eq.numbers[lastIndex]

	rSub := eq.result - z
	rDiv := eq.result / z

	newEqNumbers := eq.numbers[0:lastIndex]
	subEq := Equation{rSub, newEqNumbers}
	divEq := Equation{rDiv, newEqNumbers}

	// dd("====================")
	// dd("more")
	// dd("lastIndex")
	// dd(lastIndex)
	// dd("z")
	// dd(z)
	// dd("rSub")
	// dd(rSub)
	// dd("rDvi")
	// dd(rDiv)
	// dd("  recursive")
	// dd(subEq)
	// dd(isValidEquation(subEq))
	// dd(divEq)
	// dd(isValidEquation(divEq))
	// dd("=======================")

	return isValidEquation(subEq) || isValidEquation(divEq)
}

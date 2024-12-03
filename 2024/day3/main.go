package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInput() string {
	buf, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func evaluateMulExpressions(expressions string) int {
	regex := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches := regex.FindAllStringSubmatch(expressions, -1)

	runningTotal := 0
	for i := 0; i < len(matches); i++ {
		match := matches[i]
		num1, err1 := strconv.Atoi(match[1])
		num2, err2 := strconv.Atoi(match[2])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		runningTotal += num1 * num2
	}

	return runningTotal
}

func solution1() {
	input := getInput()

	total := evaluateMulExpressions(input)

	fmt.Printf("Part 1 Solution: %d\n", total)
}

func solution2() {
	runningTotal := 0
	input, before, after := getInput(), "", ""
	found := true
	for found {
		before, after, _ = strings.Cut(input, "don't()")
		runningTotal += evaluateMulExpressions(before)
		_, input, found = strings.Cut(after, "do()")
	}

	fmt.Printf("Part 2 Solution: %d\n", runningTotal)
}

func main() {
	solution1()
	solution2()
}

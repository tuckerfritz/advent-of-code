package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func solution1() {
	buf, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	regex := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches := regex.FindAllStringSubmatch(string(buf), -1)

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

	fmt.Printf("Part 1 Solution: %d\n", runningTotal)
}

func main() {
	solution1()
}

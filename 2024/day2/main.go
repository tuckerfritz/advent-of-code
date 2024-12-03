package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMatrix() [][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]int

	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		row := make([]int, len(numbers))
		for i, v := range numbers {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row[i] = num
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func isIntArraySafe(row []int) bool {
	maxDelta := -1
	ascending, descending := true, true
	for j := 1; j < len(row); j++ {
		if row[j] <= row[j-1] {
			ascending = false
		}
		if row[j] >= row[j-1] {
			descending = false
		}
		currentDelta := abs(row[j] - row[j-1])
		if currentDelta > maxDelta {
			maxDelta = currentDelta
		}
	}
	return (ascending || descending) && (maxDelta > 0 && maxDelta <= 3)
}

func removeIndex(slice []int, index int) []int {
	newArray := make([]int, 0, len(slice))
	newArray = append(newArray, slice[:index]...)
	return append(newArray, slice[index+1:]...)
}

func solution1() {
	matrix := getMatrix()
	numSafe := 0

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		if isIntArraySafe(row) {
			numSafe += 1
		}
	}

	fmt.Printf("Part 1: Number of safe reports: %d\n", numSafe)
}

func solution2() {
	matrix := getMatrix()
	numSafe := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if isIntArraySafe(removeIndex(matrix[i], j)) {
				numSafe += 1
				break
			}
		}
	}

	fmt.Printf("Part 2: Number of safe reports: %d\n", numSafe)
}

func main() {
	solution1()
	solution2()
}

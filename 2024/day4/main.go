package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := make([][]string, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		matrix = append(matrix, letters)
	}

	return matrix
}

func checkLeftRight(matrix [][]string, i int, j int) bool {
	if j+3 >= len(matrix[i]) {
		return false
	}
	return matrix[i][j] == "X" && matrix[i][j+1] == "M" &&
		matrix[i][j+2] == "A" && matrix[i][j+3] == "S"
}

func checkRightLeft(matrix [][]string, i int, j int) bool {
	if j-3 < 0 {
		return false
	}
	return matrix[i][j] == "X" && matrix[i][j-1] == "M" &&
		matrix[i][j-2] == "A" && matrix[i][j-3] == "S"
}

func checkBottomTop(matrix [][]string, i int, j int) bool {
	if i-3 < 0 {
		return false
	}
	return matrix[i][j] == "X" && matrix[i-1][j] == "M" &&
		matrix[i-2][j] == "A" && matrix[i-3][j] == "S"
}

func checkTopBottom(matrix [][]string, i int, j int) bool {
	if i+3 >= len(matrix) {
		return false
	}
	return matrix[i][j] == "X" && matrix[i+1][j] == "M" &&
		matrix[i+2][j] == "A" && matrix[i+3][j] == "S"
}

func checkDiagonalTowardsTopRight(matrix [][]string, i int, j int) bool {
	if i+3 >= len(matrix) || j-3 < 0 {
		return false
	}
	return matrix[i][j] == "X" && matrix[i+1][j-1] == "M" &&
		matrix[i+2][j-2] == "A" && matrix[i+3][j-3] == "S"
}

func checkDiagonalTowardsBottomRight(matrix [][]string, i int, j int) bool {
	if i+3 >= len(matrix) || j+3 >= len(matrix[i]) {
		return false
	}
	return matrix[i][j] == "X" && matrix[i+1][j+1] == "M" &&
		matrix[i+2][j+2] == "A" && matrix[i+3][j+3] == "S"
}

func checkDiagonalTowardsTopLeft(matrix [][]string, i int, j int) bool {
	if i-3 < 0 || j-3 < 0 {
		return false
	}
	return matrix[i][j] == "X" && matrix[i-1][j-1] == "M" &&
		matrix[i-2][j-2] == "A" && matrix[i-3][j-3] == "S"
}

func checkDiagonalTowardsBottomLeft(matrix [][]string, i int, j int) bool {
	if i-3 < 0 || j+3 >= len(matrix[i]) {
		return false
	}
	return matrix[i][j] == "X" && matrix[i-1][j+1] == "M" &&
		matrix[i-2][j+2] == "A" && matrix[i-3][j+3] == "S"
}

func solution1() {
	matrix := getInput()
	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if checkLeftRight(matrix, i, j) {
				total += 1
			}
			if checkRightLeft(matrix, i, j) {
				total += 1
			}
			if checkBottomTop(matrix, i, j) {
				total += 1
			}
			if checkTopBottom(matrix, i, j) {
				total += 1
			}
			if checkDiagonalTowardsTopRight(matrix, i, j) {
				total += 1
			}
			if checkDiagonalTowardsBottomRight(matrix, i, j) {
				total += 1
			}
			if checkDiagonalTowardsTopLeft(matrix, i, j) {
				total += 1
			}
			if checkDiagonalTowardsBottomLeft(matrix, i, j) {
				total += 1
			}
		}
	}
	fmt.Printf("Part 1 Solution: %d\n", total)
}

func xmasOutOfBounds(matrix [][]string, i int, j int) bool {
	if i-1 < 0 || i+1 >= len(matrix) ||
		j-1 < 0 || j+1 >= len(matrix[i]) {
		return true
	}
	return false
}

func checkXmasBottomLeftAndTopRight(matrix [][]string, i int, j int) bool {
	if (matrix[i-1][j+1] == "S" && matrix[i+1][j-1] == "M") ||
		(matrix[i-1][j+1] == "M" && matrix[i+1][j-1] == "S") {
		return true
	}
	return false
}

func checkXmasTopLeftAndBottomRight(matrix [][]string, i int, j int) bool {
	if (matrix[i-1][j-1] == "S" && matrix[i+1][j+1] == "M") ||
		(matrix[i-1][j-1] == "M" && matrix[i+1][j+1] == "S") {
		return true
	}
	return false
}
func isXmas(matrix [][]string, i int, j int) bool {
	return matrix[i][j] == "A" && checkXmasBottomLeftAndTopRight(matrix, i, j) &&
		checkXmasTopLeftAndBottomRight(matrix, i, j)
}

func solution2() {
	matrix := getInput()
	total := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if xmasOutOfBounds(matrix, i, j) {
				continue
			}
			if isXmas(matrix, i, j) {
				total += 1
			}
		}
	}

	fmt.Printf("Part 2 Solution: %d\n", total)
}

func main() {
	solution1()
	solution2()
}

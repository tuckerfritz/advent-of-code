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

	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)

	for scanner.Scan() {
		positions := strings.Split(scanner.Text(), "")
		grid = append(grid, positions)
	}

	return grid
}

func getDirections() map[string][2]int {
	return map[string][2]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}
}

func getInitialCoords(grid [][]string) [2]int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			char := grid[i][j]
			if char == "^" || char == ">" || char == "v" || char == "<" {
				return [2]int{i, j}
			}
		}
	}
	panic("Could not find starting position!")
}

func processRoute(grid [][]string, directions map[string][2]int, x int, y int, numPositions int) int {
	guard := grid[x][y]
	direction := directions[guard]
	nextX, nextY := direction[0]+x, direction[1]+y
	if nextX > len(grid)-1 || nextX < 0 || nextY > len(grid[x])-1 || nextY < 0 {
		return numPositions + 1
	}
	if grid[nextX][nextY] == "#" {
		if guard == ">" {
			grid[x][y] = "v"
		} else if guard == "v" {
			grid[x][y] = "<"
		} else if guard == "<" {
			grid[x][y] = "^"
		} else if guard == "^" {
			grid[x][y] = ">"
		}
		return processRoute(grid, directions, x, y, numPositions)
	} else if grid[nextX][nextY] == "X" {
		grid[x][y] = "X"
		grid[nextX][nextY] = guard
		return processRoute(grid, directions, nextX, nextY, numPositions)
	} else if grid[nextX][nextY] == "." {
		grid[x][y] = "X"
		grid[nextX][nextY] = guard
		return processRoute(grid, directions, nextX, nextY, numPositions+1)
	}
	panic("Unknown character in input!")
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}

func solution1() {
	grid := getInput()
	directions := getDirections()
	coords := getInitialCoords(grid)
	numPositions := processRoute(grid, directions, coords[0], coords[1], 0)
	printGrid(grid)
	fmt.Printf("Part 1 Solution: %d\n", numPositions)
}

func main() {
	solution1()
}

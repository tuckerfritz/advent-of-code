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

func isLoop(visited map[[2]int][]string, guard string, x int, y int) bool {
	coords := [2]int{x, y}
	guardPermutations, exists := visited[coords]
	if !exists {
		visited[coords] = []string{guard}
		return false
	}
	for i := 0; i < len(guardPermutations); i++ {
		if guard == guardPermutations[i] {
			return true
		}
	}
	visited[coords] = append(visited[coords], guard)
	return false
}

func processRoute(grid [][]string, visited map[[2]int][]string, directions map[string][2]int, x int, y int, numPositions int) (map[[2]int][]string, int, bool) {
	guard := grid[x][y]
	direction := directions[guard]
	nextX, nextY := direction[0]+x, direction[1]+y
	if isLoop(visited, guard, x, y) {
		return visited, numPositions, true
	}
	if nextX > len(grid)-1 || nextX < 0 || nextY > len(grid[x])-1 || nextY < 0 {
		return visited, numPositions + 1, false
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
		return processRoute(grid, visited, directions, x, y, numPositions)
	} else if grid[nextX][nextY] == "X" {
		grid[x][y] = "X"
		grid[nextX][nextY] = guard
		return processRoute(grid, visited, directions, nextX, nextY, numPositions)
	} else if grid[nextX][nextY] == "." {
		grid[x][y] = "X"
		grid[nextX][nextY] = guard
		return processRoute(grid, visited, directions, nextX, nextY, numPositions+1)
	}
	panic("Unknown character in input!")
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}

func solution1() map[[2]int][]string {
	grid := getInput()
	directions := getDirections()
	coords := getInitialCoords(grid)
	visited, numPositions, _ := processRoute(grid, make(map[[2]int][]string), directions, coords[0], coords[1], 0)
	fmt.Printf("Part 1 Solution: %d\n", numPositions)
	return visited
}

func copyGrid(grid [][]string) [][]string {
	duplicate := make([][]string, len(grid))
	for i := range grid {
		duplicate[i] = make([]string, len(grid[i]))
		copy(duplicate[i], grid[i])
	}
	return duplicate
}

func solution2(visited map[[2]int][]string) {
	originalGrid := getInput()
	directions := getDirections()
	coords := getInitialCoords(originalGrid)
	totalObstaclePositions := 0

	for visitedCoords := range visited {
		grid := copyGrid(originalGrid)
		if coords[0] == visitedCoords[0] && coords[1] == visitedCoords[1] {
			continue
		}
		grid[visitedCoords[0]][visitedCoords[1]] = "#"
		_, _, isLoop := processRoute(grid, make(map[[2]int][]string), directions, coords[0], coords[1], 0)
		if isLoop {
			totalObstaclePositions += 1
		}
		grid[visitedCoords[0]][visitedCoords[1]] = "."
	}
	fmt.Printf("Part 2 Solution: %d\n", totalObstaclePositions)
}

func main() {
	visited := solution1()
	solution2(visited)
}

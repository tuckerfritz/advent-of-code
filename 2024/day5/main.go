package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRulesAndInputLists() (map[int]map[int]struct{}, [][]int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ruleSet = make(map[int]map[int]struct{})
	var listsOfPages = make([][]int, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		if line == "" {
			break
		}
		var numbers = strings.Split(line, "|")
		page1, err1 := strconv.Atoi(numbers[0])
		page2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}

		set, exists := ruleSet[page1]
		if exists {
			set[page2] = struct{}{}
		} else {
			newSet := make(map[int]struct{})
			newSet[page2] = struct{}{}
			ruleSet[page1] = newSet
		}

	}

	for scanner.Scan() {
		var numbers = strings.Split(scanner.Text(), ",")
		var pages = make([]int, 0)
		for i := 0; i < len(numbers); i++ {
			number, err := strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			}
			pages = append(pages, number)
		}
		listsOfPages = append(listsOfPages, pages)
	}

	return ruleSet, listsOfPages
}

func getMiddleNumberOfList(list []int) int {
	if len(list)%2 != 0 {
		return list[len(list)/2]
	} else {
		index1 := len(list) / 2
		index2 := (len(list) / 2) + 1
		return (list[index1] + list[index2]) / 2
	}
}

func solution1() {
	ruleSet, listsOfPages := getRulesAndInputLists()
	total := 0

	for i := 0; i < len(listsOfPages); i++ {
		isValidList := true
		for j := len(listsOfPages[i]) - 1; j >= 0; j-- {
			currentPage := listsOfPages[i][j]
			for k := j - 1; k >= 0; k-- {
				pageToCheck := listsOfPages[i][k]
				_, exists := ruleSet[currentPage][pageToCheck]
				if exists {
					isValidList = false
					break
				}
			}
			if !isValidList {
				break
			}
		}
		if isValidList {
			total += getMiddleNumberOfList(listsOfPages[i])
		}
	}

	fmt.Printf("Part 1 Solution: %d\n", total)

}

func main() {
	solution1()
}
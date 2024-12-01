package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		var numbers []string
		numbers = strings.Fields(scanner.Text())
		int1, err1 := strconv.Atoi(numbers[0])
		int2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}
		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}
	return list1, list2
}

func solution1() {
	list1, list2 := getLists()
	sort.Ints(list1)
	sort.Ints(list2)
	totalDistance := 0
	for i := 0; i < len(list1) && i < len(list2); i++ {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	fmt.Printf("Part 1 Answer: %d\n", totalDistance)
}

func solution2() {
	list2Occurences := make(map[int]int)
	list1, list2 := getLists()
	for _, v := range list2 {
		list2Occurences[v] = list2Occurences[v] + 1
	}
	similarityScore := 0
	for _, v := range list1 {
		similarityScore += v * list2Occurences[v]
	}
	fmt.Printf("Part 2 Answer: %d\n", similarityScore)
}

func main() {
	solution1()
	solution2()
}

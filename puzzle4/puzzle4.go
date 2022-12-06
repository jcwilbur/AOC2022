package main

import (
	"AOC2022/helpers"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	fourB()
}

func fourA() {
	inputStrings := helpers.ReadFileAsStrings("input4.txt")
	fmt.Println(inputStrings)

	pairs := make([][2][]int, len(inputStrings))
	for i, v := range inputStrings {
		elfPair := strings.Split(v, ",")
		for j, nums := range elfPair {
			indices := strings.Split(nums, "-")
			fmt.Println(indices)
			min, err := strconv.Atoi(indices[0])
			if err != nil {
				panic("error!")
			}
			max, err := strconv.Atoi(indices[1])
			if err != nil {
				panic("Other error!")
			}
			fullRange := makeRange(min, max)
			pairs[i][j] = fullRange
		}
	}

	countOfContains := countContains(pairs)

	fmt.Println(countOfContains)

}

func countContains(pairs [][2][]int) int {
	internalCounter := 0
	for _, pair := range pairs {
		var smallerRange []int
		var largerRange []int
		if len(pair[0]) < len(pair[1]) {
			smallerRange = pair[0]
			largerRange = pair[1]
		} else {
			smallerRange = pair[1]
			largerRange = pair[0]
		}

		smallerIsContainedWithinLarger := true

		for _, value := range smallerRange {
			if !slices.Contains(largerRange, value) {
				smallerIsContainedWithinLarger = false
				break
			}
		}
		if smallerIsContainedWithinLarger {
			internalCounter++
		}
	}
	return internalCounter
}

func fourB() {
	inputStrings := helpers.ReadFileAsStrings("input4.txt")
	fmt.Println(inputStrings)

	pairs := make([][2][]int, len(inputStrings))
	for i, v := range inputStrings {
		elfPair := strings.Split(v, ",")
		for j, nums := range elfPair {
			indices := strings.Split(nums, "-")
			fmt.Println(indices)
			min, err := strconv.Atoi(indices[0])
			if err != nil {
				panic("error!")
			}
			max, err := strconv.Atoi(indices[1])
			if err != nil {
				panic("Other error!")
			}
			fullRange := makeRange(min, max)
			pairs[i][j] = fullRange
		}
	}

	countOfContains := countOverlaps(pairs)

	fmt.Println(countOfContains)

}

func countOverlaps(pairs [][2][]int) int {
	internalCounter := 0
	for _, pair := range pairs {
		var smallerRange []int
		var largerRange []int
		if len(pair[0]) < len(pair[1]) {
			smallerRange = pair[0]
			largerRange = pair[1]
		} else {
			smallerRange = pair[1]
			largerRange = pair[0]
		}

		for _, value := range smallerRange {
			if slices.Contains(largerRange, value) {
				internalCounter++
				break
			}
		}

	}
	return internalCounter
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

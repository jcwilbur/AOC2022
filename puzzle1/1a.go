package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var elves []int
	elves = append(elves, 0)
	var counter = 0
	for scanner.Scan() {

		fmt.Println(scanner.Text())
		text := scanner.Text()
		if text == "" {
			fmt.Println("New elf")
			elves = append(elves, 0)
			fmt.Println(elves)
			counter++
		} else {
			var i int
			i, err = strconv.Atoi(text)
			elves[counter] += i
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var minValue, maxValue = MinMax(elves)
	fmt.Println(minValue)
	fmt.Println(maxValue)
	fmt.Println("second part")
	sort.Ints((elves))
	fmt.Println(elves)
	var top3Sum = elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Println(top3Sum)

}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

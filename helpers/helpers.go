package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileAsStrings(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

func ReadFileAs2dInts(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var output [][]int
	counter := 0
	for scanner.Scan() { //for each line
		output = append(output, []int{})
		for _, val := range scanner.Text() { //for each char in line
			stringChar := string(val)
			intChar, err := strconv.Atoi(stringChar)
			if err != nil {
				panic("NaN")
			}
			output[counter] = append(output[counter], intChar)

		}
		//output = append(output, scanner.Text())
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}

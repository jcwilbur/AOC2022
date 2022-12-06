package main

import (
	"AOC2022/helpers"
	"fmt"
)

const markerLength = 14

func main() {
	var inputStrings = helpers.ReadFileAsStrings("input6.txt")
	input := inputStrings[0]

	for i := 0; i < len(input)-markerLength+1; i++ {
		candidate := input[i : i+markerLength]
		if isMarker(candidate) {
			fmt.Println(candidate)
			fmt.Println(i + markerLength)
			break
		}
	}
}

func isMarker(candidate string) bool {
	answer := true
	for i := 0; i < len(candidate); i++ { //loop through each char
		//see if it matches any chars after it
		for j := i + 1; j < len(candidate); j++ {
			if candidate[i] == candidate[j] {
				answer = false
				break
			}
		}
	}
	return answer
}

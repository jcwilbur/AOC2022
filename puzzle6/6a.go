package main

import (
	"AOC2022/helpers"
	"fmt"
)

const markerLength = 14

func main() {
	var inputStrings = helpers.ReadFileAsStrings("input6.txt")
	input := inputStrings[0] //dont need an array this time, just a string

	for i := 0; i < len(input)-markerLength+1; i++ { //loop through each char in the string, stopping in time to not out-of-bounds
		candidate := input[i : i+markerLength] //grab the candidate marker starting with the active character of length markerLength
		if isMarker(candidate) {               //if it's a marker, print it and stop looping
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

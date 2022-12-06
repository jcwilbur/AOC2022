package main

import (
	"AOC2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	quantity int
	from     int
	to       int
}

const crateWidth = 4

func main() {
	fiveA()
}

func fiveA() {
	var inputStrings = helpers.ReadFileAsStrings("input5.txt")

	crates := extractCrates(inputStrings)
	instructions := extractInstructions(inputStrings)

	for _, singleInstruction := range instructions {
		crates = executeInstruction(crates, singleInstruction)
	}

	result := readTopLetterOfEachCrate(crates)
	fmt.Println(result)
}

func readTopLetterOfEachCrate(crates [][]string) string {
	secretMessage := ""
	for _, crate := range crates {
		secretMessage += crate[0]
	}
	return secretMessage
}

func executeInstruction(crates [][]string, singleInstruction instruction) [][]string {
	//grab block to move
	//startIndex :=
	block := make([]string, 0)
	block = append(block, crates[singleInstruction.from-1][0:singleInstruction.quantity]...)
	//block = reverseSlice(block)
	//prepend block at destination
	crates[singleInstruction.to-1] = append(block, crates[singleInstruction.to-1]...)
	//remove block from source
	crates[singleInstruction.from-1] = crates[singleInstruction.from-1][singleInstruction.quantity:]
	return crates
}

func reverseSlice(s []string) []string {
	ss := s
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

func extractInstructions(inputStrings []string) []instruction {
	var instructions []instruction
	for _, inputLine := range inputStrings {
		if inputLine != "" && inputLine[0:4] == "move" {
			linepieces := strings.Split(inputLine, " ")
			myQuantity, err := strconv.Atoi(linepieces[1])
			if err != nil {
				panic("uh oh")
			}
			myFrom, err := strconv.Atoi(linepieces[3])
			if err != nil {
				panic("uh oh")
			}
			myTo, err := strconv.Atoi(linepieces[5])
			if err != nil {
				panic("uh oh")
			}
			myInstruction := instruction{quantity: myQuantity, from: myFrom, to: myTo}
			instructions = append(instructions, myInstruction)
		}
	}
	return instructions
}

func extractCrates(inputStrings []string) [][]string {
	var cratelines []string
	for i := 0; inputStrings[i] != ""; i++ {
		cratelines = append(cratelines, inputStrings[i])
	}
	//destroy last line (containing column numbers)
	cratelines = cratelines[:len(cratelines)-1]

	numColumns := 0
	for _, crateline := range cratelines {
		if len(crateline)/4 > numColumns {
			numColumns = (len(crateline) / 4) + 1
		}
	}
	var crates = make([][]string, numColumns)
	for _, crateline := range cratelines {
		for i := 0; i < len(crateline); i += crateWidth {
			startingPoint := i
			endingPont := i + crateWidth - 1
			if crateline[startingPoint:endingPont] != "   " {
				crateNum := i / crateWidth
				crateValStart := i + 1
				crateValEnd := i + crateWidth - 2

				crateVal := crateline[crateValStart:crateValEnd]
				crates[crateNum] = append(crates[crateNum], crateVal)
			}
		}

	}
	return crates
}

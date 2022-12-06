package main

import (
	"AOC2022/helpers"
	"fmt"
	"strings"
)

type sack struct {
	front string
	back  string
}

func main() {
	threeB()
}

func threeA() {
	inputStrings := helpers.ReadFileAsStrings("input3.txt")

	sacks := []sack{}
	for _, line := range inputStrings {
		half1 := line[0 : len((line))/2]
		half2 := line[len(line)/2:]
		sacks = append(sacks, sack{front: half1, back: half2})
	}
	duplicates := []string{}
	duplicateRunes := []rune{}
	for _, v := range sacks {
		duplicateFound := false
		for _, frontchar := range v.front {
			if !duplicateFound {

				for _, backchar := range v.back {
					if !duplicateFound {

						if frontchar == backchar {
							duplicates = append(duplicates, string(frontchar))
							if frontchar >= 97 { //it's lower case
								duplicateRunes = append(duplicateRunes, frontchar-96)
								duplicateFound = true
							} else { //it's upper case
								duplicateRunes = append(duplicateRunes, frontchar-38)
								duplicateFound = true
							}

						}
					}
				}
			}
		}
	}
	for i := 0; i < len(duplicateRunes); i++ {
		fmt.Printf("letter: %s; rune: %d", duplicates[i], duplicateRunes[i])
		fmt.Println()
	}

	sumPriority := 0

	for _, v := range duplicateRunes {
		sumPriority += int(v)
	}
	fmt.Println(sumPriority)
}

func threeB() {
	inputStrings := helpers.ReadFileAsStrings("input3.txt")

	groups := make([][3]string, len((inputStrings))/3)
	slowCounter := 0
	circleCounter := 0
	for i := 0; i < len(inputStrings); i++ {
		slowCounter = i / 3
		circleCounter = i % 3
		groups[slowCounter][circleCounter] = inputStrings[i]
	}
	duplicateFound := false
	badges := make([]rune, len(groups))
	for i, group := range groups { //for each group
		duplicateFound = false
		for _, char := range group[0] { //for every char in the first line of the group
			if !duplicateFound {
				if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) {
					if char >= 97 { //it's lower case
						badges[i] = char - 96
						duplicateFound = true
					} else { //it's upper case
						badges[i] = char - 38
						duplicateFound = true
					}
				}
			}
		}

	}
	sumPriority := 0

	for _, v := range badges {
		sumPriority += int(v)
	}
	fmt.Println(sumPriority)

}

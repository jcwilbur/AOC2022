package main

import (
	"AOC2022/helpers"
	"fmt"
)

type round struct {
	opponentMove string
	myMove       string
}

func main() {
	var inputStrings = helpers.ReadFileAsStrings("input2.txt")
	fmt.Println(inputStrings)
	rounds := []round{}
	for i, line := range inputStrings {
		rounds = append(rounds, round{opponentMove: line[0:1], myMove: line[2:3]})
		fmt.Println(rounds[i])
	}

	for _, v := range rounds {
		calculatePoints(v)
	}
}

func calculatePoints(v round) {
	panic("unimplemented")
}

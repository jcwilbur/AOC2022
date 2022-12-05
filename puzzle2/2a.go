package main

import (
	"AOC2022/helpers"
	"fmt"
)

type round struct {
	OpponentMove string
	MyMove       string
}

const RockScore = 1
const PaperScore = 2
const ScissorsScore = 3
const WinScore = 6
const DrawScore = 3
const LoseScore = 0

func main() {
	twoB()
}

func twoA() {
	var inputStrings = helpers.ReadFileAsStrings("input2.txt")
	//fmt.Println(inputStrings)
	rounds := []round{}
	for _, line := range inputStrings {
		rounds = append(rounds, round{OpponentMove: line[0:1], MyMove: line[2:3]})
		//fmt.Println(rounds[i])
	}
	totalscore := 0
	for _, v := range rounds {
		totalscore += calculatePoints2a(v)
	}
	fmt.Println(totalscore)
}

func calculatePoints2a(v round) int {
	roundScore := 0
	switch v.MyMove {
	case "X": //I chose Rock
		roundScore += RockScore
		switch v.OpponentMove {
		case "A":
			//Rock v Rock = Draw
			roundScore += DrawScore
		case "B": //Paper v Rock = Lose
			roundScore += LoseScore
		case "C": //Scissors v Rock = Win
			roundScore += WinScore
		}
	case "Y":
		roundScore += PaperScore
		switch v.OpponentMove {
		case "A":
			//Rock v Paper = Win
			roundScore += WinScore
		case "B": //Paper v Paper = Draw
			roundScore += DrawScore
		case "C": //Scissors v Paper = Lose
			roundScore += LoseScore
		}
	case "Z":
		roundScore += ScissorsScore
		switch v.OpponentMove {
		case "A":
			//Rock v Scissors = Lose
			roundScore += LoseScore
		case "B": //Paper v Scissors = Win
			roundScore += WinScore
		case "C": //Scissors v Scissors = Draw
			roundScore += DrawScore
		}
	}
	return roundScore
}

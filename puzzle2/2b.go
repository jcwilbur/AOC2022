package main

import (
	"AOC2022/helpers"
	"fmt"
)

func twoB() {
	var inputStrings = helpers.ReadFileAsStrings("input2.txt")
	//fmt.Println(inputStrings)
	rounds := []round{}
	for _, line := range inputStrings {
		rounds = append(rounds, round{OpponentMove: line[0:1], MyMove: line[2:3]})
		//fmt.Println(rounds[i])
	}
	totalscore := 0
	for _, v := range rounds {
		totalscore += calculatePoints2b(v)
	}
	fmt.Println(totalscore)
}

func calculatePoints2b(v round) int {
	roundScore := 0
	switch v.MyMove {
	case "X": //I need to lose
		roundScore += LoseScore
		switch v.OpponentMove {
		case "A":
			//Opponent threw Rock
			roundScore += ScissorsScore
		case "B": //Opponent threw Paper
			roundScore += RockScore
		case "C": //Opponent threw Scissors
			roundScore += PaperScore
		}
	case "Y": //I need to Draw
		roundScore += DrawScore
		switch v.OpponentMove {
		case "A": //Opponent threw rock
			roundScore += RockScore
		case "B": //Opponent threw paper
			roundScore += PaperScore
		case "C": //Opponent threw scissors
			roundScore += ScissorsScore
		}
	case "Z": //I need to win
		roundScore += WinScore
		switch v.OpponentMove {
		case "A": //Opponent threw Rock
			roundScore += PaperScore
		case "B": //Opponent threw Paper
			roundScore += ScissorsScore
		case "C": //Opponent threw Scissors
			roundScore += RockScore
		}
	}
	return roundScore
}

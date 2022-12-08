package main

import (
	"AOC2022/helpers"
	"fmt"
)

func main() {
	intMatrix := helpers.ReadFileAs2dInts("input8.txt")
	//eightA(intMatrix)
	eightB(intMatrix)
}

func eightB(intMatrix [][]int) {
	treeMatrix := makeTreeMatrix(intMatrix)
	fmt.Println(treeMatrix)
}

type Tree struct {
	height    int
	rowIndex  int
	colIndex  int
	isVisible bool
}

func makeTreeMatrix(intMatrix [][]int) [][]Tree {
	var treeMatrix [][]Tree
	for i, row := range intMatrix {
		treeMatrix = append(treeMatrix, []Tree{})
		for j, treeVal := range row {
			treeMatrix[i] = append(treeMatrix[i], Tree{height: treeVal, rowIndex: i, colIndex: j, isVisible: false})
		}
	}
	return treeMatrix
}

func eightA(intMatrix [][]int) {
	//make a matrix of Trees
	treeMatrix := makeTreeMatrix(intMatrix)

	//scan for visible horizontally
	for i := 0; i < len(treeMatrix); i++ {
		maxHeight := -1
		for j := 0; j < len(treeMatrix[i]); j++ { //check left to right
			if treeMatrix[i][j].height > maxHeight {
				maxHeight = treeMatrix[i][j].height
				treeMatrix[i][j].isVisible = true
			}
		}
		maxHeight = -1
		for j := len(treeMatrix[i]) - 1; j >= 0; j-- { //check right to left
			if treeMatrix[i][j].height > maxHeight {
				maxHeight = treeMatrix[i][j].height
				treeMatrix[i][j].isVisible = true
			}
		}
	}
	//scan for visible vertically
	for i := 0; i < len(treeMatrix); i++ {
		maxHeight := -1
		for j := 0; j < len(treeMatrix[i]); j++ { //check top to bottom
			if treeMatrix[j][i].height > maxHeight {
				maxHeight = treeMatrix[j][i].height
				treeMatrix[j][i].isVisible = true
			}
		}
		maxHeight = -1
		for j := len(treeMatrix[i]) - 1; j >= 0; j-- { //check bottom to top
			if treeMatrix[j][i].height > maxHeight {
				maxHeight = treeMatrix[j][i].height
				treeMatrix[j][i].isVisible = true
			}
		}
	}

	//count visible trees
	count := 0
	for _, row := range treeMatrix {
		fmt.Println()
		for _, tree := range row {
			fmt.Print(tree.isVisible)
			fmt.Print((" "))
			if tree.isVisible {
				count++
			}
		}
	}
	fmt.Println(count)
}

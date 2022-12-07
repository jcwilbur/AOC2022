package main

import (
	"AOC2022/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Size     int
	Children []*Node
	Parent   *Node
	Name     string
}

const sizeLimit = 100000
const totalSpace = 70000000
const spaceNeeded = 30000000

func main() {
	inputStrings := helpers.ReadFileAsStrings("input.txt")
	sevenB(inputStrings)
}

func sevenB(inputStrings []string) {
	directoryTree := buildTree(inputStrings)
	spaceAvailable := totalSpace - directoryTree.Size
	spaceToClear := spaceNeeded - spaceAvailable
	bestBet := directoryTree
	answer := calculateAnswerForSevenB(directoryTree, spaceToClear, bestBet)
	fmt.Println(answer.Name)
	fmt.Println(answer.Size)
}

func calculateAnswerForSevenB(nodeToSearch Node, spaceToClear int, currentBestBet Node) Node {

	//create table of directories and sizes
	if nodeToSearch.Children == nil {
		//I'm a file, so don't add anything to total
	} else {
		//I'm a directory, so see if I'm a better fit than the current Node
		if nodeToSearch.Size >= spaceToClear && nodeToSearch.Size < currentBestBet.Size {
			currentBestBet = nodeToSearch
		}
		//repeat for my children
		for _, child := range nodeToSearch.Children {
			//fmt.Println("searching child node " + child.Name + " running total is " + strconv.Itoa(runningTotal))
			currentBestBet = calculateAnswerForSevenB(*child, spaceToClear, currentBestBet)
			//fmt.Println("done searching child node " + child.Name + " running total is " + strconv.Itoa(runningTotal))
		}
	}

	return currentBestBet
}
func sevenA(inputStrings []string) {
	directoryTree := buildTree(inputStrings)
	answer := calculateAnswerForSevenA(directoryTree, 0)
	fmt.Println("answer is...")
	fmt.Println(answer)
}

func calculateAnswerForSevenA(directoryTree Node, runningTotal int) int {

	if directoryTree.Children == nil {
		//I'm a file, so don't add anything to total
	} else {
		//I'm a directory, so check my file size and add if under the limit
		if directoryTree.Size <= sizeLimit {
			runningTotal += directoryTree.Size
			fmt.Println("added directory " + directoryTree.Name + " with file size " + strconv.Itoa(directoryTree.Size))
			fmt.Println("running total is " + strconv.Itoa(runningTotal))
			//fmt.Print(directoryTree.Name)
		}
		for _, child := range directoryTree.Children {
			fmt.Println("searching child node " + child.Name + " running total is " + strconv.Itoa(runningTotal))
			runningTotal = calculateAnswerForSevenA(*child, runningTotal)
			fmt.Println("done searching child node " + child.Name + " running total is " + strconv.Itoa(runningTotal))
		}
	}

	return runningTotal
}

func buildTree(inputStrings []string) Node {
	root := Node{Size: 0, Name: "/"}
	currentNode := &root
	for _, line := range inputStrings {
		//fmt.Println(line)
		lineSections := strings.Split(line, " ")
		switch lineSections[0] {
		case "$":
			switch lineSections[1] {
			case "cd":
				//"cd command")
				switch lineSections[2] {
				case "/":
					currentNode = &root
				case "..":
					currentNode = currentNode.Parent
				default:
					//set currentNode to currentNode's child with name linesections[2]
					shouldIWorry := true
					for _, potentialTarget := range currentNode.Children {
						if potentialTarget.Name == lineSections[2] {
							currentNode = potentialTarget
							shouldIWorry = false
							break
						}
					}
					if shouldIWorry {
						panic("directory not found")
					}
				}

			case "ls":
				//"ls command - ignore?")
				//do nothing?
			}
		case "dir":
			//"directory")
			newNode := Node{Size: 0, Name: lineSections[1], Parent: currentNode}
			currentNode.Children = append(currentNode.Children, &newNode)
		default:
			//"file")

			size, err := strconv.Atoi(lineSections[0])
			if err != nil {
				panic("NaN")
			}
			newNode := Node{Size: size, Name: lineSections[1], Parent: currentNode}
			currentNode.Children = append(currentNode.Children, &newNode)
			//TODO: cascade up file size
			adderNode := currentNode
			for adderNode != nil {
				adderNode.Size += size
				adderNode = adderNode.Parent
			}

		}
	}
	return root
}

package main

import (
	"adventofcode2020/utils"
	"fmt"
)

func main() {
	input := utils.ReadInput("day3/input")
	symbol := '#'

	part1(input, symbol)
	part2(input, symbol)
}

func part1(input []string, symbol rune)  {
	horizontalSpeed := 1
	treeCount := 0
	horizontalIndex := horizontalSpeed
	for _, line := range input[1:] {
		if toboggan(line, horizontalIndex, symbol) {
			treeCount++
		}
		horizontalIndex += horizontalSpeed
	}
	fmt.Println("Part 1: ", treeCount, "trees passed")
}

func part2(input []string, symbol rune)  {
	multiplyTrees := 1
	treeCount := 0
	verticalCount := 0

	// Traversal
	horizontalSpeeds := []int{1, 3, 5, 7, 1}
	verticalSpeeds := []int{1, 1, 1, 1, 2}

	// Loop through all traversals
	for i := 0; i < len(horizontalSpeeds); i++ {
		horizontalIndex := horizontalSpeeds[i]

		for _, line := range input[1:] {
			verticalCount++
			if verticalCount == verticalSpeeds[i] {
				if toboggan(line, horizontalIndex, symbol) {
					treeCount++
				}
				horizontalIndex += horizontalSpeeds[i]
				verticalCount = 0
			}
		}
		multiplyTrees = multiplyTrees * treeCount
		treeCount = 0
	}

	fmt.Println("Part 2: ", multiplyTrees, "trees passed")
}


func toboggan(input string, hortizontalSpeed int, symbol rune) bool {
	index :=  hortizontalSpeed % len(input)
	if rune(input[index]) == symbol {
		return true
	}
	return false
}
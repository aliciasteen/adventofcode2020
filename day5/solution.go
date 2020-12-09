package main

import (
	"adventofcode2020/utils"
	"fmt"
	"sort"
)

func main() {
	input := utils.ReadInput("day5/input")

	highestSeatID := 0
	var seatIDs []int
	for _, line := range input {
		row := getRow(line[:7])
		column := getColumn(line[7:])
		seatID := getSeatID(row, column)
		seatIDs = append(seatIDs, seatID)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	fmt.Println("Highest Seat ID:", highestSeatID)
	fmt.Println("Seat IDs:", findMissingID(seatIDs))
}

func getRow(input string) int {
	min := 0
	max := 127
	for _, character := range input {
		if character == 'F' {
			max = max - ((max-min)/2)
		} else if character == 'B' {
			min = max - ((max-min)/2)
		}
	}
	return min
}

func getColumn(input string) int {
	min := 0
	max := 7
	for _, character := range input {
		if character == 'L' {
			max = max - ((max-min)/2)
		} else if character == 'R' {
			min = max - ((max-min)/2)
		}
	}
	return min
}

func getSeatID(row, column int) int {
	return row * 8 + column
}

func findMissingID(seatIDs []int) int {
	sort.Ints(seatIDs)
	for i, id := range seatIDs {
		if i+1 != len(seatIDs) {
			if id+1 != seatIDs[i+1] {
				return id+1

			}
		}
	}
	return 0
}
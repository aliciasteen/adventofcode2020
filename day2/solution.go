package main

import (
	"adventofcode2020/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var countPart1, countPart2 int
	input := utils.ReadInput("day2/input")
	for _, line := range input {
		countPart1 += validatePart1(line)
		countPart2 += validatePart2(line)
	}
	fmt.Println("Part 1: ", countPart1, "passwords meet requirements")
	fmt.Println("Part 2: ", countPart2, "passwords meet requirements")

}

func validatePart1(input string) int {
	s := strings.Split(input, " ")

	low, _ := strconv.Atoi(strings.Split(s[0], "-")[0])
	high, _ := strconv.Atoi(strings.Split(s[0], "-")[1])

	character := s[1][:1]
	password := s[2]

	count := strings.Count(password, character)

	if count >= low && count <= high {
		return 1
	}
	return 0
}

func validatePart2(input string) int {
	s := strings.Split(input, " ")

	positionOne, _ := strconv.Atoi(strings.Split(s[0], "-")[0])
	positionOne = positionOne -1
	positionTwo, _ := strconv.Atoi(strings.Split(s[0], "-")[1])
	positionTwo = positionTwo -1

	character := s[1][:1]
	password := s[2]

	var count int
	if string(password[positionOne]) == character {
		count++
	}
	if string(password[positionTwo]) == character {
		count++
	}

	if count == 1 {
		return 1
	} else {
		return 0
	}
}
package main

import (
	"adventofcode2020/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.ReadInputMultiLine("day6/input", "")
 	sum := 0
	for _, group := range input {
		_, questions := processGroup(group)
		sum += questions
	}
	fmt.Println("Number of questions anwsered yes:", sum)
}

func processGroup(input string) (numberOfPeople, numberOfQuestions int) {
	people := strings.Fields(input)
	input = strings.ReplaceAll(input, " ", "")
	questions := make(map[rune]int)
	for _, char := range input {
		questions[char]++
	}
	everyoneAnwsered := 0
	for _, value := range questions {
		if value == len(people) {
			everyoneAnwsered ++
		}
	}
	return len(people), everyoneAnwsered
}
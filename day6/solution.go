package main

import (
	"adventofcode2020/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := utils.ReadInputMultiLine("day6/input", "")

	sum := 0
	for _, line := range input {
		_, questions := processGroup(line)
		sum += questions
	}
	fmt.Println("Number of questions anwsered yes:", sum)
}

func processGroup(input string) (numberOfPeople, numberOfQuestions int) {
	people := strings.Fields(input)
	sort.Strings(people)
	input = strings.ReplaceAll(input, " ", "")
	questions := make(map[string]int)
	for _, char := range input {
		if _, exists := questions[string(char)]; exists {
			questions[string(char)]++
		} else {
			questions[string(char)] = 1
		}
	}
	everyoneAnwsered := 0
	for _, value := range questions {
		if value == len(people) {
			everyoneAnwsered ++
		}
	}
	return len(people), everyoneAnwsered
}
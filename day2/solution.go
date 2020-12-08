package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count int
	input := readInput("day2/input")
	for _, line := range input {
		count += validate(line)
	}
	fmt.Println(count, "passwords meet requirements")
}

func readInput(path string) []string {
	var input []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func validate(input string) int {
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
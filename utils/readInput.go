package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(path string) []string {
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

func ReadInputMultiLine(path, delimiter string) []string {
	input := []string{""}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if scanner.Text() != delimiter {
			input[i] = input[i] + scanner.Text() + " "
			i--
		} else {
			input = append(input, "")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
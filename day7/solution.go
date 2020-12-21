package main

import (
	"adventofcode2020/utils"
	"fmt"
	"strconv"
	"strings"
)

type innerBag struct {
	colour string
	quantity int
}

func main() {
	input := utils.ReadInput("day7/input")
	rules := processRules(input)
	bagCount := 0
	for outerBag, _ := range rules {
		if solutionPart1(rules, outerBag, "shiny gold") {
			bagCount++
		}
	}

	fmt.Println(solutionPart2(rules, "shiny gold")-1)

}

func processRules(input []string) map[string][]innerBag {
	rules := make(map[string][]innerBag)
	for _, line := range input {
		outerBag := before(line, " bags contain")
		innerBags := strings.Split(before(after(line, "bags contain "), "."),  ", ")
		var innerBagSlice []innerBag

		if innerBags[0] == "no other bags" {
			rules[outerBag] = nil
		} else {
			for _, bag := range innerBags {
				bagQuantity, _ := strconv.Atoi(bag[:1])
				bagColour := strings.Split(bag," bag")[0][2:]
				b := innerBag{colour: bagColour, quantity: bagQuantity}
				innerBagSlice = append(innerBagSlice, b)
			}
		}
		rules[outerBag] = innerBagSlice
	}
	return rules
}

func solutionPart1(rules map[string][]innerBag, outerBag, targetColour string) bool {
	innerBags := rules[outerBag]
	for _, bag := range innerBags {
		if bag.colour == targetColour {
			return true
		}
		if solutionPart1(rules, bag.colour, targetColour) {
			return true
		}
	}
	return false
}

func solutionPart2(rules map[string][]innerBag, currentBag string) int {
	if len(rules[currentBag]) == 0 {
		return 1
	} else {
		total := 1
		for _, bag := range rules[currentBag] {
			total += solutionPart2(rules, bag.colour) * bag.quantity
		}
		return total
	}
}

// Helper functions

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

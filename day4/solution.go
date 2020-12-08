package main

import (
	"adventofcode2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInputMultiLine("day4/input", "")
	validPassportCount := 0
	optionalFields := []string{"cid"}
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, line := range input {
		if verifyPassport(line, requiredFields, optionalFields) {
			validPassportCount++
		}
	}
	fmt.Println(validPassportCount)
}

func verifyPassport(input string, requiredFields, optionalFields []string) bool {
	passport := make(map[string]string)
	i := strings.Fields(input)
	for _, line := range i {
		passport[line[:3]] = line[4:]
	}
	for _, field := range requiredFields {
		if _, found := passport[field]; !found {
			return false
		}
	}
	return verifyValues(passport)
}

func verifyValues(passport map[string]string) bool {
	byr, _ := strconv.Atoi(passport["byr"])
	iyr, _ := strconv.Atoi(passport["iyr"])
	eyr, _ := strconv.Atoi(passport["eyr"])
	return validateYear(byr, 1920, 2002) &&
		validateYear(iyr, 2010, 2020) &&
		validateYear(eyr, 2020, 2030) &&
		validateHeight(passport["hgt"]) &&
		validateRegex(passport["hcl"], `^#(?:[0-9a-fA-F]{6})$`) &&
		validateEyeColor(passport["ecl"]) &&
		validateRegex(passport["pid"], `^(?:[0-9]{9})$`)

}

func validateEyeColor(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}


func validateRegex(s, p string) bool {
	matched, _ := regexp.MatchString(p, s)
	return matched
}

func validateHeight(s string) bool {
	height, _ := strconv.Atoi(s[:len(s)-2])
	unit := s[len(s)-2:]
	if unit == "cm" {
		return height >= 150 && height <= 193
	} else if unit == "in" {
		return height >= 59 && height <= 76
	}
	return false
}

func validateYear(year, lowerBound, upperBound int) bool {
	return year >= lowerBound && year <= upperBound
}
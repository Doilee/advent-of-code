package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid int
	cid int //optional
}

func main() {
	bytes, _ := os.ReadFile("2020/day4.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n\n")
	passports := make([]map[string]string, 0)

	for _, line := range lines {
		cleaned := strings.Replace(line, "\n", " ", -1)

		passport := make(map[string]string)

		fields := strings.Split(cleaned, " ")

		for _, field := range fields {
			keyValue := strings.Split(field, ":")
			key, value := keyValue[0], keyValue[1]

			passport[key] = value
			//passport[key] := value
		}
		passports = append(passports, passport)
	}

	validPassportsPartOne := 0
	validPassportsPartTwo := 0

	for _, passport := range passports {
		if determinePartOneValidPassport(passport) {
			validPassportsPartOne++
		}
		if determinePartTwoValidPassport(passport) {
			validPassportsPartTwo++
		}
	}

	fmt.Println("Part One: ")
	fmt.Println(validPassportsPartOne)
	fmt.Println("Part Two: ")
	fmt.Println(validPassportsPartTwo)
}

func determinePartOneValidPassport(passport map[string]string) bool {
	_, hasCountryId := passport["cid"]

	if len(passport) == 8 || (len(passport) == 7 && !hasCountryId) {
		return true
	}

	return false
}

func determinePartTwoValidPassport(passport map[string]string) bool {
	check := 0

	if byr, exists := passport["byr"]; exists {
		if validateStringAsNumber(byr, 1920, 2002) {
			check++
		}
	}

	if iyr, exists := passport["iyr"]; exists {
		if validateStringAsNumber(iyr, 2010, 2020) {
			check++
		}
	}

	if eyr, exists := passport["eyr"]; exists {
		if validateStringAsNumber(eyr, 2020, 2030) {
			check++
		}
	}

	if hgt, exists := passport["hgt"]; exists {
		characterSeperationIndex := len(hgt) - 2
		unit := hgt[characterSeperationIndex:]
		number := hgt[0:characterSeperationIndex]
		if unit == "cm" && validateStringAsNumber(number, 150, 193) {
			check++
		} else if unit == "in" && validateStringAsNumber(number, 59, 76) {
			check++
		}
	}

	if hcl, exists := passport["hcl"]; exists {
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		if re.MatchString(hcl) {
			check++
		}
	}

	if ecl, exists := passport["ecl"]; exists {
		re := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
		if re.MatchString(ecl) {
			check++
		}
	}

	if pid, exists := passport["pid"]; exists {
		re := regexp.MustCompile(`^\d{9}$`)
		if re.MatchString(pid) {
			check++
		}
	}

	return check == 7
}

func validateStringAsNumber(value string, atleast int, atmost int) bool {
	toValidate, _ := strconv.Atoi(value)

	return toValidate >= atleast && toValidate <= atmost
}

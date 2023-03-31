package main

import (
	"fmt"
	"os"
	strconv "strconv"
	"strings"
)

type puzzleInputStruct struct {
	minimum   int
	maximum   int
	character string
	password  string
}

func day2() {
	bytes, _ := os.ReadFile("2020/input/day2.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")
	values := make([]puzzleInputStruct, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		minMax := strings.Split(parts[0], "-")
		values[i].minimum, _ = strconv.Atoi(minMax[0])
		values[i].maximum, _ = strconv.Atoi(minMax[1])
		values[i].character = strings.TrimSuffix(parts[1], ":")
		values[i].password = parts[2]
	}

	correctAmountPart1 := 0
	correctAmountPart2 := 0

	for _, value := range values {
		if strings.Count(value.password, value.character) >= value.minimum &&
			strings.Count(value.password, value.character) <= value.maximum {
			correctAmountPart1++
		}

		if string(value.password[value.minimum-1]) == value.character &&
			string(value.password[value.maximum-1]) != value.character {
			correctAmountPart2++
		}

		if string(value.password[value.minimum-1]) != value.character &&
			string(value.password[value.maximum-1]) == value.character {
			correctAmountPart2++
		}
	}

	fmt.Println(correctAmountPart1)
	fmt.Println(correctAmountPart2)
}

func main() {
	day2()
}

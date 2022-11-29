package main

import (
	"fmt"
	"os"
	strconv "strconv"
	"strings"
)

func day1() {
	arr := [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	var lastNumber int

	for i := 0; i < len(arr); i++ {
		var currentNumber = arr[i]
		if i == 0 {
			fmt.Println("N/A - no previous measurement")
		} else if currentNumber > lastNumber {
			fmt.Println("Increasing")
		} else if currentNumber < lastNumber {
			fmt.Println("Decreasing")
		}
		lastNumber = arr[i]
	}
}

type valueStruct struct {
	direction string
	amount    int
}

func day2() {
	bytes, _ := os.ReadFile("data.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")
	values := make([]valueStruct, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		values[i].direction = parts[0]
		values[i].amount, _ = strconv.Atoi(parts[1])
	}

	y := 0
	x := 0

	for _, value := range values {
		if value.direction == "up" {
			y -= value.amount
		} else if value.direction == "down" {
			y += value.amount
		} else if value.direction == "forward" {
			x += value.amount
		} else {
			x -= value.amount
		}
	}

	fmt.Printf("%v", x*y)
}

func year2020day1() {
	bytes, _ := os.ReadFile("day1.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")
	values := make([]int, len(lines))

	for i, line := range lines {
		values[i], _ = strconv.Atoi(line)
	}

	for _, value := range values {
		for _, value2 := range values {
			for _, value3 := range values {
				if value+value2+value3 == 2020 {
					fmt.Printf("%v * %v * %v = %v\n", value, value2, value3, value*value2*value3)
				}
			}
		}
	}
}

type puzzleInputStruct struct {
	minimum   int
	maximum   int
	character string
	password  string
}

func year2020day2() {
	bytes, _ := os.ReadFile("day2.txt")
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
	year2020day2()
}

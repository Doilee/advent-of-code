package main

import (
	"fmt"
	"os"
	strconv "strconv"
	"strings"
)

type valueStruct struct {
	direction string
	amount    int
}

func day2() {
	bytes, _ := os.ReadFile("2021/input/02.txt")
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

func main() {
	day2()
}

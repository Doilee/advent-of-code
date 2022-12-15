package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2022/day9/larger-example-data.txt")
	puzzleInput := string(bytes)
	data := strings.Split(puzzleInput, "\n")

	xLength := 350
	yLength := 350

	matrix := make([][]string, yLength)
	for i := 0; i < yLength; i++ {
		matrix[i] = make([]string, xLength)
	}

	//coordinates Y and X
	rope := make([][]int, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = []int{100, 100}
	}

	for _, input := range data {
		direction, amount := splitInput(input, " ")

		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				rope[0][0]++
			case "U":
				rope[0][1]++
			case "L":
				rope[0][0]--
			case "D":
				rope[0][1]--
			}

			for index, knot := range rope {
				if index == 0 {
					continue
				}

				rope[index] = setKnot(rope[index-1], knot)
			}

			matrix[rope[9][0]][rope[9][1]] = "#"
		}

		fmt.Println(rope)
	}

	answer := 0

	for _, row := range matrix {
		for _, value := range row {
			if value == "#" {
				answer++
			}
		}
	}

	fmt.Println("The count is: ", answer)
}

func splitInput(s, sep string) (string, int) {
	x := strings.Split(s, sep)
	amount, _ := strconv.Atoi(x[1])
	return x[0], amount
}

func setKnot(leading []int, following []int) []int {
	if leading[0] > following[0]+1 {
		following[0] = leading[0] - 1
		following[1] = leading[1]
	} else if leading[0] < following[0]-1 {
		following[0] = leading[0] + 1
		following[1] = leading[1]
	}

	if leading[1] > following[1]+1 {
		following[0] = leading[0]
		following[1] = leading[1] - 1
	} else if leading[1] < following[1]-1 {
		following[0] = leading[0]
		following[1] = leading[1] + 1
	}

	return following
}

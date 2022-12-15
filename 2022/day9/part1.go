package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	bytes, _ := os.ReadFile("2022/day9/data.txt")
	puzzleInput := string(bytes)
	data := strings.Split(puzzleInput, "\n")

	xLength := 500
	yLength := 500

	matrix := make([][]string, yLength)
	for i := 0; i < yLength; i++ {
		matrix[i] = make([]string, xLength)
	}

	// for logging
	for x, row := range matrix {
		for y, _ := range row {
			matrix[x][y] = "."
		}
	}

	//coordinates Y and X
	head := []int{250, 250}
	tail := []int{250, 250}

	for _, input := range data {
		matrix[tail[0]][tail[1]] = "#"

		direction, amount := splitInput(input, " ")

		fmt.Print("\033[H\033[2J")

		for i := 0; i < amount; i++ {
			switch direction {
			case "R":
				head[0]++
			case "U":
				head[1]++
			case "L":
				head[0]--
			case "D":
				head[1]--
			}

			if head[0] > tail[0]+1 {
				tail[0] = head[0] - 1
				tail[1] = head[1]
			} else if head[0] < tail[0]-1 {
				tail[0] = head[0] + 1
				tail[1] = head[1]
			}

			if head[1] > tail[1]+1 {
				tail[0] = head[0]
				tail[1] = head[1] - 1
			} else if head[1] < tail[1]-1 {
				tail[0] = head[0]
				tail[1] = head[1] + 1
			}

			matrix[tail[0]][tail[1]] = "#"
		}
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

func visualizeData(head []int, tail []int, matrix [][]string) {

	for x, row := range matrix {
		for y, _ := range row {
			if head[0] == x && head[1] == y {
				fmt.Print(" H")
			} else if tail[0] == x && tail[1] == y {
				fmt.Print(" T")
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(head)
	fmt.Println(tail)

	time.Sleep(10 * time.Millisecond)

	fmt.Print("\033[H\033[2J")
}

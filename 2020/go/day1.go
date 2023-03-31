package main

import (
	"fmt"
	"os"
	strconv "strconv"
	"strings"
)

func day1() {
	bytes, _ := os.ReadFile("2020/input/day1.txt")
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

func main() {
	day1()
}

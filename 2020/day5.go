package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2020/day5.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")

	highestSeatId := 0
	var seatIds []int

	for _, line := range lines {
		minRow := 0
		maxRow := 127
		minColumn := 0
		maxColumn := 7
		rowDivider := 64
		columnDivider := 4
		var chosenRow int
		var chosenColumn int

		for _, command := range line {

			switch command {
			// lower half
			case 'F':
				if rowDivider == 1 {
					chosenRow = minRow
				} else {
					maxRow = maxRow - rowDivider
					rowDivider = rowDivider / 2
				}
			// upper half
			case 'B':
				if rowDivider == 1 {
					chosenRow = maxRow
				} else {
					minRow = minRow + rowDivider
					rowDivider = rowDivider / 2
				}
			// lower half
			case 'L':
				if columnDivider == 1 {
					chosenColumn = minColumn
				} else {
					maxColumn = maxColumn - columnDivider
					columnDivider = columnDivider / 2
				}
			// upper half
			case 'R':
				if columnDivider == 1 {
					chosenColumn = maxColumn
				} else {
					minColumn = minColumn + columnDivider
					columnDivider = columnDivider / 2
				}
			}
		}

		seatId := chosenRow*8 + chosenColumn

		seatIds = append(seatIds, seatId)

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	fmt.Println("Answer part one:")
	fmt.Println(highestSeatId)

	check := false

	sort.Ints(seatIds)

	for i, seatId := range seatIds {
		if i == 0 || i == len(seatIds)-1 {
			continue
		}

		if check && seatIds[i-1] == seatId-2 {
			fmt.Println("Answer part two:")
			fmt.Println(seatId - 1)
		}

		check = false

		if seatIds[i+1] == seatId+2 {
			check = true
		}

	}
}

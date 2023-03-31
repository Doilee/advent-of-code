package main

import (
	"fmt"
	"os"
)

const PART_ONE = 4
const PART_TWO = 14

func main() {
bytes, _ := os.ReadFile("2022/input/06.txt")
	signal := string(bytes)

	var lastFour []rune

	for index, character := range signal {
		unique := make(map[rune]bool, 26)
		lastFour = append(lastFour, character)

		if len(lastFour) > PART_ONE {
			lastFour = lastFour[1:]
		}

		lastFourAreUnique := true
		for _, char := range lastFour {

			if unique[char] {
				lastFourAreUnique = false
				break
			}

			unique[char] = true
		}

		if lastFourAreUnique && len(lastFour) == PART_ONE {
			fmt.Println(index + 1)
			break
		}
	}
}

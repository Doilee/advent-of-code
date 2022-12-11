package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2022/day4/day4.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")

	fullyContainsCount := 0
	overlapCount := 0
	for _, line := range lines {
		fmt.Println(line)
		parts := strings.Split(line, ",")
		unparsedSectionsFirstElf := parts[0]
		unparsedSectionsSecondElf := parts[1]

		sectionsFirstElf := turnStringRangeIntoArrayOfValues(unparsedSectionsFirstElf)
		sectionsSecondElf := turnStringRangeIntoArrayOfValues(unparsedSectionsSecondElf)

		intersection := intersect.Simple(sectionsFirstElf, sectionsSecondElf)
		if len(intersection) == len(sectionsFirstElf) || len(intersection) == len(sectionsSecondElf) {
			fullyContainsCount++
		}
		if len(intersection) > 0 {
			overlapCount++
		}
	}

	fmt.Println("Answer part 1: ")
	fmt.Println(fullyContainsCount)
	fmt.Println("Answer part 2: ")
	fmt.Println(overlapCount)
}

func turnStringRangeIntoArrayOfValues(s string) []int {
	parts := strings.Split(s, "-")
	var slice []int
	max, _ := strconv.Atoi(parts[1])
	for i, _ := strconv.Atoi(parts[0]); i <= max; i++ {
		slice = append(slice, i)
	}

	return slice
}

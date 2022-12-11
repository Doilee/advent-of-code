package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"math"
	"os"
	"reflect"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2022/day3.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")

	// get rune ints
	fmt.Println('a')
	fmt.Println('z')
	fmt.Println('A')
	fmt.Println('Z')
	var totalPriority uint64

	groups := make(map[uint][]string)

	for i, rucksack := range lines {
		var duplicateItem rune
		groupId := uint(math.Floor(float64(i) / 3.0))

		groups[groupId] = append(groups[groupId], rucksack)

		firstCompartment := rucksack[0 : len(rucksack)/2]
		secondCompartment := rucksack[len(rucksack)/2:]

		for _, item := range rucksack {
			inFirstCompartment := strings.Count(firstCompartment, string(item))
			inSecondCompartment := strings.Count(secondCompartment, string(item))

			if inFirstCompartment > 0 && inSecondCompartment > 0 {
				duplicateItem = item
			}
		}

		totalPriority += getCharacterPriority(duplicateItem)
	}

	fmt.Println("Answer part 1: ")
	fmt.Println(totalPriority)

	totalPriority = 0

	for _, group := range groups {

		firstIntersectionGroup := intersect.Simple(group[0], group[1])
		intersection := intersect.Simple(firstIntersectionGroup, group[2])

		totalPriority = totalPriority + getCharacterPriority(rune(reflect.ValueOf(intersection[0]).Uint()))
	}

	fmt.Println("Answer part 2: ")
	fmt.Println(totalPriority)
}

func getCharacterPriority(character rune) uint64 {
	if character >= 97 && character <= 122 {
		return uint64(character) - 96
	} else if character >= 65 && character <= 90 {
		return uint64(character) - 38 // 65 - 38 = 27..
	}

	return 0
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2022/day1.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n\n")

	answerPartOne := 0
	var caloriesPerElf []int

	for _, line := range lines {
		calorieSum := 0
		calories := strings.Split(line, "\n")

		for _, calorie := range calories {
			toInt, _ := strconv.Atoi(calorie)
			calorieSum += toInt
		}

		if calorieSum > answerPartOne {
			answerPartOne = calorieSum
		}

		caloriesPerElf = append(caloriesPerElf, calorieSum)
	}
	sort.Ints(caloriesPerElf)
	highestThreeCalorieElves := caloriesPerElf[len(caloriesPerElf)-3:]
	answerPartTwo := sum(highestThreeCalorieElves)

	fmt.Println(answerPartOne)
	fmt.Println(answerPartTwo)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

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

	var caloriesPerElf []int

	for _, line := range lines {
		calorieSum := 0
		calories := strings.Split(line, "\n")

		for _, calorie := range calories {
			toInt, _ := strconv.Atoi(calorie)
			calorieSum += toInt
		}

		caloriesPerElf = append(caloriesPerElf, calorieSum)
	}
	sort.Ints(caloriesPerElf)
	highestThreeCalorieElves := caloriesPerElf[len(caloriesPerElf)-3:]

	answerPartOne := caloriesPerElf[len(caloriesPerElf)-1]
	answerPartTwo := array_sum(highestThreeCalorieElves)

	fmt.Println(answerPartOne)
	fmt.Println(answerPartTwo)
}

func array_sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

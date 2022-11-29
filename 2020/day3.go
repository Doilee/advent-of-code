package main

import (
	"fmt"
	"os"
	"strings"
)

const maxXDistance = 31

func parseInput() {
	bytes, _ := os.ReadFile("2020/day3.txt")
	puzzleInput := string(bytes)
	lines := strings.Split(puzzleInput, "\n")

	x := 0
	treesEncountered := 0

	for _, line := range lines {
		if line[x%31] == '#' {
			treesEncountered++
		}
		x += 3
	}

	fmt.Println("Part 1 solution is:", getTreesEncountered(lines, 3, 1))

	part2Solution := getTreesEncountered(lines, 1, 1) *
		getTreesEncountered(lines, 3, 1) *
		getTreesEncountered(lines, 5, 1) *
		getTreesEncountered(lines, 7, 1) *
		getTreesEncountered(lines, 1, 2)

	fmt.Println("Part 2 solution is:", part2Solution)
}

func getTreesEncountered(lines []string, right int, down int) int {
	x := 0
	treesEncountered := 0

	for y := 0; y < len(lines); y += down {
		if lines[y][x%maxXDistance] == '#' {
			treesEncountered++
		}
		x += right
	}

	return treesEncountered
}

func main() {
	parseInput()
}

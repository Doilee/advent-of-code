package main

import (
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	Height int
	Top    bool
	Left   bool
	Right  bool
	Bottom bool
}

func main() {
	bytes, _ := os.ReadFile("2022/day8.txt")
	puzzleInput := string(bytes)
	data := strings.Split(puzzleInput, "\n")

	yLength := len(data)
	xLength := len(data[0])

	trees := make([][]int, yLength)
	for i := 0; i < yLength; i++ {
		trees[i] = make([]int, xLength)
	}

	for y, row := range data {
		for x, tree := range row {
			trees[y][x] = int(tree)
		}
	}

	// getPartOneAnswer(yLength, xLength, trees)

	partTwoAnswer := 0
	for y, row := range trees {
		for x, height := range row {
			scenicScore := calculateScenicScore(height, y, x, yLength, xLength, trees)
			if scenicScore > partTwoAnswer {
				partTwoAnswer = scenicScore
			}
		}
	}

	fmt.Println(partTwoAnswer)

}

func calculateScenicScore(treeHeight int, treeY int, treeX int, yLength int, xLength int, trees [][]int) int {
	toLeft := 0
	// To left
	for y := treeY - 1; y >= 0; y-- {
		toLeft++
		if trees[y][treeX] >= treeHeight {
			break
		}
	}

	toRight := 0
	// To Right
	for y := treeY + 1; y < yLength; y++ {
		toRight++
		if trees[y][treeX] >= treeHeight {
			break
		}
	}

	toTop := 0
	// To Top
	for x := treeX - 1; x >= 0; x-- {
		toTop++

		if trees[treeY][x] >= treeHeight {
			break
		}
	}

	toBottom := 0
	// To Bottom
	for x := treeX + 1; x < xLength; x++ {
		toBottom++
		if trees[treeY][x] >= treeHeight {
			break
		}
	}

	return toLeft * toRight * toBottom * toTop
}

func getPartOneAnswer(yLength int, xLength int, trees [][]int) {
	matrix := make([][]bool, yLength) // initialize a slice of dy slices
	for i := 0; i < yLength; i++ {
		matrix[i] = make([]bool, xLength) // initialize a slice of dx unit8 in each of dy slices
	}

	// From left first
	for y, row := range trees {
		highestLeftTreeHeight := -1

		for x, tree := range row {
			isVisible := highestLeftTreeHeight < tree

			matrix[x][y] = isVisible

			if isVisible {
				highestLeftTreeHeight = tree
			}
		}
	}

	// From Top
	for x := 0; x < xLength; x++ {
		highestTopTreeHeight := -1

		for y := 0; y < yLength; y++ {
			treeHeight := trees[y][x]

			isVisible := highestTopTreeHeight < treeHeight

			if !matrix[x][y] {
				matrix[x][y] = isVisible
			}

			if isVisible {
				highestTopTreeHeight = treeHeight
			}
		}
	}

	// From Right
	for y := yLength - 1; y >= 0; y-- {
		highestRightTreeHeight := -1

		for x := xLength - 1; x >= 0; x-- {
			treeHeight := trees[y][x]

			isVisible := highestRightTreeHeight < treeHeight

			if !matrix[x][y] {
				matrix[x][y] = isVisible
			}

			if isVisible {
				highestRightTreeHeight = treeHeight
			}
		}
	}

	// From Bottom
	for x := xLength - 1; x >= 0; x-- {
		highestBottomTreeHeight := -1

		for y := yLength - 1; y >= 0; y-- {
			treeHeight := trees[y][x]

			isVisible := highestBottomTreeHeight < treeHeight

			if !matrix[x][y] {
				matrix[x][y] = isVisible
			}

			if isVisible {
				highestBottomTreeHeight = treeHeight
			}
		}
	}

	visibleTrees := 0

	for _, column := range matrix {
		for _, tree := range column {
			if tree {
				visibleTrees++
			}
		}
	}

	fmt.Println(visibleTrees)
}

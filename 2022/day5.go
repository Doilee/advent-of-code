package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("2022/day5.txt")
	puzzleInput := string(bytes)
	data := strings.Split(puzzleInput, "\n\n")

	crates := strings.Split(data[0], "\n")

	crateMap := make(map[int]string)

	crates = crates[:len(crates)-1]

	for _, crate := range crates {
		for x, crateCharacter := range crate {
			for z := 0; z <= 9; z++ {
				if x == 1+4*z && crateCharacter != ' ' {
					crateMap[z+1] += string(crateCharacter)
				}
			}
		}
	}

	commands := strings.Split(data[1], "\n")

	for _, command := range commands {

		parts := strings.Split(command, " ")
		move, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		// part 1
		//for i := 0; i < move; i++ {
		//	crateToMove := crateMap[from][0:1]
		//	crateMap[to] = crateToMove + crateMap[to]
		//	crateMap[from] = strings.Replace(crateMap[from], crateToMove, "", 1)
		//}

		// part 2
		cratesToMove := crateMap[from][0:move]
		crateMap[to] = cratesToMove + crateMap[to]
		crateMap[from] = strings.Replace(crateMap[from], cratesToMove, "", 1)
	}

	fmt.Println(crateMap)
	topCrates := ""
	for z := 1; z <= 9; z++ {
		topCrates += crateMap[z][0:1]
	}

	fmt.Println(topCrates)

}

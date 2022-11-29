package main

import (
	"fmt"
)

func day1() {
	arr := [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	var lastNumber int

	for i := 0; i < len(arr); i++ {
		var currentNumber = arr[i]
		if i == 0 {
			fmt.Println("N/A - no previous measurement")
		} else if currentNumber > lastNumber {
			fmt.Println("Increasing")
		} else if currentNumber < lastNumber {
			fmt.Println("Decreasing")
		}
		lastNumber = arr[i]
	}
}
func main() {
	day1()
}

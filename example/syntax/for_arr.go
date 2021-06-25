package main

import "fmt"

func forArr() {
	var scores [10]int
	var sum = 0

	for i := 0; i < len(scores); i++ {
		fmt.Print(i, scores[i], "   ")
	}
	fmt.Println()

	for score := range scores {
		sum += score
	}

	fmt.Println(sum)
}

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

	var values = [...]string{"", "", ""}

	values[0] = "1"
	values[1] = "2"
	values[2] = "3"

	fmt.Println(values)
}

package day1

import (
	"fmt"
)

func SolveTask2() int {
	input := readInput()

	increased := 0
	for i, _ := range input[:len(input)-3] {
		previous := input[i] + input[i+1] + input[i+2]
		next := input[i+1] + input[i+2] + input[i+3]

		if next > previous {
			increased++
		}
	}

	fmt.Printf("Day 1 Task 2 Output: %d\n", increased)
	return increased
}

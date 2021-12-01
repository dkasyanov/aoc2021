package day1

import (
	"fmt"
)

func SolveTask1() int {
	input := readInput()

	increased := 0
	previous := input[0]
	for _, el := range input[1:] {
		if el > previous {
			increased++
		}
		previous = el
	}

	fmt.Printf("Day 1 Task 1 Output: %d\n", increased)
	return increased
}

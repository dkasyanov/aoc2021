package day2

import "fmt"

type position1 struct {
	Horizontal int
	Vertical   int
}

func SolveTask1() int {
	input := readInput()

	pos := position1{
		Horizontal: 0,
		Vertical:   0,
	}

	for _, el := range input {
		switch el.Direction {
		case "forward":
			pos.Horizontal += el.Value
			break
		case "up":
			pos.Vertical -= el.Value
			break
		case "down":
			pos.Vertical += el.Value
			break
		}
	}

	result := pos.Horizontal * pos.Vertical

	fmt.Printf("Day 2 Task 1 Output: %d\n", result)

	return result

}

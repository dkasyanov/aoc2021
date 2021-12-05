package day2

import "fmt"

type position2 struct {
	Horizontal int
	Vertical   int
	Aim        int
}

func SolveTask2() int {
	input := readInput()

	pos := position2{
		Horizontal: 0,
		Vertical:   0,
		Aim:        0,
	}

	for _, el := range input {
		switch el.Direction {
		case "forward":
			pos.Horizontal += el.Value
			pos.Vertical += el.Value * pos.Aim
			break
		case "up":
			pos.Aim -= el.Value
			break
		case "down":
			pos.Aim += el.Value
			break
		}
	}

	result := pos.Horizontal * pos.Vertical

	fmt.Printf("Day 2 Task 2 Output: %d\n", result)

	return result

}

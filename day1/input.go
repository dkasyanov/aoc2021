package day1

import (
	"bufio"
	"os"
	"strconv"
)

func readInput() []int {
	file, err := os.Open("day1/task1_input.txt")
	if err != nil {
		panic("Cannot open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Cannot parse input value")
		}

		data = append(data, num)
	}
	return data
}

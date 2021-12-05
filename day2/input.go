package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Direction string
	Value     int
}

func readInput() []Input {
	file, err := os.Open("day2/task1_input.txt")
	if err != nil {
		panic("Cannot open input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []Input
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		value, err := strconv.Atoi(input[1])
		if err != nil {
			panic("Cannot parse value")
		}
		data = append(data, Input{
			Direction: input[0],
			Value:     value,
		})
	}
	return data
}

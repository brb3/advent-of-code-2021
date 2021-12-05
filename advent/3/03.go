package main

import (
	"advent_utils"
	"fmt"
	"strings"
)

func main() {
	lines := advent_utils.LoadData("input.txt")
	fmt.Println("Part One:", partOne(lines))
}

func partOne(lines []string) int {
	var gamma int
	var epsilon int

	lineLength := len(lines[0])
	rowCount := len(lines)
	transposed := make([]string, lineLength)
	for i := 0; i < lineLength; i++ {
		for row := 0; row < rowCount; row++ {
			transposed[i] += string(lines[row][i])
		}
	}

	transposedLength := len(transposed)

	for i, line := range transposed {
		if strings.Count(line, "0") >= (len(lines) / 2) {
			gamma |= 1 << (transposedLength - 1 - i)
		} else {
			epsilon |= 1 << (transposedLength - 1 - i)
		}
	}

	return gamma * epsilon
}

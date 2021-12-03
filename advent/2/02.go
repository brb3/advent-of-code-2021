package main

import (
	"advent_utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := advent_utils.LoadData("input.txt")
	fmt.Println("Part One:", partOne(lines))
}

func partOne(input []string) int {
	horizontal := 0
	depth := 0

	for i := 0; i < len(input); i++ {
		split := strings.Split(input[i], " ")
		value, _ := strconv.ParseInt(split[1], 10, 0)
		switch split[0] {
		case "forward":
			horizontal += int(value)
		case "down":
			depth += int(value)
		case "up":
			depth -= int(value)
		}
	}

	return horizontal * depth
}

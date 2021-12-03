package main

import (
	"advent_utils"
	"fmt"
)

func main() {
	lines := advent_utils.ConvertStringArrayToIntArray(advent_utils.LoadData("input.txt"))
	fmt.Println("Part One:", partOne(lines))
	fmt.Println("Part Two:", partTwo(lines))
}

// Compares each line of `lines` with the previous line.
// Returns the amount of times the value increases.
func partOne(lines []int) int {
	var increases int

	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			increases++
		}
	}

	return increases
}

// Reads each line of `lines` and adds the next two lines to create a sliding window of measurements.
// Returns the
func partTwo(lines []int) int {
	var sliding []int
	var sum int

	for i := 0; i < len(lines); i++ {
		if i+2 > len(lines)-1 {
			break
		}
		sum += lines[i] + lines[i+1] + lines[i+2]
		sliding = append(sliding, sum)
		sum = 0
	}

	return partOne(sliding)
}

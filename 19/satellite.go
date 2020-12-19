package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Found %v valid messaged in input.\n", part1(SatelliteInput))

	fmt.Println("Part 2:")
	fmt.Printf("Found %v valid messaged in input.\n", part2(SatelliteInput))
}

func part1(input string) int {
	reg, messages := parseInput(input)

	valid := 0

	for _, m := range messages {

		if reg.MatchString(m) {
			valid++
		}
	}

	return valid
}

func part2(input string) int {
	input = strings.Replace(input, "8: 42", "8: 42 | 42 8", 1)
	input = strings.Replace(input, "11: 42 31", "11: 42 31 | 42 11 31", 1)

	reg, messages := parseInput(input)

	valid := 0

	for _, m := range messages {

		if reg.MatchString(m) {
			valid++
		}
	}

	return valid
}

package main

import (
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	var cubesSlice [][]int

	for i, l := range lines {

		if cubesSlice == nil {
			cubesSlice = make([][]int, len(lines))
		}

		for j, r := range l {

			if cubesSlice[i] == nil {
				cubesSlice[i] = make([]int, len(l))
			}

			active := 0

			if r == '#' {
				active = 1
			}

			cubesSlice[i][j] = active
		}
	}

	return cubesSlice
}

var ConwayCubesInput = `##.#....
...#...#
.#.#.##.
..#.#...
.###....
.##.#...
#.##..##
#.####..`

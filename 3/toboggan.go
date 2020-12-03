package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Forrest has %v rows\n", len(ForrestSlice))
	fmt.Printf("Forrest Slice has %v cols\n", len(ForrestSlice[0]))

	var part1Gradient = [2]int{3,1}

	part1Encounters := countTreeEncounters(ForrestSlice, part1Gradient)

	fmt.Println("Part 1:")
	fmt.Printf("Encountered %v tress\n", part1Encounters)

	// -------

	part2Gradients := [][2]int{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}

	product := multiplyTreeCounts(ForrestSlice, part2Gradients)

	fmt.Println("Part 2:")
	fmt.Printf("Product of encountered tress is: %v\n", product)
}

func countTreeEncounters(forrestSlice []string, gradient [2]int) int {
	acrossStep := gradient[0]
	downStep := gradient[1]

	forrestSliceWidth := len(forrestSlice[0])

	across := 0
	count := 0

	for rowIndex, row := range forrestSlice {

		// fmt.Printf("Row: %v\n", row);
		
		if rowIndex % downStep != 0 {
			continue
		}

		// Fake repeating section with modulo arithmetic
		var sliceAcross int

		if across == 0 {
			sliceAcross = 0
		} else {
			sliceAcross = across % forrestSliceWidth
		}

		// fmt.Printf("Found: %v at row %v, col %v\n\n", string(row[sliceAcross]), rowIndex, sliceAcross);

		if string(row[sliceAcross]) == "#" {
			count++
		}

		across += acrossStep
	}

	return count
}

func multiplyTreeCounts(forrestSlice []string, gradients [][2]int) int {
	runningProduct := 1

	for _, gradient := range gradients {
		runningProduct *= countTreeEncounters(forrestSlice, gradient)
	}

	return runningProduct
}

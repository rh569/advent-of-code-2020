package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Active after 6 cycles: %v\n", part1(ConwayCubesInput))

	fmt.Println("Part 2:")
	fmt.Printf("Active after 6 cycles: %v\n", part2(ConwayCubesInput))
}

func part1(input string) int {
	startingCubes := parseInput(input)

	cycles := 6
	//can only expand out by 1 shell each cycle, so a dimension len(input) + 2*cycles should be fine
	dimensionSize := len(startingCubes) + 2*cycles

	pokédim := initialisePocketDim(startingCubes, dimensionSize)

	for c := 0; c < cycles; c++ {
		pokédim = doCycle(pokédim)
	}

	return countActive(pokédim)
}

// Create a 3D array and copy the given starting slice to its centre
func initialisePocketDim(startingSlice [][]int, sizeOverride int) [][][]int {
	const DEFAULT_POCKET_SIZE = 24
	size := DEFAULT_POCKET_SIZE

	if sizeOverride > 0 {
		size = sizeOverride
	}

	if len(startingSlice) >= size {
		log.Fatalf("Dimension not big enough (%v) for starting slice (%v)\n", size, len(startingSlice))
	}

	pocketDimension := makeEmptyDimension(size)

	// copy in slice
	// assumes slice is square, but does not have to be odd/even in length

	halfDim := size / 2
	halfSlice := len(startingSlice) / 2
	startPos := halfDim - halfSlice

	// (z,) y, x refer to dim; i, j to the slice
	for y, i := startPos, 0; i < len(startingSlice); i++ {

		for x, j := startPos, 0; j < len(startingSlice); j++ {
			pocketDimension[halfDim][y][x] = startingSlice[i][j]
			x++
		}

		y++
	}

	return pocketDimension
}

func makeEmptyDimension(size int) [][][]int {
	pocketDimension := make([][][]int, size)

	// index by z slice first
	for z := range pocketDimension {
		pocketDimension[z] = make([][]int, size)

		// then y
		for y := range pocketDimension[z] {
			pocketDimension[z][y] = make([]int, size)

			// then x
			for x := range pocketDimension[z][y] {
				pocketDimension[z][y][x] = 0
			}
		}
	}

	return pocketDimension
}

// Takes in one generation of the dimension and outputs the next
func doCycle(dimIn [][][]int) [][][]int {
	dimOut := makeEmptyDimension(len(dimIn))

	// index by z
	for z := range dimIn {
		// then y
		for y := range dimIn[z] {
			// then x
			for x := range dimIn[z][y] {
				neighbours := countNeighbours(dimIn, z, y, x)

				// Only setting 1s here, new dim is initialised to 0s
				if dimIn[z][y][x] == 0 {
					// inactive
					if neighbours == 3 {
						dimOut[z][y][x] = 1
					}
				} else {
					// active
					if neighbours == 2 || neighbours == 3 {
						dimOut[z][y][x] = 1
					}
				}
			}
		}
	}

	return dimOut
}

func countNeighbours(dimIn [][][]int, z1, y1, x1 int) int {
	count := 0
	size := len(dimIn)

	for z := z1 - 1; z <= z1+1; z++ {

		if z < 0 || z >= size {
			continue
		}

		for y := y1 - 1; y <= y1+1; y++ {

			if y < 0 || y >= size {
				continue
			}

			for x := x1 - 1; x <= x1+1; x++ {

				if x < 0 || x >= size {
					continue
				}

				// don't count ourselves
				if z == z1 && y == y1 && x == x1 {
					continue
				}

				if dimIn[z][y][x] == 1 {
					count++
				}
			}
		}
	}

	return count
}

func countActive(dimIn [][][]int) int {
	active := 0

	// index by z
	for z := range dimIn {
		// then y
		for y := range dimIn[z] {
			// then x
			for x := range dimIn[z][y] {

				if dimIn[z][y][x] == 1 {
					active++
				}
			}
		}
	}

	return active
}

// Print each k/z slice of the dimension
func printByZ(dim [][][]int) {
	halfDim := len(dim)/2 + 1

	for z, slice := range dim {
		fmt.Printf("\nZ offset = %v: (z=%v)\n", z+1-halfDim, z)

		for _, row := range slice {

			for _, cube := range row {

				if cube == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}

			fmt.Print("\n")
		}
	}
}

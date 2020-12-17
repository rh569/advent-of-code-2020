package main

import "log"

func part2(input string) int {
	startingCubes := parseInput(input)

	cycles := 6
	//can only expand out by 1 shell each cycle, so a dimension len(input) + 2*cycles should be fine
	dimensionSize := len(startingCubes) + 2*cycles

	pokédim := initialisePocket4Dim(startingCubes, dimensionSize)

	for c := 0; c < cycles; c++ {
		pokédim = do4Cycle(pokédim)
	}

	return count4Active(pokédim)
}

// Create a 4D array and copy the given starting slice to its centre
func initialisePocket4Dim(startingSlice [][]int, sizeOverride int) [][][][]int {
	const DEFAULT_POCKET_SIZE = 24
	size := DEFAULT_POCKET_SIZE

	if sizeOverride > 0 {
		size = sizeOverride
	}

	if len(startingSlice) >= size {
		log.Fatalf("Dimension not big enough (%v) for starting slice (%v)\n", size, len(startingSlice))
	}

	pocketDimension := makeEmpty4Dimension(size)

	// copy in slice
	// assumes slice is square, but does not have to be odd/even in length

	halfDim := size / 2
	halfSlice := len(startingSlice) / 2
	startPos := halfDim - halfSlice

	// (z, y,) x, w refer to dim; i, j to the slice
	for x, i := startPos, 0; i < len(startingSlice); i++ {

		for w, j := startPos, 0; j < len(startingSlice); j++ {
			pocketDimension[halfDim][halfDim][x][w] = startingSlice[i][j]
			w++
		}

		x++
	}

	return pocketDimension
}

func makeEmpty4Dimension(size int) [][][][]int {
	pocketDimension := make([][][][]int, size)

	// index by z slice first
	for z := range pocketDimension {
		pocketDimension[z] = make([][][]int, size)

		// then y
		for y := range pocketDimension[z] {
			pocketDimension[z][y] = make([][]int, size)

			// then x
			for x := range pocketDimension[z][y] {
				pocketDimension[z][y][x] = make([]int, size)

				// then w
				for w := range pocketDimension[z][y][x] {
					pocketDimension[z][y][x][w] = 0

				}
			}
		}
	}

	return pocketDimension
}

// Takes in one generation of the dimension and outputs the next
func do4Cycle(dimIn [][][][]int) [][][][]int {
	dimOut := makeEmpty4Dimension(len(dimIn))

	// index by z
	for z := range dimIn {
		// then y
		for y := range dimIn[z] {
			// then x
			for x := range dimIn[z][y] {
				// then w
				for w := range dimIn[z][y][x] {
					neighbours := count4Neighbours(dimIn, z, y, x, w)

					// Only setting 1s here, new dim is initialised to 0s
					if dimIn[z][y][x][w] == 0 {
						// inactive
						if neighbours == 3 {
							dimOut[z][y][x][w] = 1
						}
					} else {
						// active
						if neighbours == 2 || neighbours == 3 {
							dimOut[z][y][x][w] = 1
						}
					}
				}
			}
		}
	}

	return dimOut
}

func count4Neighbours(dimIn [][][][]int, z1, y1, x1, w1 int) int {
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

				for w := w1 - 1; w <= w1+1; w++ {

					if w < 0 || w >= size {
						continue
					}

					// don't count ourselves
					if z == z1 && y == y1 && x == x1 && w == w1 {
						continue
					}

					if dimIn[z][y][x][w] == 1 {
						count++
					}
				}
			}
		}
	}

	return count
}

func count4Active(dimIn [][][][]int) int {
	active := 0

	// index by z
	for z := range dimIn {
		// then y
		for y := range dimIn[z] {
			// then x
			for x := range dimIn[z][y] {
				// then w
				for w := range dimIn[z][y][x] {

					if dimIn[z][y][x][w] == 1 {
						active++
					}
				}
			}
		}
	}

	return active
}

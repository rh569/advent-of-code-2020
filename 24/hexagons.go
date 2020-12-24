package main

import "fmt"

const (
	WHITE TileState = iota
	BLACK
)

type TileState int

type FloorState map[Vector]TileState

type Vector struct {
	x, y int
}

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Tiles that are black: %v\n", part1(TilesInput))

	fmt.Println("Part 2:")
	fmt.Printf("Tiles that are black after 100 days: %v\n", part2(TilesInput))
}

func part1(input string) int {
	instructions := parseInput(input)
	lobbyFloor := make(FloorState)

	lobbyFloor = tileFloor(lobbyFloor, instructions)

	return countBlackTiles(lobbyFloor)
}

func part2(input string) int {
	// Set up the initial floor state
	instructions := parseInput(input)
	lobbyFloor := make(FloorState)
	lobbyFloor = tileFloor(lobbyFloor, instructions)

	// Automaton-ise
	for i := 0; i < 100; i++ {
		lobbyFloor = nextGeneration(lobbyFloor)
	}

	return countBlackTiles(lobbyFloor)
}

func tileFloor(floor FloorState, instructions []TilingInstruction) FloorState {

	for _, ins := range instructions {
		floor = executeInstruction(floor, ins)
	}

	return floor
}

// Flip a single tile by following the TilingInstruction's Directions
func executeInstruction(floor FloorState, instruction TilingInstruction) FloorState {
	pos := Vector{0, 0}

	for _, dir := range instruction {
		switch dir {
		case EAST:
			pos.x += 2
		case SOUTH_EAST:
			pos.x += 1
			pos.y += 1
		case SOUTH_WEST:
			pos.x += -1
			pos.y += 1
		case WEST:
			pos.x += -2
		case NORTH_WEST:
			pos.x += -1
			pos.y += -1
		case NORTH_EAST:
			pos.x += 1
			pos.y += -1
		}
	}

	state, set := floor[pos]

	if set {
		floor[pos] = BLACK - state // toggle state
	} else {
		floor[pos] = BLACK
	}

	return floor
}

func countBlackTiles(floor FloorState) int {
	count := 0

	for _, v := range floor {
		if v == BLACK {
			count++
		}
	}

	return count
}

//---

func nextGeneration(floor FloorState) FloorState {
	// can't just iterate over map values now, as there are many un-set tiles
	// so, we'll loop over the keys once to get the bounds of the floor
	// then loop over every tile in bounds
	var largestPosX, largestNegX, largestPosY, largestNegY int

	for k, _ := range floor {
		if k.x > largestPosX {
			largestPosX = k.x
		}

		if k.x < largestNegX {
			largestNegX = k.x
		}

		if k.y > largestPosY {
			largestPosY = k.y
		}

		if k.y < largestNegY {
			largestNegY = k.y
		}
	}

	// Increase bounds by one tile in every dir
	largestPosX += 2
	largestNegX += -2
	largestPosY += 1
	largestNegY += -1

	// Set a new state
	newFloor := make(FloorState)

	for x := largestNegX; x <= largestPosX; x++ {
		for y := largestNegY; y <= largestPosY; y++ {
			pos := Vector{x, y}
			state, set := floor[pos]
			numNeighbours := countAdjacent(floor, pos)

			if set && state == BLACK {
				if numNeighbours == 0 || numNeighbours > 2 {
					newFloor[pos] = WHITE
				} else {
					newFloor[pos] = BLACK
				}
			} else {
				if numNeighbours == 2 {
					newFloor[pos] = BLACK
				} else {
					newFloor[pos] = WHITE
				}
			}
		}
	}

	return newFloor
}

// Counts adjacent black tiles
func countAdjacent(floor FloorState, pos Vector) int {

	neighbours := make([]Vector, 6)

	neighbours[EAST] = Vector{pos.x + 2, pos.y}
	neighbours[SOUTH_EAST] = Vector{pos.x + 1, pos.y + 1}
	neighbours[SOUTH_WEST] = Vector{pos.x - 1, pos.y + 1}
	neighbours[WEST] = Vector{pos.x - 2, pos.y}
	neighbours[NORTH_WEST] = Vector{pos.x - 1, pos.y - 1}
	neighbours[NORTH_EAST] = Vector{pos.x + 1, pos.y - 1}

	count := 0

	for _, nPos := range neighbours {
		state, set := floor[nPos]

		if set && state == BLACK {
			count++
		}
	}

	return count
}

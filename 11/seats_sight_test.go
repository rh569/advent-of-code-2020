package main

import (
	"testing"
)

var newGen1Input = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

var newGen2Input = `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`

var newGen3Input = `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`

var oneUnoccupied = `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`

var eightOccupied = `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`

var oneEmptySeat = `.............
.L.L.#.#.#.#.
.............`

// ---- Neighbours (Sight)

func TestNumOccupiedNeighboursGen0Sight(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 0

	result := seatMap.seats[0][0].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1CornerSight(t *testing.T) {
	seatMap := parseInput(newGen1Input)
	want := 3

	result := seatMap.seats[0][0].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1EdgeSight(t *testing.T) {
	seatMap := parseInput(newGen1Input)
	want := 5

	result := seatMap.seats[4][9].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1MiddleSight(t *testing.T) {
	seatMap := parseInput(newGen1Input)
	want := 8

	result := seatMap.seats[3][3].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursOneOcc(t *testing.T) {
	seatMap := parseInput(oneUnoccupied)
	want := 0

	result := seatMap.seats[3][3].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursEightOcc(t *testing.T) {
	seatMap := parseInput(eightOccupied)
	want := 8

	result := seatMap.seats[4][3].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursOneEmpty(t *testing.T) {
	seatMap := parseInput(oneEmptySeat)
	want := 0

	result := seatMap.seats[1][1].numOccupiedNeighbours(true)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

// ---- Next Gen

func TestNextGenerationGen0Sight(t *testing.T) {
	seatMap := parseInput(gen0Input)

	seatMap.nextGeneration(true)

	result := seatMap.toString()

	if result != newGen1Input {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, newGen1Input)
	}
}

func TestNextGenerationGen1Sight(t *testing.T) {
	seatMap := parseInput(newGen1Input)

	seatMap.nextGeneration(true)

	result := seatMap.toString()

	if result != newGen2Input {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, newGen2Input)
	}
}

func TestNextGenerationGen2Sight(t *testing.T) {
	seatMap := parseInput(newGen2Input)

	seatMap.nextGeneration(true)

	result := seatMap.toString()

	if result != newGen3Input {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, newGen3Input)
	}
}

// ---- Iterations

func TestIterationsUntilStableSight(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 6

	result := seatMap.iterationsUntilStable(10, true)

	if result != want {
		t.Fatalf("Iterations: %v Want: %v\n", result, want)
	}
}

func TestStableGeneration(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`

	seatMap.iterationsUntilStable(10, true)

	result := seatMap.toString()

	if result != want {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, want)
	}
}

func TestOccupiedWhenStableSight(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 26

	seatMap.iterationsUntilStable(10, true)

	result := seatMap.numOccupied()

	if result != want {
		t.Fatalf("Occupied: %v Want: %v\n", result, want)
	}
}

var nonSquareInput = `####
####
####`

func TestFullInputGen(t *testing.T) {
	seatMap := parseInput(nonSquareInput)

	want := `#LL#
LLLL
#LL#`

	seatMap.iterationsUntilStable(10, true)

	result := seatMap.toString()

	if result != want {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, want)
	}
}

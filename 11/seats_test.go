package main

import (
	"testing"
)

var gen0Input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

var gen1Input = `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

var gen2Input = `#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`

func TestParseInput(t *testing.T) {
	seatMap := parseInput(gen0Input)
	wantRows := 10
	wantCols := 10

	result := seatMap.toString()

	if result != gen0Input {
		t.Fatalf("Seat map:\n%v\nDifferent to input:\n%v\n", result, gen0Input)
	}

	if wantRows != seatMap.numRows() || wantCols != seatMap.numCols() {
		t.Fatalf("Size of seatmap is (col x row) %v x %v, want %v x %v", seatMap.numCols(), seatMap.numRows(), wantCols, wantRows)
	}
}

// ---- Neighbours

func TestNumOccupiedNeighboursGen0(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 0

	result := seatMap.seats[0][0].numOccupiedNeighbours(false)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1Corner(t *testing.T) {
	seatMap := parseInput(gen1Input)
	want := 2

	result := seatMap.seats[0][0].numOccupiedNeighbours(false)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1Edge(t *testing.T) {
	seatMap := parseInput(gen1Input)
	want := 5

	result := seatMap.seats[4][9].numOccupiedNeighbours(false)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

func TestNumOccupiedNeighboursGen1Middle(t *testing.T) {
	seatMap := parseInput(gen1Input)
	want := 8

	result := seatMap.seats[8][3].numOccupiedNeighbours(false)

	if result != want {
		t.Fatalf("Neighbours Occupied: %v Want: %v\n", result, want)
	}
}

// ---- Next Gen

func TestNextGenerationGen0(t *testing.T) {
	seatMap := parseInput(gen0Input)

	seatMap.nextGeneration(false)

	result := seatMap.toString()

	if result != gen1Input {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, gen1Input)
	}
}

func TestNextGenerationGen1(t *testing.T) {
	seatMap := parseInput(gen1Input)

	seatMap.nextGeneration(false)

	result := seatMap.toString()

	if result != gen2Input {
		t.Fatalf("Seat map:\n%v\nDifferent to want:\n%v\n", result, gen2Input)
	}
}

// ---- Iterations

func TestIterationsUntilStable(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 5

	result := seatMap.iterationsUntilStable(10, false)

	if result != want {
		t.Fatalf("Iterations: %v Want: %v\n", result, want)
	}
}

func TestOccupiedWhenStable(t *testing.T) {
	seatMap := parseInput(gen0Input)
	want := 37

	seatMap.iterationsUntilStable(10, false)

	result := seatMap.numOccupied()

	if result != want {
		t.Fatalf("Occupied: %v Want: %v\n", result, want)
	}
}

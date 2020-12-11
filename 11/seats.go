package main

import (
	"fmt"
	"strings"
)

func main() {
	seatMap := parseInput(SeatLayoutInput)
	iterations := seatMap.iterationsUntilStable(1000, false)
	stableOccupied := seatMap.numOccupied()

	fmt.Println("Part 1:")
	fmt.Printf("Iterations %v. Occupied %v\n", iterations, stableOccupied)

	newSeatMap := parseInput(SeatLayoutInput)
	newIterations := newSeatMap.iterationsUntilStable(1000, true)
	newStableOccupied := newSeatMap.numOccupied()

	fmt.Println("Part 2:")
	fmt.Printf("Iterations %v. Occupied %v\n", newIterations, newStableOccupied)
}

func parseInput(input string) *SeatMap {
	const L_RUNE = 76
	const HASH_RUNE = 35

	lines := strings.Split(input, "\n")

	seatMap := SeatMap{}
	seats := make([][]Seat, len(lines))

	for i, line := range lines {

		seats[i] = make([]Seat, len(line))

		for j, r := range line {

			var seat Seat

			if r == L_RUNE || r == HASH_RUNE {
				seat.comfy = true
				seat.row = i
				seat.col = j
				seat.seatMap = &seatMap
			}

			if r == HASH_RUNE {
				seat.occupied = true
			}

			seats[i][j] = seat
		}
	}

	seatMap.seats = seats
	return &seatMap
}

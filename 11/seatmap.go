package main

import (
	"fmt"
	"log"
)

// Represents all seats in a grid (2D Array)
// seats[row][col]
// floor is nil
type SeatMap struct {
	seats [][]Seat
}

func (s *SeatMap) numRows() int {
	return len(s.seats)
}

func (s *SeatMap) numCols() int {
	return len(s.seats[0])
}

func (s *SeatMap) numOccupied() int {
	numOcc := 0

	// Mark seats for change
	for i := 0; i < s.numRows(); i++ {
		for j := 0; j < s.numCols(); j++ {

			if s.seats[i][j].occupied {
				numOcc++
			}
		}
	}

	return numOcc
}

// Returns number of seats changed
func (s *SeatMap) nextGeneration(sightlineMethod bool) int {
	changed := 0

	// Mark seats for change
	for i := 0; i < s.numRows(); i++ {
		for j := 0; j < s.numCols(); j++ {

			if s.seats[i][j].determineWillChange(sightlineMethod) {
				changed++
			}
		}
	}

	// Update all seats
	for i := 0; i < s.numRows(); i++ {
		for j := 0; j < s.numCols(); j++ {
			s.seats[i][j].update()
		}
	}

	return changed
}

func (s *SeatMap) iterationsUntilStable(maxIterations int, sightlineMethod bool) int {
	iterations := 0

	for iterations <= maxIterations {
		changed := s.nextGeneration(sightlineMethod)

		if changed == 0 {
			return iterations
		}

		iterations++
	}

	// Shouldn't get here
	return -1
}

func (s *SeatMap) toString() string {
	str := ""

	for i := 0; i < s.numRows(); i++ {
		for j := 0; j < s.numCols(); j++ {

			if !s.seats[i][j].comfy {
				str += "."
			} else {

				if s.seats[i][j].occupied {
					str += "#"
				} else {
					str += "L"
				}
			}
		}

		if i < s.numRows()-1 {
			str += "\n"
		}
	}

	return str
}

func (s *SeatMap) print() {
	fmt.Printf("%v\n", s.toString())
}

// Represents a single seat's state
type Seat struct {
	comfy      bool
	occupied   bool
	row        int
	col        int
	willChange bool
	seatMap    *SeatMap
}

func (s *Seat) numOccupiedNeighbours(sightlineMethod bool) int {
	if s.seatMap == nil {
		log.Fatalf("Don't check the floor for neighbours")
	}

	if sightlineMethod {
		return numOccupiedSightline(s)
	}

	return numOccupiedAdjacent(s)
}

// Checks first seat along each sightline to see if occupied
func numOccupiedSightline(s *Seat) int {
	nOcc := 0

	grads := [][2]int{
		{-1, 0},  // N
		{-1, 1},  // NE
		{0, 1},   // E
		{1, 1},   // SE
		{1, 0},   // S
		{1, -1},  // SW
		{0, -1},  // W
		{-1, -1}, // NW
	}

	for _, grad := range grads {
		seat := findFirstSeat(s.seatMap.seats, grad, s.row, s.col)

		if !seat.comfy {
			continue
		}

		if seat.occupied {
			nOcc++
		}
	}

	return nOcc
}

// Returns the first non-floor / seat in the given direction (gradient)
func findFirstSeat(seats [][]Seat, grad [2]int, startRow int, startCol int) Seat {
	rowStep := grad[0]
	colStep := grad[1]

	if rowStep == 0 && colStep == 0 {
		panic("Can't have gradient of 0,0")
	}

	row := startRow
	col := startCol

	for row >= 0 && row < len(seats) && col >= 0 && col < len(seats[0]) {
		// Not where we started
		if row != startRow || col != startCol {

			// If not floor return seat
			if seats[row][col].comfy {
				return seats[row][col]
			}
		}

		col += colStep
		row += rowStep
	}

	return Seat{}
}

// Checks 8 adjacent seats to see if occupied
func numOccupiedAdjacent(s *Seat) int {
	nOcc := 0

	// For adjacent rows
	for i := s.row - 1; i < s.row+2; i++ {

		if i < 0 || i >= len(s.seatMap.seats) {
			continue
		}

		// For adjacent cols
		for j := s.col - 1; j < s.col+2; j++ {

			if j < 0 || j >= len(s.seatMap.seats[0]) {
				continue
			}

			// Don't count this seat
			if i == s.row && j == s.col {
				continue
			}

			// Don't check floor
			if !s.seatMap.seats[i][j].comfy {
				continue
			}

			if s.seatMap.seats[i][j].occupied {
				nOcc++
			}
		}
	}

	return nOcc
}

// Marks seat for change
// Also returns true if it is marked to change
func (s *Seat) determineWillChange(sightlineMethod bool) bool {
	if !s.comfy {
		return false
	}

	occThreshold := 3

	if sightlineMethod {
		occThreshold = 4
	}

	nOcc := s.numOccupiedNeighbours(sightlineMethod)

	if s.occupied && nOcc > occThreshold {
		s.willChange = true
	} else if !s.occupied && nOcc == 0 {
		s.willChange = true
	} else {
		s.willChange = false
	}

	return s.willChange
}

func (s *Seat) update() {
	if !s.comfy {
		return
	}

	if s.willChange {
		s.occupied = !s.occupied
	}
}

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	var highestId int
	var allIds = []int{}

	for _, pass := range Passes {
		row := findPassRow(pass[:len(pass)-3])
		col := findPassCol(pass[len(pass)-3:])

		id := calculatePassId(row, col)

		allIds = append(allIds, id)

		if id > highestId {
			highestId = id
		}
	}

	missingId, err := findMissingId(allIds)

	if err != nil {
		log.Fatalf("Error locating missing id %v", err)
	}

	fmt.Println("Part 1:")
	fmt.Printf("Highest id: %v\n", highestId)

	fmt.Println("Part 2:")
	fmt.Printf("Missing seat id: %v\n", missingId)

}

func calculatePassId(row int, col int) int {
	return row*8 + col
}

func findPassRow(first7 string) int {
	return binaryLocate(0, 127, first7)
}

func findPassCol(last3 string) int {
	return binaryLocate(0, 7, last3)
}

// Applies the low/high instruction to the range specified by low and high
// len(instruction) number of times and checks for low/high equality
// For instruction, F and L indicate lower, B and R indicate higher
// Returns number between low and high inclusive
func binaryLocate(low int, high int, instruction string) int {

	steps := len(instruction)

	if math.Exp2(float64(steps)) != float64(high-low+1) {
		log.Fatalf("%v steps covers search space of %v. Search is of a space of %v\n", steps, math.Exp2(float64(steps)), high-low+1)
	}

	for _, char := range instruction {
		letter := string(char)
		diff := high - low
		half := (diff + 1) / 2

		if letter == "F" || letter == "L" {
			high -= half
		} else if letter == "B" || letter == "R" {
			low += half
		} else {
			log.Fatalf("Invalid instruction char %v\n", letter)
		}
	}

	if low != high {
		log.Fatalf("Locate failed, low %v | high %v\n", low, high)
	}

	return low
}

func findMissingId(ids []int) (int, error) {
	var missingId int

	sort.Ints(ids)

	for i := 1; i < len(ids)-1; i++ {
		prevSeatId := ids[i-1]
		seatId := ids[i]

		if seatId-prevSeatId > 1 {
			
			if seatId-prevSeatId != 2 {
				return 0, errors.New(fmt.Sprintf("More than one seat in gap of size: %v", seatId-prevSeatId))
			}

			if missingId > 0 {
				return 0, errors.New(fmt.Sprintf("Already found a missing seat: %v", missingId))
			}

			missingId = prevSeatId + 1
		}
	}

	return missingId, nil
}

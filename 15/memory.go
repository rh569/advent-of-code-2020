package main

import (
	"fmt"
	"strconv"
	"strings"
)

var NumberInput = `10,16,6,0,1,17`

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("The 2020th number is: %v\n", part1())

	fmt.Println("Part 2:")
	fmt.Printf("The 2020th number is: %v\n", part2())
}

func part1() int {
	numbers := parseInput(NumberInput)

	return playGame(numbers, 2020)
}

func part2() int {
	numbers := parseInput(NumberInput)

	return playGame(numbers, 30000000)
}

// plays through the memory game for the given number of turns
// returns the spoken number for that turn
func playGame(startingNumbers []int, turns int) int {

	// map of the spoken numbers and the turns on which they were spoken
	spokenNumbers := make(map[int][]int)
	currentTurn := 1
	var lastSpoken int

	// Use the starting numbers
	for _, n := range startingNumbers {
		spokenNumbers[n] = []int{currentTurn}
		lastSpoken = n
		// fmt.Printf("Turn: %v, Say: %v\n", currentTurn, lastSpoken)
		currentTurn++
	}

	for i := currentTurn; i <= turns; i++ {
		prevTurns := spokenNumbers[lastSpoken]

		// Said once before, so say 0
		if len(prevTurns) == 1 {
			prevZeroes := spokenNumbers[0]
			prevZeroes = append(prevZeroes, currentTurn)
			spokenNumbers[0] = prevZeroes
			lastSpoken = 0
		}

		// Said before, so say difference
		if len(prevTurns) > 1 {
			diff := prevTurns[len(prevTurns)-1] - prevTurns[len(prevTurns)-2]
			prevDiff := spokenNumbers[diff]
			prevDiff = append(prevDiff, currentTurn)
			spokenNumbers[diff] = prevDiff
			lastSpoken = diff
		}

		// fmt.Printf("Turn: %v, Say: %v\n", currentTurn, lastSpoken)
		currentTurn++
	}

	return lastSpoken
}

func parseInput(input string) []int {
	numbers := []int{}
	numberStrings := strings.Split(input, ",")

	for _, n := range numberStrings {
		numbers = append(numbers, getInt(n))
	}

	return numbers
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

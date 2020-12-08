package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Boot Code Instruction
type Ins struct {
	code  string
	value int
}

func main() {
	instructions := parseInput(BootCodeInput)

	accumulator, finished := runInstructions(instructions)

	fmt.Println("Part 1:")

	if finished {
		log.Fatal("Reaches program end when it should have looped")
	}

	fmt.Printf("Program looped with accumulator: %v\n", accumulator)

	fmt.Println("Part 2 (Brute force):")
	newAccumulator, err := mutateInstructionsUntilDone(instructions)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("Successfully fixed program, acc: %v\n", newAccumulator)
}

// Parse the whole input returning an array of all instructions
func parseInput(bootCodeInput string) []Ins {
	lines := strings.Split(bootCodeInput, "\n")

	instructions := []Ins{}

	for _, insStr := range lines {
		instructions = append(instructions, parseInstruction(insStr))
	}

	return instructions
}

// Parse a single line of instruction
func parseInstruction(insStr string) Ins {
	var ins Ins

	ins.code = insStr[:3]
	ins.value = getSignedInt(insStr[4:])

	return ins
}

// Returns an int from string formatted as +x or -x where x is a positive int
func getSignedInt(numStr string) int {
	var num int
	sign := string(numStr[0])

	value, err := strconv.ParseInt(numStr[1:], 10, 0)

	if err != nil {
		log.Fatalf("Could not parse %v\nError: %v\n", numStr, err)
	}

	num = int(value)

	if sign == "-" {
		num *= -1
	}

	return num
}

// Run until instructions end or loop detected
// If loop detected, returns (acc, false) where acc is the value of
// the accumulator before the first repeated instruction
// If finished, returns (acc, true)
func runInstructions(instructions []Ins) (int, bool) {
	alreadyExecuted := []int{}
	var acc int

	// Simple for instead of range, so that we can execute jmp calls
	for pos := 0; pos < len(instructions); pos++ {

		if containsInt(alreadyExecuted, pos) {
			return acc, false
		}

		alreadyExecuted = append(alreadyExecuted, pos)
		ins := instructions[pos]

		switch ins.code {
		case "nop":
			continue
		case "acc":
			acc += ins.value
		case "jmp":
			pos += ins.value - 1 // loop increments by 1 anyway
		}
	}

	return acc, true
}

// Keep running the instruction set, switching nop and jmp instructions
// until we create a finishable program
// If we can finish, return the acc - otherwise error
func mutateInstructionsUntilDone(instructions []Ins) (int, error) {

	for pos, ins := range instructions {
		mutatedInstructions := []Ins{}

		mutatedInstructions = append(mutatedInstructions, instructions...)

		if ins.code == "jmp" {
			mutatedInstructions[pos].code = "nop"
		} else if ins.code == "nop" {
			mutatedInstructions[pos].code = "jmp"
		}

		acc, fin := runInstructions(mutatedInstructions)

		if fin {
			return acc, nil
		}
	}

	return 0, errors.New("Could not fix program")
}

// Returns true if the given collection of ints contains the target int
// Will sort the collection if not already sorted
func containsInt(collection []int, target int) bool {

	if len(collection) == 0 {
		return false
	}

	if !sort.IntsAreSorted(collection) {
		sort.Ints(collection)
	}

	pos := sort.SearchInts(collection, target)

	if pos == len(collection) {
		return false
	}

	return collection[pos] == target
}

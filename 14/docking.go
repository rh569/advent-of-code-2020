package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction struct {
	op      int    // 0 is mask update; 1 is memory write
	orMask  uint64 // 'XXX1X00X' -> '00010000' OR-wise bitmask
	andMask uint64 // 'XXX1X00X' -> '11111001' AND-wise bitmask
	addr    uint64 // address to write to
	val     uint64 // decimal value to write

	// Part 2
	xClearMask uint64
	floatMasks []uint64
}

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Init sum: %v\n", part1(DockingInitializationInput))

	fmt.Println("Part 2:")
	fmt.Printf("Init sum: %v\n", part2(DockingInitializationInput))
}

func part1(input string) uint64 {
	instructions := parseInput(input, false)
	memorySize := getMaxMemory(instructions)

	mem := make([]uint64, memorySize)
	mem = runInstructions(mem, instructions)

	return sumMemory(mem)
}

func part2(input string) uint64 {
	instructions := parseInput(input, true)

	// Use a map for the 'memory' this time, we'll be accessing only a fraction of the address space (2^36)
	mem := make(map[uint64]uint64)
	mem = runFloatingInstructions(mem, instructions)

	return sumMappedMemory(mem)
}

// This is horribly slow for large memory
func sumMemory(mem []uint64) uint64 {
	var total uint64 = 0

	for _, n := range mem {
		total += n
	}

	return total
}

// Sums all memory values in the given map
func sumMappedMemory(mem map[uint64]uint64) uint64 {
	var total uint64 = 0

	for _, v := range mem {
		total += v
	}

	return total
}

// returns the largest memory address referened in the given instructions
func getMaxMemory(instructions []Instruction) uint64 {
	max := uint64(0)

	for _, ins := range instructions {
		if ins.op == 0 {
			continue
		}

		if ins.addr > max {
			max = ins.addr
		}
	}

	return max + 1
}

// applies the given instructions to the provided memory array
// returns the modified memory
// it is the caller's responsibility to ensure the memory is large enough and initialized
func runInstructions(mem []uint64, instructions []Instruction) []uint64 {
	var orMask uint64
	var andMask uint64

	for _, ins := range instructions {

		if ins.op == 0 {
			orMask = ins.orMask
			andMask = ins.andMask
		} else {
			mem[ins.addr] = applyMasks(ins.val, orMask, andMask)
		}
	}

	return mem
}

// applies the given instructions to the provided memory array using the
// instruction masks on the memory addresses with 'floating' bits
// returns the modified memory
// it is the caller's responsibility to ensure the memory is large enough and initialized
func runFloatingInstructions(mem map[uint64]uint64, instructions []Instruction) map[uint64]uint64 {
	var orMask uint64
	var xClearMask uint64
	var floatMasks []uint64

	for _, ins := range instructions {

		if ins.op == 0 {
			orMask = ins.orMask
			xClearMask = ins.xClearMask
			floatMasks = ins.floatMasks
		} else {

			for _, floatMask := range floatMasks {
				mem[getAddress(ins.addr, orMask, xClearMask, floatMask)] = ins.val
			}
		}
	}
	return mem
}

func parseInput(input string, considerFloatingMasks bool) []Instruction {
	lines := strings.Split(input, "\n")

	instructions := []Instruction{}

	for _, l := range lines {
		ins := Instruction{}
		parts := strings.Split(l, " = ")

		if parts[0] == "mask" {
			mask := parts[1]
			orMaskStr := strings.ReplaceAll(mask, "X", "0")
			andMaskStr := strings.ReplaceAll(mask, "X", "1")

			orMask, _ := strconv.ParseUint(orMaskStr, 2, 64)
			andMask, _ := strconv.ParseUint(andMaskStr, 2, 64)

			ins.orMask = orMask
			ins.andMask = andMask

			if considerFloatingMasks {
				ins.xClearMask = getClearMask(mask)
				ins.floatMasks = getFloatMasks(mask)
			}
		} else {
			ins.op = 1
			addr, _ := strconv.ParseUint(parts[0][4:len(parts[0])-1], 10, 64)
			ins.addr = addr

			val, _ := strconv.ParseUint(parts[1], 10, 64)
			ins.val = val
		}

		instructions = append(instructions, ins)
	}

	return instructions
}

// returns all bits set except the Xs
func getClearMask(maskString string) uint64 {
	maskString = strings.ReplaceAll(maskString, "0", "1")
	maskString = strings.ReplaceAll(maskString, "X", "0")

	mask, _ := strconv.ParseUint(maskString, 2, 64)
	return mask
}

func getFloatMasks(maskString string) []uint64 {
	// Clear any set bits
	maskString = strings.ReplaceAll(maskString, "1", "0")

	numFloatingBits := strings.Count(maskString, "X")
	numMasks := int(math.Pow(2, float64(numFloatingBits)))

	floatMasks := make([]uint64, numMasks)

	// For 2^num(X) different floatMasks
	for i := 0; i < numMasks; i++ {

		replacementBits := []byte{}

		// Pad with zeroes
		for i := 0; i < numFloatingBits; i++ {
			replacementBits = append(replacementBits, '0')
		}

		// Get a binary representation of the bits to set (has value of iteration)
		// Not 0 led and in wrong order for string replacement
		iterationNumber := fmt.Sprintf("%b", i)

		// copy significant bits across
		copy(replacementBits[len(replacementBits)-len(iterationNumber):], iterationNumber)

		newMask := maskString

		// Replace each X in order with a digit from the iterationBits
		for _, b := range replacementBits {
			letter := string(b)

			newMask = strings.Replace(newMask, "X", letter, 1)
		}

		mask, _ := strconv.ParseUint(newMask, 2, 64)
		floatMasks[i] = mask
	}

	return floatMasks
}

// Applies the OR and AND masks to the given value
func applyMasks(val, orMask, andMask uint64) uint64 {
	val = val | orMask
	return val & andMask
}

func getAddress(addr, orMask, xClearMask, floatMask uint64) uint64 {
	addr = addr | orMask
	addr = addr & xClearMask
	return addr + floatMask
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

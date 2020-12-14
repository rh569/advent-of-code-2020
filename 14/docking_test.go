package main

import (
	"reflect"
	"testing"
)

var testInput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestApplyMasks(t *testing.T) {
	instructions := parseInput(testInput, false)

	setMaskIns := instructions[0]
	var want1 uint64 = 73
	var want2 uint64 = 101
	var want3 uint64 = 64

	result1 := applyMasks(instructions[1].val, setMaskIns.orMask, setMaskIns.andMask)
	result2 := applyMasks(instructions[2].val, setMaskIns.orMask, setMaskIns.andMask)
	result3 := applyMasks(instructions[3].val, setMaskIns.orMask, setMaskIns.andMask)

	if want1 != result1 || want2 != result2 || want3 != result3 {
		t.Fatalf("Results %v, %v and %v. Want %v, %v and %v\n", result1, result2, result3, want1, want2, want3)
	}
}

func TestPart1(t *testing.T) {
	var want uint64 = 165

	result := part1(testInput)

	if want != result {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

func TestGetFloatMasks(t *testing.T) {
	testMask := `00000000000000000000010000000001X0XX`

	want := []uint64{
		0b000000000000000000000000000000000000,
		0b000000000000000000000000000000000001,
		0b000000000000000000000000000000000010,
		0b000000000000000000000000000000000011,
		0b000000000000000000000000000000001000,
		0b000000000000000000000000000000001001,
		0b000000000000000000000000000000001010,
		0b000000000000000000000000000000001011,
	}

	result := getFloatMasks(testMask)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result %v, Want %v", result, want)
	}
}

var part2TestInput = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestPart2(t *testing.T) {
	var want uint64 = 208

	result := part2(part2TestInput)

	if want != result {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

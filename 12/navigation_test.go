package main

import (
	"reflect"
	"testing"
)

var testInput = `F10
N3
F7
R90
F11`

func TestParseInput(t *testing.T) {
	instructions := parseInput(testInput)

	want := []Instruction{
		Instruction{"F", 10},
		Instruction{"N", 3},
		Instruction{"F", 7},
		Instruction{"R", 90},
		Instruction{"F", 11},
	}

	if !reflect.DeepEqual(instructions, want) {
		t.Fatalf("Instructions: %v do not match wanted: %v", instructions, want)
	}
}

func TestManhattanDistance(t *testing.T) {
	ferry := Ferry{facing: 90}
	instructions := parseInput(testInput)

	ferry.processInstructions(instructions)

	want := 25

	result := calculateManhattanDistance(0, 0, ferry.longitude, ferry.latitude)

	if result != want {
		t.Fatalf("Distance: %v, Wanted: %v\n", result, want)
	}
}

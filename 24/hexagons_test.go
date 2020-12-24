package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	want := []TilingInstruction{
		[]Direction{SOUTH_EAST, SOUTH_EAST, NORTH_WEST, NORTH_EAST, NORTH_EAST, NORTH_EAST, WEST, SOUTH_EAST, EAST},
	}
	result := parseInput("sesenwnenenewsee")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v. Want: %v\n", result, want)
	}
}

func TestPart1(t *testing.T) {
	want := 10
	result := part1(TestInput)

	if result != want {
		t.Fatalf("Result: %v. Want: %v\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2208
	result := part2(TestInput)

	if result != want {
		t.Fatalf("Result: %v. Want: %v\n", result, want)
	}
}

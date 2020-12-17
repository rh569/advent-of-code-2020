package main

import (
	"reflect"
	"testing"
)

var testInput = `.#.
..#
###`

func TestParseInput(t *testing.T) {
	cubes := parseInput(testInput)

	want := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}

	if !reflect.DeepEqual(cubes, want) {
		t.Fatalf("Result: %v, want %v\n", cubes, want)
	}
}

func TestPart1(t *testing.T) {
	want := 112
	result := part1(testInput)

	if result != want {
		t.Fatalf("Result: %v, want %v\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	want := 848
	result := part2(testInput)

	if result != want {
		t.Fatalf("Result: %v, want %v\n", result, want)
	}
}

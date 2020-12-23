package main

import "testing"

func TestPart1Ten(t *testing.T) {
	want := "92658374"
	result := part1(TestInput, 10)

	if result != want {
		t.Fatalf("Result: %v, want %v\n", result, want)
	}
}

func TestPart1Hundred(t *testing.T) {
	want := "67384529"
	result := part1(TestInput, 100)

	if result != want {
		t.Fatalf("Result: %v, want %v\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	want := 149245887792
	result := part2(TestInput)

	if result != want {
		t.Fatalf("Result: %v, want %v\n", result, want)
	}
}

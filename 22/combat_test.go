package main

import "testing"

func TestPart1(t *testing.T) {
	want := 306
	result := part1(TestInput)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

func TestPart2(t *testing.T) {
	want := 291
	result := part2(TestInput)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

func TestPart2Loop(t *testing.T) {
	want := 105
	result := part2(LoopTestInput)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

func TestDeckPairsMatchTrue(t *testing.T) {
	want := true
	decksA := [2][]int{
		{3, 4, 6, 10},
		{2},
	}
	decksB := [2][]int{
		{3, 4, 6, 10},
		{2},
	}

	result := deckPairsMatch(decksA, decksB)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

func TestDeckPairsMatchFalseLength(t *testing.T) {
	want := false
	decksA := [2][]int{
		{3, 4, 6, 10},
		{2},
	}
	decksB := [2][]int{
		{3, 4, 6, 10, 1},
		{2},
	}

	result := deckPairsMatch(decksA, decksB)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

func TestDeckPairsMatchFalseValues(t *testing.T) {
	want := false
	decksA := [2][]int{
		{3, 4, 6, 10},
		{2},
	}
	decksB := [2][]int{
		{3, 4, 6, 9},
		{2},
	}

	result := deckPairsMatch(decksA, decksB)

	if result != want {
		t.Fatalf("Result %v, want %v\n", result, want)
	}
}

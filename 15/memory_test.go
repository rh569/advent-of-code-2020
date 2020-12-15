package main

import (
	"testing"
)

func TestPlayGame(t *testing.T) {
	testNumbers := []int{0, 3, 6}
	want := 0

	result := playGame(testNumbers, 10)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

func TestPlayGameTwo(t *testing.T) {
	testNumbers := []int{0, 3, 6}
	want := 436

	result := playGame(testNumbers, 2020)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

func TestPlayGameThree(t *testing.T) {
	testNumbers := []int{1, 3, 2}
	want := 1

	result := playGame(testNumbers, 2020)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

func TestPlayGameFour(t *testing.T) {
	testNumbers := []int{3, 1, 2}
	want := 1836

	result := playGame(testNumbers, 2020)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

// --- Part 2

func TestPlayGameBig(t *testing.T) {
	testNumbers := []int{0, 3, 6}
	want := 175594

	result := playGame(testNumbers, 30000000)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

func TestPlayGameBigTwo(t *testing.T) {
	testNumbers := []int{2, 3, 1}
	want := 6895259

	result := playGame(testNumbers, 30000000)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

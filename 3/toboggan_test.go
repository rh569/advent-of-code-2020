package main

import (
	"testing"
)

var testForrestSlice = []string {
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestCountTreeEncounters(t *testing.T) {
	testGradient := [2]int{3,1}
	want := 7

	encounters := countTreeEncounters(testForrestSlice, testGradient)

	if encounters != want {
		t.Fatalf("Counted %v trees, wanted %v", encounters, want)
	}
}

func TestMultiplyTreeCounts(t *testing.T) {
	testGradients := [][2]int{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}

	want := 336

	product := multiplyTreeCounts(testForrestSlice, testGradients)

	if product != want {
		t.Fatalf("Product was %v, wanted %v", product, want)
	}
}
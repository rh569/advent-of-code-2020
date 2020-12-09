package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	testInput := `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

	want := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	result := parseInput(testInput)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result %v, want %v", result, want)
	}
}

func TestFindContiguousSum(t *testing.T) {
	testInput := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	want := []int{15, 25, 47, 40}
	target := 127

	result, err := findContiguousSum(testInput, target)

	if err != nil {
		t.Fatalf("Encountered err %v, want %v", err, want)
	}

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result %v, want %v", result, want)
	}
}

func TestSumMinMaxElements(t *testing.T) {
	testInts := []int{15, 25, 47, 40}
	want := 62

	result := sumMinMaxElements(testInts)

	if result != want {
		t.Fatalf("Sum was %v, want %v", result, want)
	}
}

func TestFindNonSumElementExists(t *testing.T) {
	testInput := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	copyOfInput := make([]int, len(testInput))

	copy(copyOfInput, testInput)

	n := 5

	want := 127
	result, err := findNonSumElement(testInput, n)

	if err != nil {
		t.Fatalf("Encountered err %v, want %v", err, want)
	}

	if result != want {
		t.Fatalf("Result %v, want %v", result, want)
	}

	if !reflect.DeepEqual(testInput, copyOfInput) {
		t.Fatalf("Input was mutated: %v, became %v", copyOfInput, testInput)
	}
}

// Tests the exact ranges involved
func TestFindNonSumElementExistsAdjacent(t *testing.T) {
	testInput := []int{1, 20, 15, 25, 47, 40, 60, 55, 65, 100, 87}
	copyOfInput := make([]int, len(testInput))

	copy(copyOfInput, testInput)

	n := 5

	want := 87
	result, err := findNonSumElement(testInput, n)

	if err != nil {
		t.Fatalf("Encountered err %v, want %v", err, want)
	}

	if result != want {
		t.Fatalf("Result %v, want %v", result, want)
	}

	if !reflect.DeepEqual(testInput, copyOfInput) {
		t.Fatalf("Input was mutated: %v, became %v", copyOfInput, testInput)
	}
}

func TestFindNonSumElementDoesntExist(t *testing.T) {
	testInput := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182}
	copyOfInput := make([]int, len(testInput))

	copy(copyOfInput, testInput)

	n := 5

	result, err := findNonSumElement(testInput, n)

	if err == nil {
		t.Fatalf("Expected error, found %v", result)
	}

	if !reflect.DeepEqual(testInput, copyOfInput) {
		t.Fatalf("Input was mutated: %v, became %v", copyOfInput, testInput)
	}
}

func TestHasSumInIntsTrue(t *testing.T) {
	coll := []int{4, 6, 3, 5, 8, 9, 4, 23}
	target := 29
	want := true

	result := hasSumInInts(coll, target)

	if result != want {
		t.Fatalf("Did not find sum to %v in %v. Result %v, want %v", target, coll, result, want)
	}
}

func TestHasSumInIntsTrueFirstAndLast(t *testing.T) {
	coll := []int{4, 6, 3, 5, 8, 9, 4, 23}
	target := 26
	want := true

	result := hasSumInInts(coll, target)

	if result != want {
		t.Fatalf("Did not find sum to %v in %v. Result %v, want %v", target, coll, result, want)
	}
}

func TestHasSumInIntsFalse(t *testing.T) {
	coll := []int{4, 6, 3, 5, 8, 9, 4, 23}
	target := 21
	want := false

	result := hasSumInInts(coll, target)

	if result != want {
		t.Fatalf("Found sum to %v in %v. Result %v, want %v", target, coll, result, want)
	}
}

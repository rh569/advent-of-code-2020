package main

import (
	"reflect"
	"testing"
)

var smallTestInput = `16
10
15
5
1
11
7
19
6
12
4`

var mediumTestInput = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestParseInputSmall(t *testing.T) {
	want := []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}
	result := parseInput(smallTestInput)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Parsed input %v does not match wanted %v\n", result, want)
	}
}

func TestParseInputMedium(t *testing.T) {
	want := []int{1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49}
	result := parseInput(mediumTestInput)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Parsed input %v does not match wanted %v\n", result, want)
	}
}

func TestCalculateFullChainDifferences2Elems(t *testing.T) {
	adapters := []int{2, 4}
	want1, want2, want3 := 0, 2, 1

	count1, count2, count3, err := calculateFullChainDifferences(adapters)

	if err != nil {
		t.Fatalf("Error: %v. Wanted %v, %v, %v\n", err, want1, want2, want3)
	}

	if count1 != want1 || count2 != want2 || count3 != want3 {
		t.Fatalf("Counts: 1) %v, 2) %v, 3) %v Wanted: 1) %v, 2) %v, 3) %v\n", count1, count2, count3, want1, want2, want3)
	}
}

func TestCalculateFullChainDifferencesSmall(t *testing.T) {
	adapters := parseInput(smallTestInput)
	want1, want2, want3 := 7, 0, 5

	count1, count2, count3, err := calculateFullChainDifferences(adapters)

	if err != nil {
		t.Fatalf("Error: %v. Wanted %v, %v, %v\n", err, want1, want2, want3)
	}

	if count1 != want1 || count2 != want2 || count3 != want3 {
		t.Fatalf("Counts: 1) %v, 2) %v, 3) %v Wanted: 1) %v, 2) %v, 3) %v\n", count1, count2, count3, want1, want2, want3)
	}
}

func TestCalculateFullChainDifferencesMedium(t *testing.T) {
	adapters := parseInput(mediumTestInput)
	want1, want2, want3 := 22, 0, 10

	count1, count2, count3, err := calculateFullChainDifferences(adapters)

	if err != nil {
		t.Fatalf("Error: %v. Wanted %v, %v, %v\n", err, want1, want2, want3)
	}

	if count1 != want1 || count2 != want2 || count3 != want3 {
		t.Fatalf("Counts: 1) %v, 2) %v, 3) %v Wanted: 1) %v, 2) %v, 3) %v\n", count1, count2, count3, want1, want2, want3)
	}
}

// ---

func TestTestCalculateNumberAdapterPermutationsSmall(t *testing.T) {
	adapters := parseInput(smallTestInput)
	want := int64(8)

	result := calculateNumberAdapterPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestTestCalculateNumberAdapterPermutationsMedium(t *testing.T) {
	adapters := parseInput(mediumTestInput)
	want := int64(19208)

	result := calculateNumberAdapterPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

// ---

func TestGetSubAdapterListsSmall(t *testing.T) {
	adapters := parseInput(smallTestInput)

	want := [][]int{
		[]int{1},
		[]int{4, 5, 6, 7},
		[]int{10, 11, 12},
		[]int{15, 16},
		[]int{19},
	}

	result := getSubAdapterLists(adapters)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

// --- Permutations

func TestCalculateSubListPermutations3Elems(t *testing.T) {
	adapters := []int{4, 6, 7}

	var want int64 = 2

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestCalculateSubListPermutationsSimple3ElemsNo3Diff(t *testing.T) {
	adapters := []int{4, 6, 8}

	var want int64 = 1

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestCalculateSubListPermutations4Elems(t *testing.T) {
	adapters := []int{4, 5, 6, 7}

	var want int64 = 4

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestCalculateSubListPermutationsSimple4ElemsNo3Diff(t *testing.T) {
	adapters := []int{4, 6, 7, 8}

	var want int64 = 3

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestCalculateSubListPermutations5ElemsSingles(t *testing.T) {
	adapters := []int{4, 5, 6, 7, 8}

	var want int64 = 7

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

func TestCalculateSubListPermutations5ElemsMixed(t *testing.T) {
	adapters := []int{4, 5, 7, 9, 10}

	var want int64 = 4

	result := calculateSubListPermutations(adapters)

	if result != want {
		t.Fatalf("Result: %v Want: %v\n", result, want)
	}
}

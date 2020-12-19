package main

import "testing"

var testInput = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

func TestPart1(t *testing.T) {
	want := 2
	result := part1(testInput)

	if result != want {
		t.Fatalf("Result %v, was not want %v\n", result, want)
	}
}

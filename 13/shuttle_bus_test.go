package main

import "testing"

var testInput = `939
7,13,x,x,59,x,31,19`

func TestGetNextBus(t *testing.T) {
	time, buses := parseInput(testInput)

	wantBus := 59
	wantDeparts := 944

	nextBus, departs := getNextBus(time, buses)

	if nextBus != wantBus || departs != wantDeparts {
		t.Fatalf("Bus %v at %v. Want bus %v at %v\n", nextBus, departs, wantBus, wantDeparts)
	}
}

// func TestPart1(t *testing.T) {
// 	want := 295

// 	result := part1(testInput)

// 	if result != want {
// 		t.Fatalf("Result %v. Want %v\n", result, want)
// 	}
// }

// func TestPart2(t *testing.T) {
// 	input := `0
// 7,13,x,x,59,x,31,19`

// 	var want int64 = 1068781

// 	result := part2(input)

// 	if result != want {
// 		t.Fatalf("Result %v. Want %v\n", result, want)
// 	}
// }

func TestPart2Two(t *testing.T) {
	input := `0
17,x,13,19`

	var want int64 = 3417

	result := part2(input)

	if result != want {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

func TestPart2Three(t *testing.T) {
	input := `0
67,7,59,61`

	var want int64 = 754018

	result := part2(input)

	if result != want {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

// This originally took 12 secs to run - real solution will take >12 days ...
func TestPart2Four(t *testing.T) {
	input := `0
1789,37,47,1889`

	var want int64 = 1202161486

	result := part2(input)

	if result != want {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

func TestPart2Five(t *testing.T) {
	input := `0
8,x,5,3`

	var want int64 = 48

	result := part2(input)

	if result != want {
		t.Fatalf("Result %v. Want %v\n", result, want)
	}
}

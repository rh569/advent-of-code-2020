package main

import (
	"testing"
)

const (
	TestPub1 = 5764801
	TestPub2 = 17807724
)

func TestPart1(t *testing.T) {
	want := 14897079
	result := part1(TestPub1, TestPub2)

	if result != want {
		t.Fatalf("Result: %v. Want: %v\n", result, want)
	}
}

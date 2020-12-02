package main

import (
	// "fmt"
	"reflect"
	"testing"
)

var testInput = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc",
}

func TestSplitInput(t *testing.T) {
	want := [][]string{
		{"1", "3", "a", "abcde"},
		{"1", "3", "b", "cdefg"},
		{"2", "9", "c", "ccccccccc"},
	}

	testSplitInput := splitInput(testInput)

	if !reflect.DeepEqual(testSplitInput, want) {
		t.Fatalf("Split input was %v, wanted %v", testSplitInput, want)
	}
}

func TestIsValid1(t *testing.T) {
	want := true
	input := []string{"1", "3", "a", "abcde"}
	valid := isValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

func TestIsValid2(t *testing.T) {
	want := false
	input := []string{"1", "3", "b", "cdefg"}
	valid := isValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

func TestIsValid3(t *testing.T) {
	want := true
	input := []string{"2", "9", "c", "ccccccccc"}
	valid := isValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

func TestIsActuallyValid1(t *testing.T) {
	want := true
	input := []string{"1", "3", "a", "abcde"}
	valid := isActuallyValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

func TestIsActuallyValid2(t *testing.T) {
	want := false
	input := []string{"1", "3", "b", "cdefg"}
	valid := isActuallyValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

func TestIsActuallyValid3(t *testing.T) {
	want := false
	input := []string{"2", "9", "c", "ccccccccc"}
	valid := isActuallyValid(input)

	if valid != want {
		t.Fatalf("input valid was %v, wanted %v", valid, want)
	}
}

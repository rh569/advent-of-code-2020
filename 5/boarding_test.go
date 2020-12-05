package main

import (
	"reflect"
	"testing"
)

var testInput = []string{
	"FFFFFFFLLL",
	"BBBBBBBLLL",
	"BFFFBBFRRR",
	"FFFBBBFRRR",
	"BBFFBBFRLL",
}

func TestCalculatePassId(t *testing.T) {
	want := 820
	id := calculatePassId(102, 4)

	if id != want {
		t.Fatalf("Got %v, want %v\n", id, want)
	}
}

func TestFindPassRow(t *testing.T) {
	want := []int{0, 127, 70, 14, 102}
	rows := []int{}

	for _, pass := range testInput {
		rows = append(rows, findPassRow(pass[:len(pass)-3]))
	}

	if !reflect.DeepEqual(rows, want) {
		t.Fatalf("Got %v, want %v\n", rows, want)
	}
}

func TestFindPassCol(t *testing.T) {
	want := []int{0, 0, 7, 7, 4}
	cols := []int{}

	for _, pass := range testInput {
		cols = append(cols, findPassCol(pass[len(pass)-3:]))
	}

	if !reflect.DeepEqual(cols, want) {
		t.Fatalf("Got %v, want %v\n", cols, want)
	}
}

func TestFindMissingIdOneMissing(t *testing.T) {
	testIds := []int{13, 12, 17, 14, 16} // no 15
	want := 15

	id, err := findMissingId(testIds)

	if err != nil {
		t.Fatalf("Should not err: %v", err)
	}

	if id != want {
		t.Fatalf("Found %v, want %v", id, want)
	}
}

func TestFindMissingIdMoreMissing(t *testing.T) {
	testIds := []int{12, 17, 14, 16} // no 13 or 15

	id, err := findMissingId(testIds)

	if err == nil {
		t.Fatalf("Should err. Got %v", id)
	}
}

func TestFindMissingIdBiggerGap(t *testing.T) {
	testIds := []int{12, 17, 16} // no 13 or 15

	id, err := findMissingId(testIds)

	if err == nil {
		t.Fatalf("Should err. Got %v", id)
	}
}
package main

import (
	"reflect"
	"testing"
)

var testInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

func TestParseInput(t *testing.T) {
	wantFields := []Field{
		Field{"class", 1, 3, 5, 7},
		Field{"row", 6, 11, 33, 44},
		Field{"seat", 13, 40, 45, 50},
	}
	wantYours := []int{7, 1, 14}
	wantNearby := [][]int{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}

	fields, yours, nearby := parseInput(testInput)

	if !reflect.DeepEqual(fields, wantFields) {
		t.Fatalf("Fields: Result: %v, want: %v\n", fields, wantFields)
	}

	if !reflect.DeepEqual(yours, wantYours) {
		t.Fatalf("Your ticket: Result: %v, want: %v\n", yours, wantYours)
	}

	if !reflect.DeepEqual(nearby, wantNearby) {
		t.Fatalf("Near tickets: Result: %v, want: %v\n", nearby, wantNearby)
	}
}

func TestPart1(t *testing.T) {
	want := 71
	result := part1(testInput)

	if result != want {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

var validTestInput = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestGetFieldPositions(t *testing.T) {
	fields, _, tickets := parseInput(validTestInput)

	want := map[string]int{
		"class": 1,
		"row":   0,
		"seat":  2,
	}

	result := getFieldPositions(fields, getValidTickets(fields, tickets))

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want: %v\n", result, want)
	}
}

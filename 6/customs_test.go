package main

import (
	"reflect"
	"testing"
)

var testAnswers = `abc

a
b
c

ab
ac

a
a
a
a

b

w
s
q
s`

func TestParseAnswers(t *testing.T) {
	want := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
		{"w", "s", "q", "s"},
	}

	parsedAnswers := parseAnswers(testAnswers)

	if !reflect.DeepEqual(parsedAnswers, want) {
		t.Fatalf("Parsed answers:\n%v\nDid not match wanted:\n%v\n", parsedAnswers, want)
	}
}

func TestCollateAllAnswers(t *testing.T) {
	parsedAnswers := parseAnswers(testAnswers)

	want := [][]string{
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a", "b", "c"},
		{"a"},
		{"b"},
		{"q", "s", "w"},
	}

	collated := collateAllAnswers(parsedAnswers)

	if !reflect.DeepEqual(collated, want) {
		t.Fatalf("Collated answers:\n%v\nDid not match wanted:\n%v\n", collated, want)
	}
}

func TestCollateAllMatchingGroupAnswers(t *testing.T) {
	parsedAnswers := parseAnswers(testAnswers)

	want := [][]string{
		{"a", "b", "c"},
		{},
		{"a"},
		{"a"},
		{"b"},
		{},
	}

	collated := collateAllMatchingAnswers(parsedAnswers)

	if !reflect.DeepEqual(collated, want) {
		t.Fatalf("Collated matching answers:\n%v\nDid not match wanted:\n%v\n", collated, want)
	}
}

func TestCountAllCollated(t *testing.T) {
	collated := collateAllAnswers(parseAnswers(testAnswers))

	want := 14

	count := countAllCollated(collated)

	if count != want {
		t.Fatalf("Counted collated answers:\n%v\nDid not match wanted:\n%v\n", count, want)
	}
}

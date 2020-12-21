package main

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	// result := part1(TestInput)
	result, _ := part1(TestInput)

	want := 5

	if want != result {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

func TestRemoveIngredientFirst(t *testing.T) {
	want := []string{"two", "three", "four"}

	result := removeIngredient([]string{"one", "two", "three", "four"}, "one")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

func TestRemoveIngredientLast(t *testing.T) {
	want := []string{"one", "two", "three"}

	result := removeIngredient([]string{"one", "two", "three", "four"}, "four")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

func TestRemoveIngredientMid(t *testing.T) {
	want := []string{"one", "two", "four"}

	result := removeIngredient([]string{"one", "two", "three", "four"}, "three")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

func TestRemoveIngredientOnly(t *testing.T) {
	want := []string{}

	result := removeIngredient([]string{"one"}, "one")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

func TestRemoveIngredientMissing(t *testing.T) {
	want := []string{"one"}

	result := removeIngredient([]string{"one"}, "two")

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result: %v, want : %v\n", result, want)
	}
}

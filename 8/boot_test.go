package main

import (
	"reflect"
	"testing"
)

var testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestGetSignedIntPos(t *testing.T) {
	want := 42
	num := getSignedInt("+42")

	if want != num {
		t.Fatalf("Num was %v. Want %v", num, want)
	}
}

func TestGetSignedIntNeg(t *testing.T) {
	want := -420
	num := getSignedInt("-420")

	if want != num {
		t.Fatalf("Num was %v. Want %v", num, want)
	}
}

func TestParseInstructionNop(t *testing.T) {
	want := Ins{"nop", 102}
	ins := parseInstruction("nop +102")

	if want != ins {
		t.Fatalf("Ins was %v. Want %v", ins, want)
	}
}

func TestParseInstructionAcc(t *testing.T) {
	want := Ins{"acc", 2}
	ins := parseInstruction("acc +2")

	if want != ins {
		t.Fatalf("Ins was %v. Want %v", ins, want)
	}
}

func TestParseInstructionJmp(t *testing.T) {
	want := Ins{"jmp", -11}
	ins := parseInstruction("jmp -11")

	if want != ins {
		t.Fatalf("Ins was %v. Want %v", ins, want)
	}
}

func TestParseInput(t *testing.T) {
	want := []Ins{
		Ins{"nop", 0},
		Ins{"acc", 1},
		Ins{"jmp", 4},
		Ins{"acc", 3},
		Ins{"jmp", -3},
		Ins{"acc", -99},
		Ins{"acc", 1},
		Ins{"jmp", -4},
		Ins{"acc", 6},
	}

	instructions := parseInput(testInput)

	if !reflect.DeepEqual(instructions, want) {
		t.Fatalf("Instructions were %v. Want %v", instructions, want)
	}
}

func TestContainsIntTrue(t *testing.T) {
	testCollection := []int{12, 144, 765, -9, 0, 22, 7}
	testTarget := -9

	want := true
	result := containsInt(testCollection, testTarget)

	if result != want {
		t.Fatalf("Result: %v. Want: %v. Should have found %v in %v", result, want, testTarget, testCollection)
	}
}

func TestContainsIntFalseStart(t *testing.T) {
	testCollection := []int{12, 144, 765, -9, 0, 22, 7}
	testTarget := -122

	want := false
	result := containsInt(testCollection, testTarget)

	if result != want {
		t.Fatalf("Result: %v. Want: %v. Should not have found %v in %v", result, want, testTarget, testCollection)
	}
}

func TestContainsIntFalseMid(t *testing.T) {
	testCollection := []int{12, 144, 765, -9, 0, 22, 7}
	testTarget := 122

	want := false
	result := containsInt(testCollection, testTarget)

	if result != want {
		t.Fatalf("Result: %v. Want: %v. Should not have found %v in %v", result, want, testTarget, testCollection)
	}
}

func TestContainsIntEnd(t *testing.T) {
	testCollection := []int{12, 144, 765, -9, 0, 22, 7}
	testTarget := 200

	want := false
	result := containsInt(testCollection, testTarget)

	if result != want {
		t.Fatalf("Result: %v. Want: %v. Should not have found %v in %v", result, want, testTarget, testCollection)
	}
}

func TestContainsIntEmpty(t *testing.T) {
	testCollection := []int{}
	testTarget := 122

	want := false
	result := containsInt(testCollection, testTarget)

	if result != want {
		t.Fatalf("Result: %v. Want: %v. Should not have found %v in %v", result, want, testTarget, testCollection)
	}
}

func TestRunInstructionsLoop(t *testing.T) {
	instructions := parseInput(testInput)

	wantAcc, wantFinished := 5, false

	acc, finished := runInstructions(instructions)

	if acc != wantAcc {
		t.Fatalf("Accumulator was %v, want %v", acc, wantAcc)
	}

	if finished != wantFinished {
		t.Fatalf("Program finished: %v, want: %v", finished, wantFinished)
	}
}

func TestMutateInstructionsUntilDone(t *testing.T) {
	instructions := parseInput(testInput)

	want := 8

	acc, err := mutateInstructionsUntilDone(instructions)

	if err != nil {
		t.Fatalf("Error: '%v', want %v", err, acc)
	}

	if acc != want {
		t.Fatalf("Program finished with acc: %v, want: %v", acc, want)
	}
}

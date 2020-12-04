package main

import (
	"testing"
)

func TestValidateYearValid(t *testing.T) {
	want := 2
	valid := 0

	if validateYear(2002, 2002, 2002) {
		valid++
	}
	if validateYear(2002, 2001, 2003) {
		valid++
	}

	if valid != want {
		t.Fatalf("Found %v valid, wanted %v\n", valid, want)
	}
}

func TestValidateYearInvalid(t *testing.T) {
	want := 2
	invalid := 0

	if !validateYear(2002, 2005, 2002) {
		invalid++
	}
	if !validateYear(2002, 1990, 1994) {
		invalid++
	}

	if invalid != want {
		t.Fatalf("Found %v invalid, wanted %v\n", invalid, want)
	}
}

func TestValidateHeightValid(t *testing.T) {
	want := 4
	valid := 0

	if validateHeight("59in") {
		valid++
	}
	if validateHeight("76in") {
		valid++
	}
	if validateHeight("193cm") {
		valid++
	}
	if validateHeight("150cm") {
		valid++
	}

	if valid != want {
		t.Fatalf("Found %v valid, wanted %v\n", valid, want)
	}
}

func TestValidateHeightInvalid(t *testing.T) {
	want := 5
	invalid := 0

	if !validateHeight("58in") {
		invalid++
	}
	if !validateHeight("77in") {
		invalid++
	}
	if !validateHeight("149cm") {
		invalid++
	}
	if !validateHeight("194cm") {
		invalid++
	}
	if !validateHeight("190") {
		invalid++
	}

	if invalid != want {
		t.Fatalf("Found %v invalid, wanted %v\n", invalid, want)
	}
}

func TestValidateHairColorValid(t *testing.T) {
	want := 1
	valid := 0

	if validateHairColor("#aaaa00") {
		valid++
	}

	if valid != want {
		t.Fatalf("Found %v valid, wanted %v\n", valid, want)
	}
}

func TestValidateHairColorInvalid(t *testing.T) {
	want := 3
	invalid := 0

	if !validateHairColor("aaga00") {
		invalid++
	}
	if !validateHairColor("aaaa00") {
		invalid++
	}
	if !validateHairColor("#aaa00") {
		invalid++
	}

	if invalid != want {
		t.Fatalf("Found %v invalid, wanted %v\n", invalid, want)
	}
}

func TestValidateEyeColorValid(t *testing.T) {
	want := 7
	valid := 0

	if validateEyeColor("amb") {
		valid++
	}
	if validateEyeColor("blu") {
		valid++
	}
	if validateEyeColor("brn") {
		valid++
	}
	if validateEyeColor("gry") {
		valid++
	}
	if validateEyeColor("grn") {
		valid++
	}
	if validateEyeColor("hzl") {
		valid++
	}
	if validateEyeColor("oth") {
		valid++
	}

	if valid != want {
		t.Fatalf("Found %v valid, wanted %v\n", valid, want)
	}
}

func TestValidateEyeColorInvalid(t *testing.T) {
	want := 1
	invalid := 0

	if !validateEyeColor("eye") {
		invalid++
	}

	if invalid != want {
		t.Fatalf("Found %v invalid, wanted %v\n", invalid, want)
	}
}

func TestValidatePidValid(t *testing.T) {
	want := 1
	valid := 0

	if validatePid("000000001") {
		valid++
	}

	if valid != want {
		t.Fatalf("Found %v valid, wanted %v\n", valid, want)
	}
}

func TestValidatePidInvalid(t *testing.T) {
	want := 2
	invalid := 0

	if !validatePid("0123456789") {
		invalid++
	}
	if !validatePid("01234567") {
		invalid++
	}

	if invalid != want {
		t.Fatalf("Found %v invalid, wanted %v\n", invalid, want)
	}
}

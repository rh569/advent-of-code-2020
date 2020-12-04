package main

import (
	// "fmt"
	"regexp"
)

func validateYear(year int, start int, end int) bool {
	return year >= start && year <= end
}

func validateHeight(height string) bool {

	if len(height) < 4 {
		return false
	}

	unit, numStr := height[len(height)-2:], height[:len(height)-2]
	heightValue := getInt(numStr)

	if unit == "cm" {
		return heightValue >= 150 && heightValue <= 193
	} else if unit == "in" {
		return heightValue >= 59 && heightValue <= 76
	} else {
		return false
	}
}

func validateHairColor(color string) bool {

	if len(color) != 7 {
		return false
	}

	r := regexp.MustCompile("^#[0-9a-f]{6}$")

	return r.MatchString(color)
}

func validateEyeColor(color string) bool {

	r := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")

	return r.MatchString(color)
}

func validatePid(pid string) bool {

	r := regexp.MustCompile("^[0-9]{9}$")

	return r.MatchString(pid)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Passport struct {
	byr, iyr, eyr int
	hgt, hcl, ecl string
	pid           string // can start 0
	cid           int
}

func (p *Passport) isValid() bool {
	return p.byr > 0 &&
		p.iyr > 0 &&
		p.eyr > 0 &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p *Passport) isStrictlyValid() bool {
	byrValid := validateYear(p.byr, 1920, 2002)
	iyrValid := validateYear(p.iyr, 2010, 2020)
	eyrValid := validateYear(p.eyr, 2020, 2030)

	hgtValid := validateHeight(p.hgt)
	hclValid := validateHairColor(p.hcl)
	eclValid := validateEyeColor(p.ecl)

	pidValid := validatePid(p.pid)

	return byrValid &&
		iyrValid &&
		eyrValid &&
		hgtValid &&
		hclValid &&
		eclValid &&
		pidValid
}

func parseBatch(passportBatch string) []Passport {
	passports := []Passport{}

	// Passports separated by newline
	passportStrings := strings.Split(passportBatch, "\n\n")

	for _, passportString := range passportStrings {
		passports = append(passports, parsePassportString(passportString))
	}

	// fmt.Printf("Found %v passport strings\n", len(passportStrings))

	return passports
}

func parsePassportString(passportString string) Passport {
	var passport Passport

	passportString = strings.ReplaceAll(passportString, "\n", " ")
	parts := strings.Split(passportString, " ")

	for _, part := range parts {
		keyValue := strings.Split(part, ":")

		switch keyValue[0] {
		case "byr":
			passport.byr = getInt(keyValue[1])
		case "iyr":
			passport.iyr = getInt(keyValue[1])
		case "eyr":
			passport.eyr = getInt(keyValue[1])
		case "hgt":
			passport.hgt = keyValue[1]
		case "hcl":
			passport.hcl = keyValue[1]
		case "ecl":
			passport.ecl = keyValue[1]
		case "pid":
			passport.pid = keyValue[1]
		case "cid":
			passport.cid = getInt(keyValue[1])
		default:
			fmt.Printf("Unexpected passport key: %v\n", keyValue[0])
		}
	}

	return passport
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

func countValid(batch string, strict bool) int {
	passports := parseBatch(batch)
	valid := 0

	if strict {
		for _, passport := range passports {
			if passport.isStrictlyValid() {
				valid++
			}
		}
	} else {
		for _, passport := range passports {
			if passport.isValid() {
				valid++
			}
		}
	}

	return valid
}

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Number of valid passports: %v\n", countValid(PassportBatch, false))

	fmt.Println("Part 2:")
	fmt.Printf("Number of strictly valid passports: %v\n", countValid(PassportBatch, true))
}

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	parsedAnswers := parseAnswers(Answers)

	orCount := countAllCollated(collateAllAnswers(parsedAnswers))

	fmt.Println("Part 1:")
	fmt.Printf("Count of all answers was: %v\n", orCount)

	andCount := countAllCollated(collateAllMatchingAnswers(parsedAnswers))

	fmt.Println("Part 2:")
	fmt.Printf("Count of all answers was: %v\n", andCount)
}

// Parses multiline input to array of groups' answers
// Individuals' answers remain as strings
func parseAnswers(answers string) [][]string {
	parsed := [][]string{}

	groups := strings.Split(answers, "\n\n")

	for _, group := range groups {
		parsed = append(parsed, strings.Split(group, "\n"))
	}

	return parsed
}

// returns an array of all unique answers in a group
func collateGroupAnswers(groupAnswers []string) []string {
	collated := []string{}

	for _, indAnswers := range groupAnswers {

		for _, char := range indAnswers {
			letter := string(char)

			j := sort.SearchStrings(collated, letter)

			if j < len(collated) {

				if collated[j] != letter {
					collated = append(collated, letter)
				}
			} else {
				collated = append(collated, letter)
			}

			sort.Strings(collated)
		}
	}

	return collated
}

// Collate group answers for every group in the input
func collateAllAnswers(parsedAnswers [][]string) [][]string {
	collated := [][]string{}

	for _, groupAnswers := range parsedAnswers {
		collated = append(collated, collateGroupAnswers(groupAnswers))
	}

	return collated
}

// returns an array of all matching (agreed upon) answers in a group
func collateMatchingGroupAnswers(groupAnswers []string) []string {
	collated := []string{}

	for i, indAnswers := range groupAnswers {

		answersArr := strings.Split(indAnswers, "")

		if i == 0 {
			// add all answers from the first person
			collated = append(collated, answersArr...)
			sort.Strings(collated)
			continue
		}

		// new slice for matching answers
		agreed := []string{}

		for _, letter := range answersArr {
			j := sort.SearchStrings(collated, letter)

			if j < len(collated) {

				// if their letter is already in the collated slice, they agree
				if collated[j] == letter {
					agreed = append(agreed, letter)
				}
			}
		}

		collated = append([]string{}, agreed...)
		sort.Strings(collated)
	}

	return collated
}

// Collate matching group answers for every group in the input
func collateAllMatchingAnswers(parsedAnswers [][]string) [][]string {
	collated := [][]string{}

	for _, groupAnswers := range parsedAnswers {
		collated = append(collated, collateMatchingGroupAnswers(groupAnswers))
	}

	return collated
}

func countAllCollated(collatedAnswers [][]string) int {
	var count int

	for _, group := range collatedAnswers {
		count += len(group)
	}

	return count
}

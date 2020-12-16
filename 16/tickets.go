package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Invalid sum: %v\n", part1(TicketInput))

	fmt.Println("Part 2:")
	fmt.Printf("Departure product: %v\n", part2(TicketInput))
}

type Field struct {
	name                   string
	min1, max1, min2, max2 int
}

func part1(input string) int {
	fields, _, tickets := parseInput(input)

	invalidFields := []int{}

	for _, ticket := range tickets {
		invalidFields = append(invalidFields, getInvalidTicketFields(fields, ticket)...)
	}

	sum := 0

	for _, f := range invalidFields {
		sum += f
	}

	return sum
}

func part2(input string) int {
	fields, yours, tickets := parseInput(input)

	validTickets := getValidTickets(fields, tickets)

	// map of field names, e.g. class, to their ticket index
	fieldNamePositions := getFieldPositions(fields, validTickets)

	product := 1

	for k, v := range fieldNamePositions {

		if strings.HasPrefix(k, "departure") {
			product *= yours[v]
		}
	}

	return product
}

// Only one permutation is possible (at least, it seems, in the given input - therefore we're not exiting lazily)
// After getting every possible position for a rule, we narrow down by excluding one number at a time
func getFieldPositions(fields []Field, tickets [][]int) map[string]int {
	possibleFieldPositions := make(map[string][]int)

	remainingFields := []Field{}
	remainingFields = append(remainingFields, fields...)

	// While there are fewer determined positions than initial field rules
	for len(remainingFields) > 0 {

		field := remainingFields[0]

		// Find possible field positions
		for i := 0; i < len(fields); i++ {

			allTicketsMatchPos := true

			// where every ticket matches
			for _, ticket := range tickets {
				if !matchesRule(field, ticket[i]) {
					allTicketsMatchPos = false
					break
				}
			}

			if allTicketsMatchPos {
				possiblePositions := possibleFieldPositions[field.name]
				possiblePositions = append(possiblePositions, i)
				possibleFieldPositions[field.name] = possiblePositions

				remainingFields = removeField(remainingFields, field)
			}
		}
	}

	return narrowPossiblePositions(possibleFieldPositions)
}

// For every field rule, find a rule that has exactly one possible position
// Remove that position from every other field
// Repeat until all fields narrowed
// Assumes there's only ever one field that is the correct field to find in a given iteration
func narrowPossiblePositions(possibles map[string][]int) map[string]int {
	actualPositions := make(map[string]int)

	for len(possibles) > 0 {

		for k, v := range possibles {
			if len(v) == 1 {
				actualPositions[k] = v[0]
				exclude(possibles, v[0])
				delete(possibles, k)
			}
		}
	}

	return actualPositions
}

// for every array in the map, removes all values matching ex
func exclude(m map[string][]int, ex int) map[string][]int {

	for k, v := range m {
		newV := []int{}

		for _, num := range v {
			if num != ex {
				newV = append(newV, num)
			}
		}

		m[k] = newV
	}

	return m
}

// Assumes unique field names
func removeField(fields []Field, ditch Field) []Field {
	newFields := []Field{}

	for _, f := range fields {
		if f.name != ditch.name {
			newFields = append(newFields, f)
		}
	}

	return newFields
}

// returns ture if val occurs in in []int
func contains(in []int, val int) bool {

	for _, i := range in {
		if i == val {
			return true
		}
	}

	return false
}

func getValidTickets(fields []Field, tickets [][]int) [][]int {
	invalidFields := []int{}
	validTickets := [][]int{}

	for _, ticket := range tickets {
		invalidFields = getInvalidTicketFields(fields, ticket)

		if len(invalidFields) == 0 {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

// returns all values on a ticket that don't match ANY field rules
func getInvalidTicketFields(fieldRules []Field, ticket []int) []int {
	invalid := []int{}

	for _, val := range ticket {
		matchedAny := false

		for _, rule := range fieldRules {
			if matchesRule(rule, val) {
				matchedAny = true
			}
		}

		if !matchedAny {
			invalid = append(invalid, val)
		}
	}

	return invalid
}

func matchesRule(rule Field, val int) bool {
	return (val >= rule.min1 && val <= rule.max1) || (val >= rule.min2 && val <= rule.max2)
}

// Returns the string key for the given value
// Values are assumed to be unique
// If val does not exist, returns ""
func getKey(m map[string]int, val int) string {
	for k, v := range m {
		if v == val {
			return k
		}
	}

	return ""
}

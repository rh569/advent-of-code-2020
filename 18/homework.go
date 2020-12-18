package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("The sum of sums is: %v\n", part1(MathHomeworkInput))

	fmt.Println("Part 2:")
	fmt.Printf("The sum of sums is: %v\n", part2(MathHomeworkInput))
}

func part1(input string) int {
	sums := parseInput(input)
	sumOfSums := 0

	for _, sum := range sums {
		sumOfSums += evaluateSum(sum, false)
	}

	return sumOfSums
}

func part2(input string) int {
	sums := parseInput(input)
	sumOfSums := 0

	for _, sum := range sums {
		sumOfSums += evaluateSum(sum, true)
	}

	return sumOfSums
}

func evaluateSum(sum Sum, opPrecedence bool) int {
	// while there are multiple nested expressions
	for len(sum.s) > 1 {

		// find the nested expressions matching the highest n
		for i, nestedExp := range sum.s {

			// for those with the highest n, evaluate
			if nestedExp.n == sum.n {

				if opPrecedence {
					sum.s[i].e = evaluateExpressionWithOpPrecedence(nestedExp.e)
				} else {
					sum.s[i].e = evaluateExpression(nestedExp.e)
				}
				sum.s[i].n--
			}
		}

		sum.s = collapseTerms(sum.s)

		sum.n--
	}

	var lastSum []string

	if opPrecedence {
		lastSum = evaluateExpressionWithOpPrecedence(sum.s[0].e)
	} else {
		lastSum = evaluateExpression(sum.s[0].e)
	}

	return getInt(lastSum[0])
}

func evaluateExpression(e []string) []string {
	var total int
	var op string

	for _, char := range e {

		if char == "+" || char == "*" {
			op = char
			continue
		}

		num := getInt(char)

		switch op {
		case "+":
			total += num
		case "*":
			total *= num
		default:
			total = num
		}
	}

	return []string{fmt.Sprint(total)}
}

func evaluateExpressionWithOpPrecedence(e []string) []string {
	var total int

	containsPlus := checkForPlus(e)

	// Deal with all additions
	for containsPlus {
		newExpression := []string{}

		for i, char := range e {

			if char == "+" {

				// copy prior chars
				if i > 2 {
					newExpression = append(newExpression, e[:i-1]...)
				}

				// evaluate addition and append result
				newExpression = append(newExpression, fmt.Sprint(getInt(e[i-1])+getInt(e[i+1])))

				// copy remaining chars
				if i < len(e)-3 {
					newExpression = append(newExpression, e[i+2:]...)
				}

				break
			}
		}

		e = newExpression
		containsPlus = checkForPlus(e)
	}

	if len(e) == 1 {
		return e
	}

	total = getInt(e[0])

	for i := 2; i < len(e); i++ {
		if i%2 == 0 {
			total *= getInt(e[i])
		}
	}

	return []string{fmt.Sprint(total)}
}

func checkForPlus(in []string) bool {
	containsPlus := false

	for _, char := range in {

		if char == "+" {
			containsPlus = true
			break
		}
	}

	return containsPlus
}

// Collapsed adjacent terms of the same nesting, returning the collapsed array
func collapseTerms(expressions []NestedExpression) []NestedExpression {
	collpasedExpressions := []NestedExpression{}

	// fist expression will always come first, may be appended to
	newExpression := expressions[0]

	// for each expression, but the first
	for i := 1; i < len(expressions); i++ {
		currentExpression := expressions[i]

		// current n matches the expression we're building
		if currentExpression.n == newExpression.n {
			newExpression.e = append(newExpression.e, currentExpression.e...)
		} else {
			collpasedExpressions = append(collpasedExpressions, newExpression)
			newExpression = currentExpression
		}
	}

	// If the last element was included in the collapse
	if newExpression.e != nil {
		collpasedExpressions = append(collpasedExpressions, newExpression)
	}

	return collpasedExpressions
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

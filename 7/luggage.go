package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BagRule struct {
	name     string
	contains []InnerBag
}

type InnerBag struct {
	name   string
	number int
}

func main() {
	rules := parseRules(Rules)

	parents := calculateNumberOfParents(rules, "shiny gold")

	fmt.Println("Part 1:")
	fmt.Printf("Found %v possible parent bags\n", parents)

	children := calculateNumberOfChildren(rules, "shiny gold")

	fmt.Println("Part 2:")
	fmt.Printf("Must contain %v child bags\n", children)
}

// Parse the full list of rules
func parseRules(rules string) []BagRule {
	lines := strings.Split(rules, "\n")

	bagRules := []BagRule{}

	for _, rule := range lines {
		bagRules = append(bagRules, parseRule(rule))
	}

	return bagRules
}

// Parse an individual rule
func parseRule(ruleString string) BagRule {
	var rule BagRule

	parts := strings.Split(ruleString, " bags contain ")
	rule.name = parts[0]
	rule.contains = []InnerBag{}

	innerBagsString := parts[1]

	if innerBagsString != "no other bags." {
		innerBagsString = strings.ReplaceAll(innerBagsString, ".", ", ")
		innerBagsString = strings.ReplaceAll(innerBagsString, "bags", "bag")

		innerBags := strings.Split(innerBagsString, " bag, ")

		for _, innerBagString := range innerBags {
			if len(innerBagString) == 0 {
				continue
			}

			number := getInt(string(innerBagString[0]))
			rule.contains = append(rule.contains, InnerBag{innerBagString[2:], number})
		}
	}

	return rule
}

// given a ruleset and a target bag type, will count the number
// of bags that can contain the target bag either directly or indirectly
func calculateNumberOfParents(rules []BagRule, targetBagName string) int {
	// either the target bag name, or names of bags that ultimately contain the target name
	childrenToConsider := []string{targetBagName}
	var childrenFoundLastLoop int
	childrenFound := len(childrenToConsider)

	var loops int

	// loop until we've found all matching bags
	// This is horrible, but without having a proper data structure implementing the directed graph, this has to do...
	for childrenFound-childrenFoundLastLoop > 0 {

		if loops >= 500 {
			log.Fatalf("Looped %v times! Exiting now.", loops)
		} else {
			loops++
		}

		// for every bag rule, check if bagRule.contains has any of our childrenToConsider
		// if so increment count and add to children to consider
		for _, bagRule := range rules {

			// ignore bags already considered
			if containsString(childrenToConsider, bagRule.name) {
				continue
			}

			// check each inner bag
			for _, innerBag := range bagRule.contains {

				// can contain a bag we're considering and we don't already have this bagRule considered
				if containsString(childrenToConsider, innerBag.name) && !containsString(childrenToConsider, bagRule.name) {
					// this bag can contain the target
					childrenToConsider = append(childrenToConsider, bagRule.name)
				}
			}
		}

		childrenFoundLastLoop = childrenFound
		childrenFound = len(childrenToConsider)
	}

	return len(childrenToConsider) - 1
}

// The recursive method returns one too many as it has to include the parent bag as 1 at each level
// including the initial target bag which is not a child
func calculateNumberOfChildren(rules []BagRule, bagName string) int {
	allBags := calculateNumberOfChildrenPlusParent(rules, getRuleByName(rules, bagName))

	return allBags - 1
}

// recurses over a given ruleset counting the number
// of bags that are contained within the target bagRule including itself
func calculateNumberOfChildrenPlusParent(rules []BagRule, bag BagRule) int {
	var count = 1

	if len(bag.contains) == 0 {
		return 1
	}

	for _, innerBag := range bag.contains {
		innerBagRule := getRuleByName(rules, innerBag.name)

		count += (innerBag.number * calculateNumberOfChildrenPlusParent(rules, innerBagRule))
	}

	return count
}

func getRuleByName(rules []BagRule, bagName string) BagRule {

	for _, bag := range rules {
		if bag.name == bagName {
			return bag
		}
	}

	log.Fatalf("Rules do not contain bag of name %v", bagName)
	return BagRule{}
}

// naively checks if target string is in string slice
func containsString(list []string, target string) bool {
	for _, entry := range list {
		if entry == target {
			return true
		}
	}

	return false
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

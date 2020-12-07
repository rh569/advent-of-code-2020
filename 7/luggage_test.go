package main

import (
	"reflect"
	"testing"
)

var testRules = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestParseRules(t *testing.T) {
	parseTestRules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
bright white bags contain no other bags.
muted yellow bags contain no other bags.`

	want := []BagRule{
		BagRule{"light red", []InnerBag{InnerBag{"bright white", 1}, InnerBag{"muted yellow", 2}}},
		BagRule{"bright white", []InnerBag{}},
		BagRule{"muted yellow", []InnerBag{}},
	}

	bags := parseRules(parseTestRules)

	if !reflect.DeepEqual(bags, want) {
		t.Fatalf("Bag rules found: %v\nBag rules wanted: %v\n", bags, want)
	}
}

func TestCalculateNumberOfParents(t *testing.T) {
	rules := parseRules(testRules)
	want := 4

	parents := calculateNumberOfParents(rules, "shiny gold")

	if parents != want {
		t.Fatalf("Found %v parents, want %v", parents, want)
	}
}

func TestCalculateNuberOfChildren(t *testing.T) {
	var testChildrenRules = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

	rules := parseRules(testChildrenRules)
	want := 126

	children := calculateNumberOfChildren(rules, "shiny gold")

	if children != want {
		t.Fatalf("Found %v children, want %v", children, want)
	}
}

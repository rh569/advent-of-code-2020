package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Cups after 100 moves: %v\n", part1(CupsInput, 100))

	fmt.Println("Part 2:")
	fmt.Printf("Star cup product after 10M moves: %v\n", part2(CupsInput))
}

func part1(input string, moves int) string {
	cups := parseInput(input)
	startingLabel := cups[0].label

	sort.Slice(cups, func(i, j int) bool {
		return cups[i].label < cups[j].label
	})

	playGame(cups, moves, startingLabel)
	return toString(cups)
}

func part2(input string) int {
	numCups := 1000000
	givenCups := parseInput(input)
	cups := make([]*Cup, numCups)
	copy(cups, givenCups)

	for i := len(givenCups); i < len(cups); i++ {
		cups[i] = &Cup{
			label: i + 1,
		}
	}

	for i := 0; i < len(cups); i++ {
		prevPos := (i + numCups - 1) % numCups
		nextPos := (i + numCups + 1) % numCups

		cups[i].prev = cups[prevPos]
		cups[i].next = cups[nextPos]
	}

	startingLabel := cups[0].label

	sort.Slice(cups, func(i, j int) bool {
		return cups[i].label < cups[j].label
	})

	playGame(cups, 10000000, startingLabel)

	one, found := findCup(cups[0], 1)

	if found {
		fmt.Println(one.next.label)
		fmt.Println(one.next.next.label)
		return one.next.label * one.next.next.label
	}

	panic("couldn't find cup 1")
}

func playGame(cups []*Cup, moves int, startingLabel int) {
	current, _ := findCupFast(cups, []*Cup{}, startingLabel)

	for i := 0; i < moves; i++ {
		var destination *Cup

		// get next 3
		nextThree := []*Cup{current.next, current.next.next, current.next.next.next}

		// close the circle
		current.next = nextThree[2].next // current points to after the 3
		nextThree[2].next.prev = current // the cup after the three points back to current

		// find destination
		for i := 1; i < 5; i++ {
			searchLabel := (current.label - i + len(cups)) % len(cups)

			// I feel like this should be automatic with the right modulo arithmetic above...
			if searchLabel == 0 {
				searchLabel = len(cups) - searchLabel
			}

			dest, found := findCupFast(cups, nextThree, searchLabel)

			if found {
				destination = dest
				break
			}
		}

		// insert the 3
		nextThree[2].next = destination.next
		nextThree[0].prev = destination
		destination.next = nextThree[0]
		nextThree[2].next.prev = nextThree[2]

		current = current.next
	}
}

func findCupFast(cups []*Cup, nextThree []*Cup, label int) (*Cup, bool) {

	for _, c := range nextThree {
		if c.label == label {
			return nil, false
		}
	}

	index := sort.Search(len(cups), func(i int) bool { return cups[i].label >= label })

	if index < len(cups) && cups[index].label == label {
		return cups[index], true
	} else {
		fmt.Printf("Couldn't find: %v at all.\n", label)
		return nil, false
	}
}

// Too slow, would have taken ~40m to do 1M moves
func findCup(startingCup *Cup, label int) (*Cup, bool) {

	done := false
	startLabel := startingCup.label
	current := startingCup

	for !done {
		if current.next.label == label {
			return current.next, true
		}

		if current.next.label == startLabel {
			done = true
		} else {
			current = current.next
		}
	}

	return nil, false
}

func toString(cups []*Cup) string {
	var current Cup
	str := ""

	// Print from after 1
	for _, cup := range cups {
		if cup.label == 1 {
			current = *cup.next
		}
	}

	done := false

	for !done {
		str += fmt.Sprint(current.label)

		current = *current.next

		done = current.label == 1
	}

	return str
}

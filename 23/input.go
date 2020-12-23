package main

var TestInput = `389125467`

var CupsInput = `712643589`

type Cup struct {
	label int
	prev  *Cup
	next  *Cup
}

func parseInput(input string) []*Cup {
	numCups := len(input)
	cups := make([]*Cup, numCups)

	// Store the cups in an array
	for i, r := range input {
		cups[i] = &Cup{
			label: getIntFromRune(r),
		}
	}

	// Set up pointers
	for i := 0; i < numCups; i++ {
		prevPos := (i + numCups - 1) % numCups
		nextPos := (i + numCups + 1) % numCups

		cups[i].prev = cups[prevPos]
		cups[i].next = cups[nextPos]
	}

	return cups
}

func getIntFromRune(r rune) int {
	return int(r - rune(48))
}

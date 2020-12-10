package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	adapters := parseInput(AdaptersInput)

	count1, count2, count3, err := calculateFullChainDifferences(adapters)

	if err != nil {
		log.Fatalf("Error calculating differences: %v\n", err)
	}

	fmt.Println("Part 1:")
	fmt.Printf("Differences: 1J %v, 2J %v, 3J %v\n", count1, count2, count3)
	fmt.Printf("Product of 1J and 3J = %v\n", count1*count3)

	allPermutations := calculateNumberAdapterPermutations(adapters)

	fmt.Println("Part 2:")
	fmt.Printf("Permutations = %v\n", allPermutations)
}

// Loops though a sorted adapter list and returns the number of
// 1 jolt, 2 jolt and 3 jolt differences required to connet the 0 jolt
// outlet to the device (max + 3)
// Will error if a link cannot be found
func calculateFullChainDifferences(adapters []int) (int, int, int, error) {
	var count1, count2, count3, prevOut, startPos int

	for i := startPos; i < len(adapters); i++ {
		diff := adapters[i] - prevOut

		switch diff {
		case 1:
			count1++
		case 2:
			count2++
		case 3:
			count3++
		default:
			return 0, 0, 0, errors.New(fmt.Sprintf("Unsupported jolt difference. Current adapter: %v, Previous: %v", adapters[i], prevOut))
		}

		prevOut = adapters[i]
	}

	// Count 3 incremented to account for diff to device
	count3++

	return count1, count2, count3, nil
}

// Calculates the total number of possible permutations for getting from
// 0J -> (max +1)J with the available adapters
// The 3J differences from Part 1 are not permutable and can be used to break up the
// full adapter list into managable sub lists
func calculateNumberAdapterPermutations(adapters []int) int64 {
	var totalPermutations int64 = 1

	// add 0 to the start of the adapters, it is relevant as we could skip
	// a 1 or 2 Jolt adapter
	adapters = append([]int{0}, adapters...)

	subLists := getSubAdapterLists(adapters)

	for _, subList := range subLists {
		totalPermutations *= calculateSubListPermutations(subList)
	}

	return totalPermutations
}

// Split adapter list based on 3J jumps
func getSubAdapterLists(adapters []int) [][]int {
	subLists := [][]int{}
	prevOut := 0

	// Position for the start of the next sub list
	prevBoundary := 0

	for i := 0; i < len(adapters); i++ {
		diff := adapters[i] - prevOut

		if diff == 3 {
			subLists = append(subLists, adapters[prevBoundary:i])
			prevBoundary = i
		}

		prevOut = adapters[i]
	}

	// Add the last sub list
	subLists = append(subLists, adapters[prevBoundary:len(adapters)])

	return subLists
}

// Node to contain adapter value and refs to
// +1, +2 and +3 nodes which may be nil
type JoltNode struct {
	value int
	l1    *JoltNode
	l2    *JoltNode
	l3    *JoltNode
}

// Creates a totally unnecessary tree of adapters to subsequent possible adapters
// only the array is really required, but this was my first time using pointers
// so I don't mind the over-engineering too much
func calculateSubListPermutations(subList []int) int64 {

	// Trivial cases
	if len(subList) <= 2 {
		return 1
	}

	// array containing pointers to all nodes of the resulting permutation tree
	nodes := []*JoltNode{
		&JoltNode{value: subList[0]},
	}

	// Check each adapter node in nodes (which will get bigger)
	for n := 0; n < len(nodes); n++ {

		currentNode := nodes[n]
		subListPosition := sort.SearchInts(subList, currentNode.value)

		// Check up to the next 3 adapters, the 4th can never be in range
		for i := 1; i < 4; i++ {
			checkIndex := subListPosition + i

			// reached the end
			if checkIndex >= len(subList) {
				break
			}

			switch subList[checkIndex] {
			case currentNode.value + 1:
				node := JoltNode{value: subList[checkIndex]}
				nodes = append(nodes, &node)
				currentNode.l1 = &node
			case currentNode.value + 2:
				node := JoltNode{value: subList[checkIndex]}
				nodes = append(nodes, &node)
				currentNode.l2 = &node
			case currentNode.value + 3:
				node := JoltNode{value: subList[checkIndex]}
				nodes = append(nodes, &node)
				currentNode.l3 = &node
			default:
				continue
			}
		}
	}

	var permutations int64

	// Check how many nodes we have containing the last element
	// This is equal to the number of permutations
	for _, node := range nodes {
		if node.value == subList[len(subList)-1] {
			permutations++
		}
	}

	return permutations
}

// Parses the string input returning a sorted array of adapter ratings
func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	adapters := []int{}

	for _, line := range lines {
		adapters = append(adapters, getInt(line))
	}

	sort.Ints(adapters)

	return adapters
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

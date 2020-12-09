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
	xmasData := parseInput(XMASEncryptedInput)
	n := 25

	nonSum, err := findNonSumElement(xmasData, n)

	if err != nil {
		log.Fatalf("Encountered an error searching data: %v\n", err)
	}

	fmt.Println("Part 1:")
	fmt.Printf("%v is not a sum of its previous %v\n", nonSum, n)

	contSum, sumErr := findContiguousSum(xmasData, nonSum)

	if sumErr != nil {
		log.Fatalf("Encountered an error searching for sum: %v\n", err)
	}

	fmt.Println("Part 2:")
	fmt.Printf("%v sum to %v\n", contSum, nonSum)
	fmt.Printf("Min-Max sum is %v\n", sumMinMaxElements(contSum))
}

func parseInput(input string) []int {
	ints := []int{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		ints = append(ints, getInt(line))
	}

	return ints
}

func sumMinMaxElements(collection []int) int {
	sort.Ints(collection)

	return collection[0] + collection[len(collection)-1]
}

// finds the first contiguous set in collection that sums to target
// returns an error if none can be found
func findContiguousSum(collection []int, target int) ([]int, error) {

	// For each starting position
	for i, _ := range collection {
		var runningSum int
		sumComponents := []int{}

		for j := i; j < len(collection); j++ {
			if collection[j] > target {
				break
			}

			runningSum += collection[j]

			if runningSum > target {
				break
			}

			sumComponents = append(sumComponents, collection[j])

			// found a contiguous sum
			if runningSum == target {
				return sumComponents, nil
			}
		}
	}

	return []int{}, errors.New("Could not find a contiguous sum")
}

// For the given array, finds the first int that is not a sum of
// two elements in the preceding n entries
// Returns an error if all entries are sums or if n >= len(collection)
func findNonSumElement(collection []int, n int) (int, error) {

	if n >= len(collection) {
		return 0, errors.New(fmt.Sprintf("n: %v is too big", n))
	}

	for i := n + 1; i < len(collection); i++ {
		lastN := make([]int, n)

		copy(lastN, collection[i-n:i])

		if !hasSumInInts(lastN, collection[i]) {
			return collection[i], nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Could not find a missing sum in %v rolling entries", n))
}

// Returns true if two different entries in the given []int sum to
// the given target entry
// Will sort the given array
func hasSumInInts(collection []int, targetSum int) bool {
	if len(collection) == 0 {
		return false
	}

	sort.Ints(collection)

	for i, num := range collection {

		// If we're checking the last element, there's no sum here
		if i == len(collection)-1 {
			return false
		}

		// unchecked ints other than num in the collection
		uncheckedInts := make([]int, len(collection)-1-i)

		// we have checked elements before num, only copy elements after num
		if i < len(collection)-1 {
			copy(uncheckedInts, collection[i+1:])
		}

		if containsInt(uncheckedInts, targetSum-num) {
			return true
		}
	}

	return false
}

// Returns true if the given collection of ints contains the target int
// Will sort the collection if not already sorted
func containsInt(collection []int, target int) bool {

	if len(collection) == 0 {
		return false
	}

	if !sort.IntsAreSorted(collection) {
		sort.Ints(collection)
	}

	pos := sort.SearchInts(collection, target)

	if pos == len(collection) {
		return false
	}

	return collection[pos] == target
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

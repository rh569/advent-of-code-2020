package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Answer: %v\n", part1(ShuttleBusInput))

	fmt.Println("Part 2:")
	fmt.Printf("Answer: %v\n", part2(ShuttleBusInput))
}

func part2(input string) int64 {
	_, buses := parseInput(input)

	return getStartTimeOfBusParade(buses)
}

// returns the timestamp at which the first bus in the parade departs
func getStartTimeOfBusParade(buses []int) int64 {
	allDeparted := false
	var initialTime int64 = 0
	var timeStep int64 = 1
	var n int64 = 1

	for !allDeparted {
		var time int64 = initialTime + (n * timeStep)
		zeroes := 0
		stepChanged := false

		// fmt.Printf("n: %v, ti: %v, ts: %v, t: %v\n", n, initialTime, timeStep, time)

		for busOffset := 0; busOffset < len(buses); busOffset++ {
			bus := buses[busOffset]

			if bus == 0 {
				zeroes++

				// exit condition
				if zeroes == len(buses) {
					allDeparted = true
					break
				}

				continue
			}

			success, _ := checkForSpecificBus(time+int64(busOffset), int64(bus))

			if success {
				initialTime = time
				timeStep *= int64(bus)
				stepChanged = true
				buses[busOffset] = 0 // no longer consider factored bus
				// fmt.Println(buses)
			}
		}

		if stepChanged {
			n = 0
		} else {
			n++
		}
	}

	return initialTime
}

// returns true, 0 if a bus is found departing at that time
// otherwise returns false, time%bus
func checkForSpecificBus(time int64, bus int64) (bool, int64) {
	rem := time % bus

	if rem == 0 {
		return true, 0
	}

	return false, rem
}

func part1(input string) int {
	time, buses := parseInput(input)
	nextBus, departs := getNextBus(time, buses)

	return nextBus * (departs - time)
}

// returns bus number, departure time
func getNextBus(startTime int, buses []int) (int, int) {
	currentTime := startTime
	nextBus := 0

	for nextBus == 0 {
		success, bus := checkForAnyBus(currentTime, buses)

		if success {
			nextBus = bus
		} else {
			currentTime++
		}
	}

	return nextBus, currentTime
}

// returns true, busId if a bus is found departing at that time
func checkForAnyBus(currentTime int, buses []int) (bool, int) {

	for _, bus := range buses {

		if bus == 0 {
			continue
		}

		if currentTime%bus == 0 {
			return true, bus
		}
	}

	return false, 0
}

// returns the earlist time, and all buses - buses with id 0 have stopped service
func parseInput(input string) (int, []int) {
	lines := strings.Split(input, "\n")

	earliest := getInt(lines[0])

	busStrings := strings.Split(lines[1], ",")
	buses := []int{}

	for _, bus := range busStrings {

		if bus != "x" {
			buses = append(buses, getInt(bus))
		} else {
			buses = append(buses, 0)
		}

	}

	return earliest, buses
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

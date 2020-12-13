package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Manhattan distance travelled: %v\n", part1())

	fmt.Println("Part 2:")
	fmt.Printf("Manhattan distance travelled: %v\n", part2())
}

func part1() int {
	var startLong, startLat, startFacing = 0, 0, 90
	ferry := Ferry{
		Moveable{startLong, startLat},
		startFacing,
		Waypoint{Moveable{0, 0}},
	}
	instructions := parseInput(NavigationInput)

	ferry.processInstructions(instructions)

	return calculateManhattanDistance(startLong, startLat, ferry.longitude, ferry.latitude)
}

func part2() int {
	var startLong, startLat = 0, 0
	ferry := Ferry{
		waypoint: Waypoint{Moveable{10, 1}},
	}
	instructions := parseInput(NavigationInput)

	ferry.processWaypointInstructions(instructions)

	return calculateManhattanDistance(startLong, startLat, ferry.longitude, ferry.latitude)
}

func calculateManhattanDistance(long1 int, lat1 int, long2 int, lat2 int) int {
	longDist := float64(long1 + long2)
	latDist := float64(lat1 + lat2)

	return int(math.Abs(longDist) + math.Abs(latDist))
}

func parseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")

	instructions := []Instruction{}

	for _, line := range lines {
		instructions = append(instructions, Instruction{line[:1], getInt(line[1:])})
	}

	return instructions
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

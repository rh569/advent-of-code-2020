package main

import (
	"fmt"
)

const F_NORTH, F_EAST, F_SOUTH, F_WEST = 0, 90, 180, 270

type Ferry struct {
	Moveable
	facing   int // 0, 90, 180, 270 deg; 0 is North, 90 is East, etc.
	waypoint Waypoint
}

type Instruction struct {
	op  string // Operation N, S, R, L, etc.
	val int    // Value for operation
}

func (f *Ferry) moveForwards(distance int) {

	switch f.facing {
	case F_NORTH:
		f.latitude += distance
	case F_EAST:
		f.longitude += distance
	case F_SOUTH:
		f.latitude -= distance
	case F_WEST:
		f.longitude -= distance
	}
}

func (f *Ferry) turn(direction string, degrees int) {

	if direction == "R" {
		f.facing = (f.facing + degrees) % 360
	} else {
		f.facing = (f.facing + 360 - degrees) % 360
	}
}

func (f *Ferry) processInstructions(instructions []Instruction) {

	for _, ins := range instructions {

		if ins.op == "R" || ins.op == "L" {
			f.turn(ins.op, ins.val)
			continue
		}

		if ins.op == "F" {
			f.moveForwards(ins.val)
			continue
		}

		f.moveDirection(ins.op, ins.val)
	}
}

func (f *Ferry) moveToWaypoint() {
	f.longitude += f.waypoint.longitude
	f.latitude += f.waypoint.latitude
}

func (f *Ferry) processWaypointInstructions(instructions []Instruction) {

	for _, ins := range instructions {

		if ins.op == "R" || ins.op == "L" {
			f.waypoint.turn(ins.op, ins.val)
			continue
		}

		if ins.op == "F" {

			for i := 0; i < ins.val; i++ {
				f.moveToWaypoint()
			}
			continue
		}

		f.waypoint.moveDirection(ins.op, ins.val)
	}
}

func (f *Ferry) print() {
	fmt.Printf("(%v, %v) facing: %v\n", f.longitude, f.latitude, f.facing)
}

type Waypoint struct {
	Moveable
}

func (w *Waypoint) turn(direction string, degrees int) {
	currentLong := w.longitude
	currentLat := w.latitude

	if degrees == 180 {
		w.longitude = -currentLong
		w.latitude = -currentLat
		return
	}

	if (direction == "R" && degrees == 90) || (direction == "L" && degrees == 270) {
		w.longitude = currentLat
		w.latitude = -currentLong
		return
	}

	if (direction == "R" && degrees == 270) || (direction == "L" && degrees == 90) {
		w.longitude = -currentLat
		w.latitude = currentLong
	}
}

type Moveable struct {
	longitude int // East-West; -ve is West, +ve is East
	latitude  int // North-South; -ve is South, +ve is North
}

func (m *Moveable) moveDirection(direction string, distance int) {

	switch direction {
	case "N":
		m.latitude += distance
	case "E":
		m.longitude += distance
	case "S":
		m.latitude -= distance
	case "W":
		m.longitude -= distance
	}
}

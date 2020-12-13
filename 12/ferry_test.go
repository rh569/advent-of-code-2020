package main

import (
	"testing"
)

// --- Forwards

func TestMoveForwardsEast(t *testing.T) {
	ferry := Ferry{facing: 90}
	want := 11

	ferry.moveForwards(11)

	if ferry.longitude != want {
		t.Fatalf("Longitude: %v, Wanted: %v\n", ferry.longitude, want)
	}
}

func TestMoveForwardsSouth(t *testing.T) {
	ferry := Ferry{facing: 180}
	want := -13

	ferry.moveForwards(13)

	if ferry.latitude != want {
		t.Fatalf("Latitude: %v, Wanted: %v\n", ferry.latitude, want)
	}
}

func TestMoveForwardsTwice(t *testing.T) {
	ferry := Ferry{facing: 270}
	want := -10

	ferry.moveForwards(3)
	ferry.moveForwards(7)

	if ferry.longitude != want {
		t.Fatalf("Longitude: %v, Wanted: %v\n", ferry.longitude, want)
	}
}

// --- Turning

func TestTurnRight90(t *testing.T) {
	ferry := Ferry{facing: 0}
	want := 90

	ferry.turn("R", 90)

	if ferry.facing != want {
		t.Fatalf("Facing: %v, Wanted: %v\n", ferry.facing, want)
	}
}

func TestTurnLeft180(t *testing.T) {
	ferry := Ferry{facing: 180}
	want := 0

	ferry.turn("L", 180)

	if ferry.facing != want {
		t.Fatalf("Facing: %v, Wanted: %v\n", ferry.facing, want)
	}
}

func TestTurnRightPast360(t *testing.T) {
	ferry := Ferry{facing: 270}
	want := 90

	ferry.turn("R", 180)

	if ferry.facing != want {
		t.Fatalf("Facing: %v, Wanted: %v\n", ferry.facing, want)
	}
}

func TestTurnLeftPast0(t *testing.T) {
	ferry := Ferry{facing: 90}
	want := 270

	ferry.turn("L", 180)

	if ferry.facing != want {
		t.Fatalf("Facing: %v, Wanted: %v\n", ferry.facing, want)
	}
}

// --- Move Direction

func TestMoveDirectionSouth(t *testing.T) {
	ferry := Ferry{facing: 90}
	wantLat := -11
	wantLong := 0

	ferry.moveDirection("S", 11)

	if ferry.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", ferry.longitude, wantLong)
	}

	if ferry.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", ferry.latitude, wantLat)
	}

	if ferry.facing != 90 {
		t.Fatalf("Facing: %v, Wanted: %v\n", ferry.facing, 90)
	}
}

// --- Compound

func TestProcessInstructions(t *testing.T) {
	ferry := Ferry{facing: 90}
	instructions := parseInput(testInput)

	ferry.processInstructions(instructions)

	wantLong := 17
	wantLat := -8

	if ferry.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", ferry.longitude, wantLong)
	}

	if ferry.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", ferry.latitude, wantLat)
	}
}

// -- Waypoint

func TestWaypointTurnRight90(t *testing.T) {
	wp := Waypoint{
		Moveable{-11, 15},
	}

	wantLong := 15
	wantLat := 11

	wp.turn("R", 90)

	if wp.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", wp.longitude, wantLong)
	}

	if wp.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", wp.latitude, wantLat)
	}
}

func TestWaypointTurnLeft90(t *testing.T) {
	wp := Waypoint{
		Moveable{-11, 15},
	}

	wantLong := -15
	wantLat := -11

	wp.turn("L", 90)

	if wp.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", wp.longitude, wantLong)
	}

	if wp.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", wp.latitude, wantLat)
	}
}

func TestWaypointTurnLeft180(t *testing.T) {
	wp := Waypoint{
		Moveable{-11, 15},
	}

	wantLong := 11
	wantLat := -15

	wp.turn("L", 180)

	if wp.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", wp.longitude, wantLong)
	}

	if wp.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", wp.latitude, wantLat)
	}
}

func TestProcessWaypointInstructions(t *testing.T) {
	ferry := Ferry{waypoint: Waypoint{Moveable{10, 1}}}
	instructions := parseInput(testInput)

	ferry.processWaypointInstructions(instructions)

	wantLong := 214
	wantLat := -72

	if ferry.longitude != wantLong {
		t.Fatalf("Longitude: %v, Wanted: %v\n", ferry.longitude, wantLong)
	}

	if ferry.latitude != wantLat {
		t.Fatalf("Latitude: %v, Wanted: %v\n", ferry.latitude, wantLat)
	}
}

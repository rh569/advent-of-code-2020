package tile

import (
	"reflect"
	"testing"
)

var testString = `Tile 121:
..##
.###
.#.#
#.#.`

func TestFromString(t *testing.T) {
	wantId := 121
	wantSides := [4]int{2, 3, 2, 1}
	wantData := [][]int{
		[]int{0, 0, 1, 1},
		[]int{0, 1, 1, 1},
		[]int{0, 1, 0, 1},
		[]int{1, 0, 1, 0},
	}

	result := FromString(testString)

	if result.Id != wantId {
		t.Fatalf("ID - Result: %v, want %v\n", result.Id, wantId)
	}

	if !reflect.DeepEqual(result.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", result.Sides, wantSides)
	}

	if !reflect.DeepEqual(result.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", result.Data, wantData)
	}
}

func TestFlipH(t *testing.T) {
	wantSides := [4]int{2, 1, 2, 3}

	wantData := [][]int{
		[]int{1, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{1, 0, 1, 0},
		[]int{0, 1, 0, 1},
	}

	tile := FromString(testString)
	tile.FlipH()

	if !reflect.DeepEqual(tile.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", tile.Sides, wantSides)
	}

	if !reflect.DeepEqual(tile.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", tile.Data, wantData)
	}
}

func TestFlipV(t *testing.T) {
	wantSides := [4]int{2, 3, 2, 1}

	wantData := [][]int{
		[]int{1, 0, 1, 0},
		[]int{0, 1, 0, 1},
		[]int{0, 1, 1, 1},
		[]int{0, 0, 1, 1},
	}

	tile := FromString(testString)
	tile.FlipV()

	if !reflect.DeepEqual(tile.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", tile.Sides, wantSides)
	}

	if !reflect.DeepEqual(tile.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", tile.Data, wantData)
	}
}

func TestTranspose(t *testing.T) {
	wantSides := [4]int{1, 2, 3, 2}

	wantData := [][]int{
		[]int{0, 0, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{1, 1, 0, 1},
		[]int{1, 1, 1, 0},
	}

	tile := FromString(testString)
	tile.transpose()

	if !reflect.DeepEqual(tile.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", tile.Sides, wantSides)
	}

	if !reflect.DeepEqual(tile.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", tile.Data, wantData)
	}
}

func TestRotate90(t *testing.T) {
	wantSides := [4]int{1, 2, 3, 2}

	wantData := [][]int{
		[]int{1, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{1, 0, 1, 1},
		[]int{0, 1, 1, 1},
	}

	tile := FromString(testString)
	tile.Rotate90()

	if !reflect.DeepEqual(tile.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", tile.Sides, wantSides)
	}

	if !reflect.DeepEqual(tile.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", tile.Data, wantData)
	}
}

func TestRotate90Counter(t *testing.T) {
	wantSides := [4]int{3, 2, 1, 2}

	wantData := [][]int{
		[]int{1, 1, 1, 0},
		[]int{1, 1, 0, 1},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 1},
	}

	tile := FromString(testString)
	tile.Rotate90Counter()

	if !reflect.DeepEqual(tile.Sides, wantSides) {
		t.Fatalf("Sides - Result: %v, want %v\n", tile.Sides, wantSides)
	}

	if !reflect.DeepEqual(tile.Data, wantData) {
		t.Fatalf("Data - Result: %v, want %v\n", tile.Data, wantData)
	}
}

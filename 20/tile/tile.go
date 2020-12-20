package tile

import (
	"fmt"
	"strconv"
	"strings"
)

const TOP = 0
const RIGHT = 1
const BOT = 2
const LEFT = 3

type Tile struct {
	Data  [][]int
	Sides [4]int
	Id    int
}

func (t *Tile) FlipH() {
	size := len(t.Data)
	newData := makeEmptyData(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newData[i][size-j-1] = t.Data[i][j]
		}
	}

	t.Data = newData

	right := t.Sides[RIGHT]
	left := t.Sides[LEFT]
	t.Sides[RIGHT] = left
	t.Sides[LEFT] = right
}

func (t *Tile) FlipV() {
	size := len(t.Data)
	newData := makeEmptyData(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newData[size-i-1][j] = t.Data[i][j]
		}
	}

	t.Data = newData

	top := t.Sides[TOP]
	bot := t.Sides[BOT]
	t.Sides[TOP] = bot
	t.Sides[BOT] = top
}

func (t *Tile) transpose() {
	size := len(t.Data)
	newData := makeEmptyData(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newData[j][i] = t.Data[i][j]
		}
	}

	t.Data = newData

	top, right, bot, left := t.Sides[TOP], t.Sides[RIGHT], t.Sides[BOT], t.Sides[LEFT]

	t.Sides[TOP] = left
	t.Sides[RIGHT] = bot
	t.Sides[BOT] = right
	t.Sides[LEFT] = top
}

func (t *Tile) Rotate90() {
	t.transpose()
	t.FlipH()
}

func (t *Tile) Rotate90Counter() {
	t.transpose()
	t.FlipV()
}

func (t *Tile) Print() {
	for _, row := range t.Data {
		for _, point := range row {
			if point == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (t *Tile) CountPoints() int {
	count := 0

	for _, row := range t.Data {
		for _, point := range row {
			if point == 1 {
				count++
			}
		}
	}

	return count
}

// func (t *Tile) rotate180() {
// 	t.flipV()
// 	t.flipH()
// }

func makeEmptyData(size int) [][]int {
	newData := make([][]int, size)

	for i := range newData {
		newData[i] = make([]int, size)
	}

	return newData
}

func FromString(input string) Tile {
	lines := strings.Split(input, "\n")

	tile := Tile{}

	idStr := strings.Replace(lines[0], "Tile ", "", 1)
	idStr = strings.Replace(idStr, ":", "", 1)
	tile.Id = getInt(idStr)

	lines = lines[1:]

	data := makeEmptyData(len(lines))

	for i, l := range lines {
		for j, r := range l {

			if r == '#' {
				data[i][j] = 1
			}
		}
	}

	tile.Data = data

	tile.Sides = calculateSides(data)

	return tile
}

func calculateSides(data [][]int) [4]int {
	var top, right, bot, left int

	for i := 0; i < len(data); i++ {
		if data[i][0] == 1 {
			left++
		}

		if data[i][len(data)-1] == 1 {
			right++
		}
	}

	for j := 0; j < len(data); j++ {
		if data[0][j] == 1 {
			top++
		}

		if data[len(data)-1][j] == 1 {
			bot++
		}
	}

	return [4]int{top, right, bot, left}
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

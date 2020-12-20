package main

import (
	"fmt"
	"strings"

	"github.com/rh569/advent-of-code-2020/20/tile"
)

const IMAGE_SIZE = 12

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Product of corners: %v\n", part1(SatelliteInput))

	fmt.Println("Part 2:")
	fmt.Printf("Not Nessy: %v\n", part2(SatelliteInput))
}

func parseInput(input string) []tile.Tile {
	tileStrs := strings.Split(input, "\n\n")

	tiles := []tile.Tile{}

	for _, t := range tileStrs {
		tiles = append(tiles, tile.FromString(t))
	}

	return tiles
}

func part1(input string) int {

	allTiles := parseInput(input)

	image := arrangeImageTiles(allTiles)

	// for _, r := range image {
	// 	for _, t := range r {
	// 		fmt.Printf("%v ", t.Id)
	// 	}
	// 	fmt.Print("\n")
	// }

	cornersProduct := image[0][0].Id
	cornersProduct *= image[0][len(image[0])-1].Id
	cornersProduct *= image[len(image)-1][0].Id
	cornersProduct *= image[len(image)-1][len(image[0])-1].Id

	return cornersProduct
}

func part2(input string) int {
	allTiles := parseInput(input)
	image := arrangeImageTiles(allTiles)

	fullImage := finaliseImage(image)

	count := fullImage.CountPoints()

	fmt.Printf("Points: %v\n", count)

	maxNess := 0
	nessSize := 15

	// For each flipside
	for f := 0; f < 2; f++ {

		// For each orientation
		for s := 0; s < 4; s++ {
			nesses := countMonsters(fullImage)

			if nesses > maxNess {
				fmt.Printf("Found %v, when f=%v, s=%v\n", nesses, f, s)
				// fullImage.Print()
				maxNess = nesses
			}

			fullImage.Rotate90()
		}

		fullImage.FlipH()
	}

	return count - maxNess*nessSize
}

func countMonsters(image tile.Tile) int {
	// .#...#.###...#.##.O#
	// O.##.OO#.#.OO.##.OOO
	// #O.#O#.O##O..O.#O##.

	ness := [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		[]int{1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1},
		[]int{0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0},
	}

	count := 0

	// For each starting point in the data
	for tI, tRow := range image.Data {
		if tI+len(ness) >= len(image.Data) {
			continue
		}

		for tJ, _ := range tRow {
			if tJ+len(ness[0]) >= len(tRow) {
				continue
			}

			found := true

			// Check for nessy
			for i, row := range ness {
				for j, point := range row {

					if point == 1 && image.Data[tI+i][tJ+j] == 0 {
						found = false
					}
				}
			}

			if found {
				count++
			}
		}
	}

	return count
}

// Remove tile borders and return the full image as one giant tile
// so that it can be rotated, flipped etc.
func finaliseImage(arrangedTiles [][]tile.Tile) tile.Tile {
	tileSizeSansBorder := len(arrangedTiles[0][0].Data[0]) - 2
	finalImageSize := IMAGE_SIZE * tileSizeSansBorder

	finalImage := tile.Tile{}
	finalData := make([][]int, finalImageSize)

	for i := range finalData {
		finalData[i] = make([]int, finalImageSize)
	}

	for tI, tileRow := range arrangedTiles {
		for tJ, tile := range tileRow {

			for i, dataRow := range tile.Data {
				if i == 0 || i == len(tile.Data)-1 {
					continue
				}

				for j, point := range dataRow {
					if j == 0 || j == len(dataRow)-1 {
						continue
					}

					if point == 1 {
						finalData[tI*tileSizeSansBorder+i-1][tJ*tileSizeSansBorder+j-1] = 1
					}
				}
			}
		}
	}

	finalImage.Data = finalData
	return finalImage
}

// Arrange all the tiles into a matching 2D image
func arrangeImageTiles(allTiles []tile.Tile) [][]tile.Tile {
	seedTile := allTiles[0]
	remainingTiles := allTiles[1:]

	initialRow, tilesLeft := assembleRow(seedTile, remainingTiles)
	remainingTiles = tilesLeft

	image := [][]tile.Tile{initialRow}

	image = assembleRemainingRows(image, remainingTiles)

	return image
}

// Piece together the remaining rows to form the full image
func assembleRemainingRows(image [][]tile.Tile, remainingTiles []tile.Tile) [][]tile.Tile {

	currentTop := image[0][0]
	currentBot := image[0][0]

	i := 0

	for len(image) < IMAGE_SIZE {
		if i >= len(remainingTiles) {
			i = 0
		}

		potentialTile := remainingTiles[i]
		matchFound := false

		// For each face
		for f := 0; f < 2; f++ {

			// For 4 sides
			// Check sums of sides before checking for actual matches
			for s := 0; s < 4; s++ {

				// Check Top
				if currentTop.Sides[tile.TOP] == potentialTile.Sides[tile.BOT] {

					// Compare bottom of potential to top of currentTop
					if matchTopBottom(potentialTile, currentTop) {
						// Remove matched tile
						remainingTiles = append(remainingTiles[:i], remainingTiles[i+1:]...)

						// generate the whole row
						newRow, newRemaining := assembleRow(potentialTile, remainingTiles)
						remainingTiles = newRemaining

						// resize and copy row into image
						newImage := make([][]tile.Tile, len(image)+1)
						newImage[0] = newRow
						copy(newImage[1:], image)
						image = newImage

						currentTop = image[0][0]

						matchFound = true
						break
					}
				}

				// Check bottom of currentBot to top of potential
				if currentBot.Sides[tile.BOT] == potentialTile.Sides[tile.TOP] {

					// Compare
					if matchTopBottom(currentBot, potentialTile) {
						// Remove matched tile
						remainingTiles = append(remainingTiles[:i], remainingTiles[i+1:]...)

						// generate the whole row
						newRow, newRemaining := assembleRow(potentialTile, remainingTiles)
						remainingTiles = newRemaining

						image = append(image, newRow)

						currentBot = image[len(image)-1][0]

						matchFound = true
						break
					}
				}

				// Rotate and try again
				potentialTile.Rotate90()
			}

			if matchFound {
				i--
				break
			}

			// Flip and try again
			potentialTile.FlipH()
		}

		i++
	}

	return image
}

// Piece together a row from a single starting tile
func assembleRow(seedTile tile.Tile, remainingTiles []tile.Tile) ([]tile.Tile, []tile.Tile) {
	// tileSize := len(seedTile.Data)

	currentLeft := seedTile
	currentRight := seedTile

	assembledRow := []tile.Tile{currentLeft}
	i := 0

	for len(assembledRow) < IMAGE_SIZE {
		if i == len(remainingTiles) {
			i = 0
		}

		potentialTile := remainingTiles[i]
		matchFound := false

		// For each face
		for f := 0; f < 2; f++ {

			// For 4 sides
			// Check sums of sides before checking for actual matches
			for s := 0; s < 4; s++ {

				// Check Left
				if currentLeft.Sides[tile.LEFT] == potentialTile.Sides[tile.RIGHT] {

					// Compare right of potential to left of currentLeft
					if matchLeftRight(potentialTile, currentLeft) {
						// add to start
						newRow := make([]tile.Tile, len(assembledRow)+1)
						newRow[0] = potentialTile
						copy(newRow[1:], assembledRow)
						assembledRow = newRow

						// new top tile
						currentLeft = assembledRow[0]

						matchFound = true
						break
					}
				}

				// Check Right
				if currentRight.Sides[tile.RIGHT] == potentialTile.Sides[tile.LEFT] {

					// Compare right of currentRight to left of potential
					if matchLeftRight(currentRight, potentialTile) {
						// add to end
						assembledRow = append(assembledRow, potentialTile)

						// new bot tile
						currentRight = assembledRow[len(assembledRow)-1]

						matchFound = true
						break
					}
				}

				// Rotate and try again
				potentialTile.Rotate90()
			}

			if matchFound {
				remainingTiles = append(remainingTiles[:i], remainingTiles[i+1:]...)
				i--
				break
			}

			// Flip and try again
			potentialTile.FlipH()
		}

		i++
	}

	return assembledRow, remainingTiles
}

// Checks if the right side of tileLeft matches the left side of tileRight
func matchLeftRight(tileLeft, tileRight tile.Tile) bool {
	match := true

	tileSize := len(tileLeft.Data)

	for i := range tileLeft.Data {
		if tileLeft.Data[i][tileSize-1] != tileRight.Data[i][0] {
			return false
		}
	}

	return match
}

// Checks if the bottom side of tileTop matches the top side of tileBot
func matchTopBottom(tileTop, tileBot tile.Tile) bool {
	match := true

	tileSize := len(tileTop.Data)

	for j, n := range tileTop.Data[tileSize-1] {
		if n != tileBot.Data[0][j] {
			return false
		}
	}

	return match
}

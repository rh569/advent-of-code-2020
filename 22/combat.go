package main

import "fmt"

func main() {
	fmt.Println("Part 1:")
	fmt.Printf("Winning score: %v\n", part1(CombatInput))

	fmt.Println("Part 2:")
	fmt.Printf("Winning score: %v\n", part2(CombatInput))

	// NOT 32827 (too low)
	// NOT 33099 (too low)
	// 32827 again...
}

func part1(input string) int {
	decks := parseInput(input)

	winningDeck := playCombat(decks)

	score := calculateScore(winningDeck)
	return score
}

func part2(input string) int {
	decks := parseInput(input)

	finishedDecks, _ := playRecursiveCombatGame(decks)

	decks = finishedDecks

	var winningDeck []int

	if len(decks[0]) > 0 {
		winningDeck = decks[0]
	} else {
		winningDeck = decks[1]
	}

	score := calculateScore(winningDeck)
	return score
}

func calculateScore(winningDeck []int) int {
	score := 0

	for i := 0; i < len(winningDeck); i++ {
		cardScore := len(winningDeck) - i
		score += (cardScore * winningDeck[i])
	}

	return score
}

// Simulates a game of 'Combat' with the two given decks
// Returns the winning deck in its final state
func playCombat(decks [2][]int) []int {
	deck1 := decks[0]
	deck2 := decks[1]

	for len(deck1) > 0 && len(deck2) > 0 {

		if deck1[0] > deck2[0] {
			deck1, deck2 = winTurn(deck1, deck2)
		} else {
			// Assuming no ties
			deck2, deck1 = winTurn(deck2, deck1)
		}
	}

	if len(deck1) > 0 {
		return deck1
	} else {
		return deck2
	}
}

// Rearranges decks based on order they're passed in (winning first)
// Moves top cards to end of winning deck
// Returns decks in same order as passed
func winTurn(winningDeck, losingDeck []int) ([]int, []int) {

	winningCard := winningDeck[0]
	losingCard := losingDeck[0]

	winningDeck = winningDeck[1:]
	losingDeck = losingDeck[1:]

	winningDeck = append(winningDeck, []int{winningCard, losingCard}...)

	return winningDeck, losingDeck
}

// Simulates a game of 'Recursive Combat' with the two given decks
// Returns both decks in their final state and whether the game was won by P1 by default
func playRecursiveCombatGame(decks [2][]int) ([2][]int, bool) {
	previousConfigurations := [][2][]int{}

	// Play until a deck length is 0 or P1 wins by default
	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		nextDecks, winsGameByDefault := playRound(decks, previousConfigurations)

		if winsGameByDefault {
			return nextDecks, true
		}

		previousConfigurations = append(previousConfigurations, decks)
		decks = nextDecks
	}

	return decks, false
}

func playRound(decks [2][]int, prevConfig [][2][]int) ([2][]int, bool) {
	// Check anti-loop condition
	if playedBefore(decks, prevConfig) {
		return decks, true
	}

	remainingCards1 := len(decks[0]) - 1
	remainingCards2 := len(decks[1]) - 1

	// Check recurse condition, if met play a recursive game
	if decks[0][0] <= remainingCards1 && decks[1][0] <= remainingCards2 {
		subDecks := [2][]int{}

		// Copy cards into subDecks
		for i, deck := range decks {
			subDeck := []int{}
			finalIndex := deck[0] + 1 // Off by 1 here and the tests still pass but the sub games are all wrong :(
			subDeck = append(subDeck, deck[1:finalIndex]...)
			subDecks[i] = subDeck
		}

		// Recurse with subDecks
		nextSubDecks, wonGameByDefault := playRecursiveCombatGame(subDecks)
		subDecks = nextSubDecks

		var winningDeckIndex int

		if wonGameByDefault || len(subDecks[0]) > 0 {
			winningDeckIndex = 0
		} else {
			winningDeckIndex = 1
		}

		winDeck, loseDeck := winTurn(decks[winningDeckIndex], decks[1-winningDeckIndex])

		decks[winningDeckIndex] = winDeck
		decks[1-winningDeckIndex] = loseDeck

		return decks, false
	}

	// Else, play a normal round
	var deck1, deck2 []int

	if decks[0][0] > decks[1][0] {
		deck1, deck2 = winTurn(decks[0], decks[1])
	} else {
		// Assuming no ties
		deck2, deck1 = winTurn(decks[1], decks[0])
	}

	decks[0] = deck1
	decks[1] = deck2

	return decks, false
}

// Checks to see if this deck configurations has been encountered before
func playedBefore(decks [2][]int, prevConfig [][2][]int) bool {
	playedBefore := false

	for _, prevDecks := range prevConfig {
		if deckPairsMatch(decks, prevDecks) {
			playedBefore = true
		}
	}

	return playedBefore
}

// Check each card for p1 and p2 are the same in each pair of decks
func deckPairsMatch(decksA, decksB [2][]int) bool {
	match := true

	// For each deck in the pair
	for i := 0; i < 2; i++ {

		// if deck lengths don't match
		if len(decksA[i]) != len(decksB[i]) {
			return false
		}

		// For each card in the A deck
		for j, card := range decksA[i] {

			// Not the same if any card fails to match
			if card != decksB[i][j] {
				return false
			}
		}
	}

	return match
}

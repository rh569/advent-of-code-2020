package main

import (
	"strconv"
	"strings"
)

func parseInput(input string) [2][]int {
	playerSections := strings.Split(input, "\n\n")
	decks := [2][]int{}

	for i, section := range playerSections {
		cardStrs := strings.Split(section, "\n")
		deck := []int{}

		for j, card := range cardStrs {
			if j == 0 {
				continue
			}

			deck = append(deck, getInt(card))
		}

		decks[i] = deck
	}

	return decks
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

var LoopTestInput = `Player 1:
43
19

Player 2:
2
29
14`

var TestInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

var CombatInput = `Player 1:
15
31
26
4
36
30
43
39
50
21
25
46
6
44
12
20
23
9
48
11
16
42
17
13
10

Player 2:
34
49
19
24
45
28
7
41
18
38
2
3
33
14
35
40
32
47
22
29
8
37
5
1
27`

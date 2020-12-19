package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Manually incremented to see if we captured any more matches...
// Not clever in any way, but it worked
// 4 is the minimum to capture every message for my input - final regex: 12,233 characters:...
const REPETITIONS = 4

func parseInput(input string) (*regexp.Regexp, []string) {
	sections := strings.Split(input, "\n\n")

	ruleRegexp := regexpFromRules(sections[0])
	messages := strings.Split(sections[1], "\n")

	return ruleRegexp, messages
}

type Rule struct {
	index int
	rule  string
}

func regexpFromRules(rulesInput string) *regexp.Regexp {
	rulesStr := strings.Split(rulesInput, "\n")
	rules := []Rule{}

	for _, line := range rulesStr {
		parts := strings.Split(line, ": ")

		rules = append(rules, Rule{getInt(parts[0]), parts[1]})
	}

	sort.Slice(rules, func(i, j int) bool {
		return rules[i].index < rules[j].index
	})

	regexpStr := "^" + buildRegexp(rules, rules[0].rule) + "$"

	fmt.Printf("Created regex consisting of %v characters:\n", len(regexpStr))

	return regexp.MustCompile(regexpStr)
}

// Recursively builds up a regexp string
// Assumes rules are in order and that their rule index matches their position in the array
func buildRegexp(rules []Rule, ruleStr string) string {

	if ruleStr == "\"a\"" || ruleStr == "\"b\"" {
		return strings.ReplaceAll(ruleStr, "\"", "")
	}

	// Special looping case - 8
	if ruleStr == "42 | 42 8" {
		s := "(?:"
		s += buildSimpleRulePart(rules, "42")
		s += ")+"
		return s
	}

	// Special looping case - 11
	if ruleStr == "42 31 | 42 11 31" {
		s := "(?:"

		for i := 0; i <= REPETITIONS; i++ {

			for j := 0; j < i; j++ {
				s += buildSimpleRulePart(rules, "42")
			}

			for j := 0; j < i; j++ {
				s += buildSimpleRulePart(rules, "31")
			}

			if i > 0 && i < REPETITIONS {
				s += "|"
			}
		}

		s += ")"
		return s
	}

	parts := strings.Split(ruleStr, " | ")

	// No OR
	if len(parts) == 1 {
		return buildSimpleRulePart(rules, parts[0])
	}

	if len(parts) == 2 {
		// Using non-capturing groups as there will be lots of them
		s := "(?:"
		s += buildSimpleRulePart(rules, parts[0])
		s += "|"
		s += buildSimpleRulePart(rules, parts[1])
		s += ")"

		return s
	} else {
		panic("Too many parts")
	}
}

func buildSimpleRulePart(rules []Rule, numsStr string) string {
	numStrs := strings.Split(numsStr, " ")

	var s string

	for _, n := range numStrs {
		s += buildRegexp(rules, rules[getInt(n)].rule)
	}

	return s
}

func getInt(str string) int {
	num, _ := strconv.ParseInt(str, 10, 0)
	return int(num)
}

var SatelliteInput = `104: 114 114
13: 114 77 | 29 105
100: 20 114 | 111 29
80: 98 29 | 102 114
81: 29 10 | 114 89
97: 114 51 | 29 124
94: 114 67 | 29 64
74: 65 114
131: 111 114 | 65 29
123: 118 114 | 39 29
17: 29 70 | 114 70
46: 89 29 | 94 114
110: 114 50 | 29 65
62: 29 33 | 114 72
5: 81 29 | 2 114
16: 29 12 | 114 103
69: 114 3 | 29 59
12: 121 29 | 128 114
43: 97 114 | 25 29
45: 64 114 | 48 29
133: 132 29 | 70 114
67: 114 114 | 29 29
106: 29 64 | 114 132
31: 29 92 | 114 38
10: 29 67 | 114 65
65: 29 114 | 29 29
113: 68 114 | 68 29
37: 64 114 | 104 29
7: 56 114 | 129 29
128: 115 29 | 112 114
117: 7 114 | 49 29
47: 29 86 | 114 60
32: 79 114 | 64 29
86: 111 29 | 65 114
91: 29 5 | 114 66
4: 79 29 | 68 114
3: 29 68 | 114 20
92: 29 43 | 114 62
70: 114 29
39: 114 127 | 29 35
83: 29 132 | 114 111
9: 29 125 | 114 63
130: 29 36 | 114 108
27: 111 114 | 132 29
41: 67 114 | 104 29
90: 114 50 | 29 20
89: 14 21
61: 114 47 | 29 126
50: 29 29 | 114 29
38: 114 9 | 29 57
29: "a"
63: 29 69 | 114 99
105: 114 70 | 29 132
77: 14 114 | 20 29
51: 19 29 | 129 114
109: 29 120 | 114 30
54: 65 114 | 79 29
30: 20 114 | 40 29
72: 29 22 | 114 28
111: 114 21 | 29 114
84: 24 29 | 17 114
127: 50 29 | 48 114
125: 29 58 | 114 88
75: 50 114 | 20 29
118: 114 106 | 29 76
19: 21 111
115: 29 50
98: 40 29 | 50 114
126: 29 134 | 114 4
103: 29 109 | 114 15
26: 114 50 | 29 64
95: 114 96 | 29 123
101: 45 29 | 78 114
18: 29 122 | 114 110
58: 74 29 | 4 114
102: 70 29 | 50 114
0: 8 11
28: 53 29 | 119 114
21: 29 | 114
11: 42 31
122: 114 111 | 29 65
66: 114 80 | 29 71
24: 20 114 | 50 29
25: 114 55 | 29 46
87: 36 114 | 32 29
55: 131 29 | 32 114
76: 70 29
44: 1 114 | 61 29
1: 52 114 | 130 29
96: 18 114 | 13 29
112: 48 114 | 64 29
71: 129 114 | 131 29
134: 114 50 | 29 48
22: 3 114 | 113 29
2: 107 29 | 26 114
99: 83 29 | 53 114
49: 100 114 | 37 29
35: 104 29 | 104 114
53: 48 29 | 132 114
48: 29 21 | 114 29
116: 29 17 | 114 6
121: 32 29 | 54 114
93: 29 16 | 114 44
119: 29 111 | 114 70
59: 79 29 | 48 114
34: 29 79 | 114 48
78: 29 64 | 114 79
107: 50 29 | 65 114
114: "b"
129: 14 29 | 104 114
15: 41 114 | 90 29
20: 21 114 | 29 29
60: 114 104
57: 117 114 | 82 29
52: 75 29 | 85 114
85: 70 29 | 40 114
40: 114 29 | 29 114
124: 23 114 | 34 29
33: 116 29 | 84 114
36: 114 68 | 29 132
88: 133 114 | 27 29
68: 29 114
8: 42
56: 29 132 | 114 40
6: 132 114 | 79 29
64: 114 114 | 114 29
23: 21 40
82: 87 114 | 101 29
14: 21 21
79: 29 29
120: 114 67 | 29 132
73: 95 29 | 91 114
42: 73 29 | 93 114
108: 29 48 | 114 111
132: 114 114 | 29 114

aaababbaaabababaaabaaabb
bbabbababbbabaababbbbbbabababaaa
baaaaaabaababbbbbaaaaaababbbbbaababbaaaaaaaaaaaa
bababbbbaabaabbaaabababbaabbabab
ababbaaabaabbabbbbaaabba
aaaabaaaababbbababbaababaabbabbababaaabbaabbbbabaabbabababbbabbbaababbaaabaaabbb
ababaababbaabbbaabbaabba
abaabbbabbaabbaaaabababaabbbaabaabbbbbaababbaabb
abbaaaabbbaabababbabbaabaaababaa
aababababaabaaabaabaabba
bbaabbbaabbabbaaaaaabbba
bbaabbaaabbabbaaaabaaaabababbaababbabbbaabbbaabb
abbabbbaabbbbabaabbabbbaaaabaaab
aaaabaabaababbbaaaababbb
baaaabaabbbbaaaabbaaabba
aabbaaaaabbbaababaaabbaaaaababaa
abbaababbaaabaaabbaaaaba
bbbbbbabaaaaaababbaabababbbaaaabbaaabbba
baaaabaaabababbbbaabbbaa
ababbbaabaaabaaabaabbaaabaabbbbaaaababab
babbbabaabaaaaaaaabbbaba
ababaabaabaaaaaabbaaaaab
aaaaabbabaaababbaaababbb
bbaaaaaaaababbbaaaaabbabbbaaabababaabbbaaaaabbaa
baaaabbababaabbbabaabaabaabbabbbaaabbaabbababaab
bbaabbabababaababbbbaabbaaababab
baaaabbbbaaababbabbbbaaa
ababbabababababaaaabaaba
abbabababbabaaabbbbbabab
ababbbbbaababbbbbbaabbbabbabbaba
abaaabbbbbabbaaabaabbbab
abbabbaabbbbbbabaababaaa
aaaaabaabaababbbaaaaabbaaababbaaaaabbbbbababbbbababbaababaabaaaaabbbbaabbbbbaabbbaabaaba
baabbababaabaaabbaaababa
bbabaabaabbaaaabbbaabbabbababbababbbbbbbaabaaabb
baabbabbaaabbbbbbbabbbba
bbaabbaaabaabaaabababbba
ababbaaabaabaaabbaabbaaababbaaab
baababbabaaaabaababbabaabbbbbbabbbbaabbbbabaaaaabbbabbbbabbaabba
abaabbbbababaabbababababaabbabab
bbbabbaaaabbbaabaaabaaaa
bbabaabbbbbbbbabbabbbaaa
abbaaaaaabaaababaabbaaba
bbaaabaabbabbbaaaaabbaaababbaaba
baabbaabbbaabaabaabaabba
aaaabbabaabbbaaabaaabbba
bbaaabaabbbaabaabbbbabab
babbbbaabbabbaaaabababba
bbbaabaabbbaabaabbaaababaaaabbbbaaaabaaabbbaabababbaaabb
baaaaabababababbaaaaaababbabbbbbbabbbbbbbabbaaaababbbbabaabaaaaabbababab
baabbaababbbbbbbbaaababbabaabbbabaaaaaaaaaaabbaabbabbbab
baaabaabbbaaaabbbabbabab
aaabaabbbbbabbbaaaaabbaa
bbbbbbabbbbababbabaaabbbaaaabbbbbaabbbbb
abababbbbaababbbbaaababbaabaabaabbbbabaabbaaaababbababbaabbabaab
babaabbabbbababbaababbbababbbbab
abaaaababaaababbbbabbabb
abaaabbbaabbbbbaabbbaaabbbaaaaabaaabaaaa
baaabaaaabbabbabbabbaaba
abaaaaaaaabbaaaabaabaabb
bbabaaabbaaaabbaabababba
aababbbaabaabaaabbbbbbba
abbabbabaabbabbabbababba
bbbbaabbaababaabaaaabaaabbaaabaa
baabbabbbbabaabaaabbabaa
aaabbabaabaaabbabbbaabbbbbaababbaaababbb
abbaababbaabbabaaabaabab
bbbaabbbbabbbbbabaaaaaabaabbaaababbaabba
abbabbabaababbbbababbabaabbabababbbbbaabbbbaabab
bbabbbbbaaabbababbababab
abbabbabaabaaabababaaaba
aabbabbaaabaabaaaaababbb
aaabbabaabaabaaaabbbbbaabbabaababbaabaabbbaaabbb
bbabaabaaaabbababaaaaaabbaaababbbabaabbbbbabbbbbbbbbabbbbbbbabbbbbbabbab
bbaaaaaababbbbaaaaaaaaab
bbaaabaabaaaabbababbbaab
ababbababababbaabbaaababaabbbbababbbaabb
ababbaaaabaabaaaaaaaaaab
baaaabbbbabbabaabaaabaaabaaaaabb
baaaaaabaaabbabaababbbab
abbabbaabbaaaaaabbaababbbbbaabbbaaaaaaab
baabaabaaababbbabbaababbababbbbbaabaaaaaabaabbabbaaabaabbaaabbbbbbbaaabaabbababa
abbbbaabbbbbbaaababbbbbaabababbbbbbbbaab
bbbababbbaaaaabaabbabaaababaabbabbaaaaabbabbabab
aaabbbbbaaabbbbbbabbaaba
bbaabbaaabaaabbbababbbba
baabbaaabbabaabbbbabbbbbbbababaaaaaaabbbabaaabaabaaabbabbbbabaaababaabaababbbbaa
bbaababbbbaaabbbbababbaabbbaaababbaabbaa
babbbbaabaabbaabbabbaabb
bbbbbbabaaababbabbbabaab
babbbbbaaaababbabbabbaba
bbaabbabaaabaabbbabababbbbaabbaababaaaabaaaaaabaaabaaaaabbabbbbabbababaaabbbabbb
baaaaabaabbaaaaababbbbab
babababaabaababababaabbbabbaabbbaabbaababaababaaabaaabab
bbbbbaaaababbbbbbabbabab
baabaaabaaabbbabababbbab
bbaaaabaaaaabbbababbbaabaaaaabba
ababababbbaabbbbbabaaaba
bbbabbbabbabaaabbbbabbaa
bababbaaabbbbabaabbaabbbbaabbabbbbaababaabbbaaabbbabbabb
aabababaabbabbabaababaab
babbbababbbaaabaabbaabba
bbabaababbaaaaaabbbbbbbb
baababbaabaabaabbbbbabab
baaabbaabbaabaababbaabbbbaaaabababbbbabbbbbbabba
abbabbaaabbbbabaabbbbaabababbbbbaaabbbababbaabaabbbbbaba
abbaaaaaababaabbaababababbbaabaabaaaabbabbbbabbababaaababbbbbabaabbaabba
aabbabbbabababaabaaabbbb
bbbabbbababbbbaabbbbbabb
abababaababaabbbbaabaabb
aaababbababbabaabbaabaaaabababbbaaababbaababaabaaaabbbbabbabaaaa
baabbababbbabbbabbbabaab
aaaaaabbbaabbaaaaabbbbbaaababbbabaaaaaaa
bbaababbabaaaabbaaabbaaa
bbaabbbbbaabaababaabaabbababbbba
abbabbaabbaaaabbaaababbabbabaabaaabbbbbabaabaaaa
baaaabaabaaabbaaaabaabba
aaaabbbbaababbbabbabaaaabbbabaab
aaabbbbbbbbbbaaaabbabbbabaaabbbb
babbbbaabaabaaaabbabbababababbaaababbabbbbbabaaabaaaabaaaaabaaaabbabaaaa
bbabaaaaababbaababbbbbab
babababaaaabbbabaaaaabbb
abbbbabbababbaabaaaabbabaabaabababbbabba
bbabbbbbbaaabaaaababbaabbaababab
baaaababbbaaababbaaabbbb
bbabbaabaaabbabababbbaaa
abbabbaaaabbabbaababbbba
bbbabbbaabbaaaabbabbbbab
bbaabbbabaaaabbaabbaabbbaabaaaababbaababbbbaabababbbabbabbbbbaba
aabbabbabbabbababababbbbabbaabba
aabbbaaabbabbaaaabbbbbab
baaaabaaabaaaabababaaabb
bbaabbbaabbaaabababaaaaa
baaaabaabbabbaaababbbaaa
bbaabbaabaabbabaaababbab
babaabbaabbaaaaaaaaaaaaa
bbaaabaabbabbaaaaabbaabb
aaabaaababbbabbbaaaababababbaababaabbbba
abbaaaabbbaabbbaaabbabaa
ababaaabaaaababababbbaaaaaaabbaa
babaabbaaaaaabbaaaaaaabbbbaabbaabaabbabbbbbaabab
aababbbbaaaabbababaaabbbbaabbaaabbbbbaaaaabbbaabbabaabaa
abaabbbbbaabbaaaaababbbaabaaaaaaababaaaa
babbabaabbaababbabbaaaabbaaabbbb
baaaabaaabaabbbaaaaabbabbbabbaababbbaaaa
abbbbababaabbaababaaaaab
babaaaaababbbabbabaabbab
bbaaaabbaabbaaabbbaaaaba
aabbaaaabaabbaaabaabbaabbaababab
bbaabbaabaabbabbaaabbbabbbbbabbababaabab
bbabaabaaaabbabaabaaaaab
bbabababbabaabbabbbabaaaabbbaaaaabaaabaabbababba
abbbbabaabaaabbbbabaaaba
ababbaaaaabbabbabbbbaaab
aaabbbabbbaabaabababaabaabbaaaaaaaaaaabbaaabbbbbbbababba
aaabbaaaabbbaababababababababbbbbabbbbabbbaabbbbaabbaaab
bbaabaaaababbbaabbbabbaa
abbaabaababbbababaabababababbbaa
babaabbbaaabbbabaaabbbbbaababbbabaabbbbaabaababb
abbabbbaaaaaabbababbbbaabbabaaabbababbaaaaabbabbababbbba
aaababbaababaababaaabbab
ababaabbbbbbbabbaabbbbab
aababbbbbbbbbaaaaaabaaba
abaaaaaabbbbaaaababaaabb
abbbaabaaabbaaabaaababbbbabaababbbaaaaab
bbbbabaabaabbaababbbabbb
bababbaaabbaababbabaaaaa
ababbababaaabbaababbabbabbaabaabbbbbabaaabababaaaaaabababbbbaaba
baabbabbbbabbaabbbbabbaa
abababaabbaaababaabaabaaabaabababbabbbba
abaaaaaaaaaabbabbababaaa
aaabaabbaaaaaabbbbbbabba
baaabaaabaababbbaaabbaaa
bbbbbabbbaabbabaaabbabaa
bbababbbbaabaaabaaabaaba
aabbaaabaaaabbabbbabaaabbaaabbba
babaabbaaabbaaabbabbbaba
abbabbabbaaaabbabbbbbaba
aaabaabaabbbababbbababab
ababbaabababbaabbbabbabb
aabaabaaababaababbaabaababaaabbaababbbbbbbbbbaab
bbabaababbaaaaaaaaabbabababaaaaababbbaaa
baabaaabaabaabaabaababbaaaaababb
abaaaaaaaaabbbabaaaaaaab
aaaabaabbbbbaabbabbbaababaaaaaabaababbbaaabbbaabaaaabbba
bbabaababbaaababbababbbb
abaaabbbbaaabaaababaaaaa
abbbaaabbbbabaaabaaaaababaaaaaababbbabbbbabbbbab
bbbabaaaabbbbbbbababbabababbaaba
babbbbbababbbbbbaabbbabb
bbaabaabbbabaababaabbbba
bbabbaabbabaaaababbababb
babababaabbaaababaaabbbb
ababbbaababbaaaaababbabb
aaaabbababbaaababaababab
baabaababbbabaaabaabbbaaabbababbbbbbabba
bbbaabbbbbababbbaaaaabbaaaaababa
abbbbaabbbaaaaaababaaaaa
bbaaababaabbbaaaaaaabaaa
bbababbbaababbbbbabbbabb
bbabbaaabaaaaababbaaabaaaabbaaaaaaababaa
bbabbaaaabbabbbaabaaaabbaabbabaa
baabbabbaababbbbbabbbaab
ababbaabaabaaaababbababb
babaabbbbbbaabbbbbbbaabbbbabbbaabaabbbaaabbbbbbaabbbbbab
abbaaaabbaaaaabaaaaabbababbabbaaababaabbbbbbbbabaaabbbaa
aabbbbbaababbababbaaaaba
abbababaaabababaabbaaabb
bbbaaababbaabbaabbaabbabaaababbbbaabbbbbabbbbaaaaabaaaaa
abbabbaaaabbabbabaaabbbb
abaabbbbbaabaaababbaaabb
aabaabaaabababbbbaabaabb
aaababbababaabbaabbabbaaababbbbbabbbbaaa
babaabbbaabbbaaaabbabaab
baaabbaabaaaabbaabbbabaa
bbabbbbbbabaabbbbbbaabbababababbbaaabaabbababbbabaabbbbbbbababaa
baabbaaababaaaabbbaaabba
bbabbaaaaababbbabaababaa
bbbbbbabababbabaaabaabbbbaaaabbaababbaaaaaaababbbaaababa
abbbbaabaaaababbaaaabbbaabbbbbab
abbabbabaaaaabbaababaabaaaaaaabaabaabbabbaaababa
bbbababbabaabaaaababbbaaababbabaabbbabbaabbbaaaa
bbbaabaaabaaaabbbabaabaa
aaababbaabbbbabaabbaabba
baaababbabaaabbbabbabaaabaabaaaaabbbbbba
ababbaababbaaaaabbbbabab
bbababbbabbbaaabaabbabbbaabaabab
ababababbbaabbbabbaaaaba
babbabbaaaababbaaabbaaba
abaaaabaabbaaaaaaabbbabb
babbaaaababbabbabababbbb
bbabbbbbbbababbbabbbaabb
baabbabaaabbbbabbbaaaabaabaaaaabbbabbbbbaabaaaabbaaaabbaabbaaaaabbbbbabb
abaaabbbabaaaabbbbabbbaa
bbbabbbabbaababaabbbabab
abaabaabaabbbaabababbbab
babbaaaabbbbaabbbabbbababbbabbabbaaaaabb
baabbaaaaabbabbbbbbbabaaaaababbaaababaaa
abbbbabbbaaaabaabbbbbbaa
bbbbaaaaaaaaababaabaaaaa
baaababbabbbaabaababbbba
bbaabbbaabaabaabbbbaaaab
abaabbbababaabbbabbbabba
baaabbaaabbbbabbbbabbbba
aabaaababaaabaaaaaaababb
baaabaababaabbbbabbaabbbbababbabbbaaabaabbbbabbb
aaaabaabbaabbaabbbaaabbb
aaabbbababbbbabaaabbbbbb
bababbaabbaabaabbbbbabaabbabbbbbaaabbaabbabaaaaa
bbbabaaababaabbabbababba
baabbaabaabbaaaabbbaaaab
abbaaaabbbbbaabbbaaaabaabababaaa
baaaaabababababbaabaaabb
abbbabaabbbbabbbbabababbaabaabbabbaaabbbbaabaababaabbabb
aabbaaabbbbababaabaaaabbabbaabbbaaaaababbaabbbba
baabbaabaaaaaabbbabbabbababbaaaaabbbabbb
bbbababbabaabaabaabaaaabbbbaabaababbbabaaaaabbbabbbabaabbbbaaaab
abaabaabbaabbababbbbbaba
aaabbbbbbbbbbabbababaabaabbbabba
bababbaabbbbbabbbbaababbbabbbabaabaabababbabaaabbbabbaba
abaabaaabbbbaaaababbbababbbbaaab
bbaabbabbaabbaabbabaaabb
abaaaababbbaabaaabaaaabbbabbbbabbbbbbaab
aabbaaabbbaabbbbaabbabab
baaababbbbaababbbaaaaaaa
ababbaaabbabaaabaaaababa
bbabaabbaaabbbbbabbbbbbbabaabaaaaabaaaaa
baabbababaaaababaabaaabaaababbbababbbbbabbabbbab
abaaabbaabbaaaaabbbababbaababbab
baaaaaabbbbaaaaabbbbbbaaabbabbba
ababbaaabbaabbabaabbabaa
aabaaababbbaaabaabaabaaaabbaababaababbabaabbabaa
abaaaababbabaaaaaababbbbbbababaa
babaabbabbbbabaababbabab
abaaaaaaabbbbbaababababbbaaabaabbbbbaaaaaaababaaabaaaaab
aaaabbabbaabbaabaaabbbbbbabbababababbbba
bbbbabaabbaababbbbbbbaba
abbabaaabaaababbbbbbabab
baabbabbabaaaabbbbbbbaab
bbaabaaaaabbabbabababbbb
bbabaaaaaaaabaabbbbbbbaa
abbbbabbbbaaaaaabababbba
bbbbbbabaaaabaabbbbbaaaaaaabbaab
babbabbaabbbabbababaabababbbbbbaabaababaaaabbbaaabbbaaba
abbbbbaabaaababbaaababbb
bbbaababaabbabbababaabaaaaaabaabababbbbaabbbbbababababbbbbbaaababaabaabb
baabaaabbababbababbbbbbabbbbaaabababbbab
abbabababababbababababba
abbabaaaaaabbbbbaababbbbaabaaaababaaaababbabbbba
abaaaaaaaaababbababaabaabbbaabbbaaababbbbbbaaaaabaaabbaa
bbbbbaaabaababbbbabbbbab
bbaaababbabaabbbbabababbabaabaaababbbbbababbaabb
aaaabaabbababbaabbabbbaa
aaaabbbbabbbaaababbabaaaabbbbabaabbbbabbababbbabbabbbaabaabbababbabaabaa
babbabbaaabaaaababbababb
aaabaabbabbbbabbbbbaabab
ababbbaaabbabaaabababaab
abbbbababaaaaabaaaababbb
ababbbbbbbabaabbaaabaaaa
abaabbbaabbbaabaaabaaaaa
bababbaaaaabaabbabaabbab
bbabbaaaabaababaaaaababa
babaabbaabababbbaabbbaab
babbbabaaabbabbbbbaaabbb
baaababbbaabaababbabbaaaabaaaaaabaabbbbaabbababb
babbbbbbbbabbababbbbaaaabaabaaaaabababaabbbbbbbaaaaaaaaa
baaabbaabaabbabababbaaab
bbbababbabaabbbbbabbbbab
aabbaaababbaaaabbabbaabb
bbaaaaaabbbaabaaabbbbabbbbbababbaaabbbbbbbbabaaaabbbabab
abbababababababaabaaaaaaaaababab
baabbabaaaaaababbabaaaaa
aabababaabaabaaaabbbbbba
abaaabbaabbabababababbbb
abbabbaaabaaaaaabbbababbaababababbbaabaabababbbabbababbaaabbbaba
babbbbaaabaaaaaabababbabaaaabbbaabaabbbaaabaabbb
baaaaaabbaaabbaaaaabbbbbbbbbbaab
abaabbbaaabababaabbbbbbbabbaababbbabbaabaabbabbbbbababaabbbaabab
bbaaaabbaaabaabbbbabaaaaabaaabbbbbabaaaaaabbbbaa
abaabaabbababababbbbabba
babababbbbaabbaababbbbbbbabaabbababbbbbbbaaaaababababbba
ababbbbabbabbbaaabbaaaaaabbbabaabbbbbabb
baabaaababaaaaaaaabababb
aabbbaaaababbbaabbabbaaabbbaaaaaababbbab
abbabbbabbaabaaaaaaabaaa
aaaaaabaaaabbbabaabbbaab
aaaaababbaaabaaabaababbabaaaabbbbaaababababaababaabbaabb
babbbababaaaabbbababbbaaabbbaaabaabbaaaaaaaabbabaaabbaab
bababbabaabbabbaaabaabab
babbbababbabaabababbabab
bbabbaabaabaabaaabbabbbb
abbbaaabbbaabaababbabbbaabaabbaa
aaaaaabbbbbabaabbabbbbabbbbbbbbbabbabbbbbabbabab
aababbbaabaabaabbbbbaabbaabaabbbbbbabbbaababaaaa
baaabaaababbbbbbaabaabbbaabbabbbabbaaaaabbaaabbaaabbaaba
baabbababaababbbabababaabaaabbbaaabbbbaa
bbabaabbbaaaaaabbbbabbbabbabbbbbabaaaabbbaabbbbaabbabaab
baaaaaababaaaaaaaaaababa
aabaaabababbbbbaaaaaaaaa
abababaaabaaabbaaaaabaaa
baabbaababbaaaababaaaabaabababbaabaababbaabaaabb
babbbababbabaaabaabbbbbabbabaaaabaaabbbb
abaabababbbabbbaaabbbbab
ababaababbababbbabbaaaaaaababbbaaaaabbaaabaababb
aaabbbabbbaabaabbbbaaaab
baaabbaaabaaaabaaabbbbab
bbbaabaabbabaaabaababbbbbabbaaaabbaaabbb
bbabaaababaabaaaabbbaaabababaabbabbbabbbabbaabba
ababaabbabbaaaaaabbbbbaabbabaabbbbbaabaabbbabaabaaaababb
babbbabaabbaaababbaaabba
baaaabbaaabbabbbbabbbbabbbabbaba
bbbbbaaaaaaaababbaaaabbaabaaabab
bbaabbaaabbbbabbbbbbaaab
ababaabaabaaabbbabbbabab
bbaabbaabaababbababbbaaa
bababbaaaabbaaabbaaaababaaaaababbbabbabbbababbbb
baaababbbaaaabbaaaabaaba
bbbabbbabbaaaabbbbbaabbaabbbaaabbabbabbb
bbaabbbbbbaaabaaaabbaaaababbabab
abaabbbaabaaaaaaaaaabaaa
aababbbbaababababbbabaab
aabbaaabaaabbbababbabbabbabbaaaaaaabbbaa
baabbabbabaaabbaabaabbaa
bbaababaaabbbbbaabbabbababaabbbbababbabbabbbbaaa
babababaaaaaababaabbbaba
bbababbbbbbbbabbabaabbbbbbababbabbbabbbb
aabababbaabaaababbaabbbbbbaaaaaaabbbbaaaababbabbaaaaaaaaabbabababbbbbbbbaaaaaaaaaaaaaaba
abbbbbbaaaabbbbababaaaba
baabaaabbaaaaaababbbbbba
bbaabbaaabbbbbaababababbbaabbbba
babbbbaababaaabbaaababaabbbabbaabaaabbabbabaaaaabbabbbba
babaaaabbabaaaabbaaaabbaaabbbbab
ababbababaaaabbaaabababb
bbbbbbababababaabaabbbba
bbaababbabbaaababababbba
aaababbabbaaabaaaabbabbaaaaabaaabababaaa
abaabaaaabaabababbbbaaba
ababababaaababbababaaaaa
bbabaaabaababaaabbabaaabbababbab
baaabbaaaabaaabaaaabaabbbaaaaabaaaaaabaababaababbaaababaabaabbaa
babbbbbaaabbabbbbaaaabbaaaabbbbbaaabaaba
aababbbaaaabbbaaababbaaabbaaaababbabaabbbaaababababaaaabaabbaaaa
bbbaabbbbbaaaabbbbaaabaaaabbabbabababaaa
abbaababababbaaabaaaabababbbaababbbaabbbabbabaab
aabbbaaaaabbaaabbbbbaaab
baaaababaabbabbbaaababab
abaabbbababbbbbbbbaaabba
babbbbaababababbaaababbb
baaababbaabbbababbaaabbbbbaaaaabbaaaaabbabbaaaaa
bababbaabbaababaaabbaabbabbaaabbbabbaaabbabbbabbaaabbaaa
bbaabbabbabbabaabbaabababaabbabbaabbaaaaaababaaabababbbb
baabbaabbbaabbaabbbbbaab
babaaaabbbbbabaabbaaabba
bbbabababbbaaabaaaaaababaaabbabb
bbababbbaabaaaababbbbaabbaaaaababbababbaaabaaabbaabbbbbb
aaaaabbaabaaabbbbabbaabb
abbbbaabaababbbbabaabaabbbbbbbbb
abaaaabaaaabbbabaaabbbbbbbbbaabababbaabb
bbbbaabbaabaaaababbabbababbabbabaababaaa
abbbaaababbaaabaabbbabba
bbaabbaaabababaaabbbabba
aabababaabbbbbaaaaaaaababbbbbaabaababbab
bbabbaababbaababaabbbbab
bbabaaaaaaaaabbababbbbbbbbaaaabbaabaabbbbbbbbabaabbbabbbbababaab
aaaaaabbaaababbaaaabbaab
bbaababbabaaaababaabaaabaaaaabaa
aaababbabbbabaaaabbbaabb
babbbbbbbbabbbbbaaaabaabbabaaaaabbbbabbb
baaabaabaaabbbabababbbab
aabaaabaaabaabaaabbbaaaa
abaabbbabbabaababbbbbaaaababbbabaabbaabbbbbabaab
aabbabbbababbabaabbbabbb
bbbaaabaabababbbbbbbabab
bbbbaabbbbbbaaaaaaaabbabbbabbbbabbbaaaab
abaabaababaaabbabbabbaba
bbbabbbaaaaabaabaaaabbbbaababaabababaaab
aaabaaabaabaaaabbbaaaaabbabbabaabbabababbaabbbabbaaaabbbbabbbbaaabaabbaababbbabbaaaabbbb
bbbababaabbbbabaababaaab
babaabbbabaabbbaabaaaabbabbbbbba
bbaaaabbaabaaaababababba
abbaaaaababaabaababbbaaabaabbbabaababaaa
aabbabbbabbaaababaaaaaabaabaababaabbbabb
bbaabbaabbabaaabababbabb
baabbaaaaabbabbbaabbaaba
aaaabbbabbaabbbbbaaaaabaabaaaabbaabababaabbbbaaaabbaabbababbaabaaabaaaab
bbbabbbbaababbabbbbbbbaabbaaaabaaaabbbababbbaabbbaabaabaabaaaababbbbbabbaabbabbb
aaabbababbaaaabbbaaaabbabbaabaabbaababab
bbaabbbaababaabbbabbaaba
abbbaababbabbaababbbbbaaabababbbbbbaaaaabbbbbbbbbbbbaaba
ababababbaaaaaabbbbbaaba
bbbbbaaabaaaaaabaaababbaaabaaaaa
aaaabbababaababaaaabbbabababbaabbbaababbabbbbbbaaaaaaaabbbbabbbbbbbbabab
aabaabbbbabbbabaaabaabab
aabaabbbababaabbbbaabaabbabbabaaabbbaaaabaabaabb
baaaabbabbbabbbabababaaa
baabaaabbaaaaababbbabababbaaabbb
babababababbaaaabaaabaabbbbbbbbb
abbbbaabbaaaababbaaaabbbaababababbbbbbbabababaabaabaabba
bbbabaaaaabbbbbaabbaabba
abababbbbaaababbaaabbabb
abaabaabbabaaaabababbbaaaaababaabbbaaabb
aaaababbbaabbaabaaabbaabbbbabbbbaaabbaaabbbaaaab
aaababbabbbabbbababababaabbaabaa
bbaaaabbaaaaaabaaabababb
abbbbaaabbbbbabbbaabaabbaaababaaabaabaaabbaabaaaabbbbbbbbbbbbbabababaaab
bbbbaabbaabaabbbaabaaaababaaabaaaaabaaba
abbabaaabaaaababbbbbbbababbbbbbbabbbaaaa
bbaabbbbababaabbaaababab
aabaabaaaaaabaabbbababaa
aabaaabaabbbaaabbbbababaababbaab
babbbbbaabababaaababaabbaabbbaab
bbabbababbaabbbbbbaaabaaaaababbaabbbbbbbaabbbbab
bbaababababbbbbabbabbaaababbbabb
aabababaaababbbabaaaaabb
babbaababaabaaabababbaababababaababbbbbb
bbbaabbbbbbaabbaabbbaabb
bbabbaaabbaababbaabbabbbbaabaabbaabbbaba
abbbbbbbabababaabaaabbab
abaaabababbbbbbaabaaabbabaaaaababbbbaaabaaaabaaabbaaabab
abbaabbbabbbbabaaabaaabb
bbababbbabbabbbabbbabbbbabbabbabbabaaabaaabaabbbbbbaaaaa
baaaabaaabaaabbaabbbabbb
aababbaaabbbabaaabaaaaabbbbbabbb
aaabbbbbaabbabbaaabbaaaaaaabbababbaabaabbbbaababaabbbabbbbabababbabbbabb
bbbbbababbabaababbaababaabaababaaaababbbaaabbbaabaabaababbababab
aaababaaaabaaabbaaabaaba
baababbbaaaabbbbbbbbabab
abababbbbbabaaaaaabbbaaaabbabaaaababbbab
ababababbbbbbaaabbaabbbaaabbbbbaababbbbbaaaabaaabaaabbbbbbaaaaabbabaaaaa
aaaaaaaaabaaabaaabbabaabbaabbaaaabbbabbabbababbbbbababaabbbbabab
baaabaaaabaaaabbbbaaaaab
babbabbaaaabbabaabbbaaaa
abbbaababbbababaaababbbbbabababbbbaabbabbbbbbbaabbbbbaba`

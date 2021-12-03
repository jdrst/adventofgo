package main

import (
	"testing"
)

func TestNumberN(t *testing.T) {
	newLine = "\n"
	ruleStrings, messages := prepInput([]byte(testinput1))
	rules := makeRules(ruleStrings)
	rule := getAllRulesFor("0", rules)
	actual := countMessages(messages, rule)
	expected := 2
	if actual != expected {
		t.Errorf("expected %d messages to match, actual %d.",
			expected, actual)
	}
}

// func BenchmarkNumberN(b *testing.B) {
// 	newLine = "\n"
// 	for i := 0; i < b.N; i++ {
// 		// for _, test := range testcases {
// 		// 	lastSpoken, _ := strconv.Atoi(strings.Split(test.input, ",")[2])
// 		// 	input := prepInput([]byte(test.input))
// 		// 	numberN(input, lastSpoken, test.rounds)
// 		// }
// 	}
// }

var testinput1 = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

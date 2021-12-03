package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

var newLine = "\r\n"

func prepInput(input []byte) ([]string, []string) {
	parts := strings.Split(strings.TrimSpace(string(input)), newLine+newLine)

	rules := strings.Split(parts[0], newLine)
	messages := strings.Split(parts[1], newLine)
	return rules, messages
}

func makeRules(ruleStrings []string) map[string][]string {
	rules := make(map[string][]string)

	for _, rulestr := range ruleStrings {
		rule := strings.Split(rulestr, ":")
		//num, _ := strconv.Atoi(rule[0])
		rules[rule[0]] = strings.Split(rule[1], "|")
	}
	return rules
}

func countMessages(messages []string, rule []string) int {
	count := 0
messages:
	for _, m := range messages {
		for _, r := range rule {
			if m == r {
				count++
				continue messages
			}
		}
	}
	return count
}

func main() {
	ruleStrings, messages := prepInput(parseInput())

	rules := makeRules(ruleStrings)
	// fullyReplaced := make(map[string][]string)
	// for len(rules) > 0 {
	// 	for key, subrules := range rules {
	// 		done := true
	// 		for _, sr := range subrules {
	// 			if strings.ContainsAny(sr, "0123456789") {
	// 				done = false
	// 			}
	// 		}
	// 		if done {
	// 			for j, sr := range subrules {
	// 				subrules[j] = strings.ReplaceAll(strings.ReplaceAll(sr, "\"", ""), " ", "")
	// 			}
	// 			fullyReplaced[key] = subrules
	// 			delete(rules, key)
	// 		}
	// 	}

	// 	if _, exists := fullyReplaced["0"]; exists {
	// 		break
	// 	}

	// 	for key, subrules := range rules {
	// 		newSr := []string{}
	// 		for _, sr := range subrules {
	// 			fields := strings.Fields(sr)
	// 			amount := 1
	// 			fooKey := ""
	// 			found := false
	// 			foo := []string{}
	// 			replaceMap := make(map[string][]string)
	// 			for _, field := range fields {
	// 				if rep, exists := fullyReplaced[field]; exists {
	// 					amount *= len(rep)
	// 					replaceMap[field] = rep
	// 					foo = rep
	// 					fooKey = field
	// 					found = true
	// 					break
	// 				}
	// 			}
	// 			replacedRules := make([]string, amount)
	// 			for i := 0; i < len(replacedRules); i++ {
	// 				if found {
	// 					replacedRules[i] = strings.ReplaceAll(sr, fooKey, foo[i])
	// 				} else {
	// 					replacedRules[i] = sr
	// 				}
	// 			}
	// 			//replace dem rules

	// 			// for i := 0; i < len(replacedRules); i++ {
	// 			// 	replacedRules[i] = sr
	// 			// 	for key, rep := range replaceMap {
	// 			// 		for j := 0; j < len(rep); j++ {
	// 			// 			replacedRules[i] = strings.ReplaceAll(sr, key, rep[j])
	// 			// 		}
	// 			// 	}
	// 			// }
	// 			newSr = append(newSr, replacedRules...)
	// 		}
	// 		rules[key] = newSr
	// 	}
	// }

	rule0 := getAllRulesFor("0", rules)

	count := countMessages(messages, rule0)

	fmt.Println(count)
	// changes := 1
	// for changes > 0 {
	// 	changes = 0
	// 	for key, rule := range rules {
	// 		for _, subrule := range rule {
	// 			rules[key] = replaceRules(rules, subrule)
	// 		}
	// 	}
	// }

	// for _, subrules := range rules {
	// 	for i, rule := range subrules {
	// 		subrules[i] = strings.ReplaceAll(strings.ReplaceAll(rule, "\"", ""), " ", "")
	// 	}
	// }
}

func getAllRulesFor(rule string, rules map[string][]string) []string {
	res := rules[rule]
	changes := true
	// loop:
	for changes {
		changes = false
	loop:
		for j, sr := range res {
			if strings.ContainsAny(sr, "0123456789") {
				changes = true
				fields := strings.Fields(sr)
				var key string
				var replace []string
				for _, field := range fields {
					if rep, exists := rules[field]; exists {
						replace = rep
						key = field
						break
					}
				}
				length := len(replace)
				if length > 1 {
					additional := make([]string, length-1)
					m := regexp.MustCompile("\\b" + key + "\\b")
					for k, l := 0, 1; l < len(replace); k, l = k+1, l+1 {
						additional[k] = m.ReplaceAllString(sr, replace[l])
					}
					res = append(res, additional...)
				}
				m := regexp.MustCompile("\\b" + key + "\\b")
				res[j] = m.ReplaceAllString(sr, replace[0])
				continue loop
			}
		}
	}
	for j, sr := range res {
		res[j] = strings.ReplaceAll(strings.ReplaceAll(sr, "\"", ""), " ", "")
	}
	return res
}

func contains(array []string, s string) bool {
	for _, a := range array {
		if a == s {
			return true
		}
	}
	return false
}

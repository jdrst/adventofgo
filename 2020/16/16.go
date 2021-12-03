package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type rule struct {
	firstBounds, secondBounds bounds
	possiblePositions         map[int]struct{}
}
type bounds struct {
	lower, upper int
}

var newLine = "\r\n"

func main() {
	lines := prepInput(parseInput())

	possiblePositions := make(map[int]struct{})
	for i := 0; i < len(strings.Split(lines[27], ",")); i++ {
		possiblePositions[i] = struct{}{}
	}
	rules := make(map[string]rule)
	for _, line := range lines[:20] {
		words := strings.Fields(line)
		var firstBounds, secondBounds bounds
		fmt.Sscanf(words[len(words)-3], "%d-%d", &firstBounds.lower, &firstBounds.upper)
		fmt.Sscanf(words[len(words)-1], "%d-%d", &secondBounds.lower, &secondBounds.upper)
		currentPossiblePositions := copy(possiblePositions)
		name := words[0]
		if len(words) > 4 {
			name = words[0] + words[1]
		}
		currentRule := rule{firstBounds, secondBounds, currentPossiblePositions}
		rules[name] = currentRule
	}

	errorRate := 0
tickets:
	for _, line := range lines[25:] {
		values := strings.Split(line, ",")
	valid:
		for _, valStr := range values {
			value, _ := strconv.Atoi(valStr)
			for _, currentRule := range rules {
				if currentRule.isAllowed(value) {
					continue valid
				}
			}
			errorRate += value
			continue tickets
		}
		//only valid tickets here
		for i, valStr := range values {
			for _, currentRule := range rules {
				value, _ := strconv.Atoi(valStr)
				if !currentRule.isAllowed(value) {
					delete(currentRule.possiblePositions, i)
				}
			}
		}
	}

	fmt.Println(errorRate)

	rulesAndPositions := findExactRulePositions(rules)

	result := 1
	for name, rule := range rulesAndPositions {
		myTicket := strings.Split(lines[22], ",")
		if strings.HasPrefix(name, "departure") {
			var i int
			for i = range rule.possiblePositions {
				break
			}
			num, _ := strconv.Atoi(myTicket[i])
			result = result * num
		}
	}
	fmt.Println(result)
}

func copy(in map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})
	for i := range in {
		result[i] = in[i]
	}
	return result
}

func (r *rule) isAllowed(value int) bool {
	return r.firstBounds.isWithin(value) || r.secondBounds.isWithin(value)
}

func (b *bounds) isWithin(value int) bool {
	return value >= b.lower && value <= b.upper
}

func findExactRulePositions(rules map[string]rule) map[string]rule {
	rulesAndPositions := make(map[string]rule)
	for len(rules) > 0 {
	rules:
		for name, currentRule := range rules {
			for i := range currentRule.possiblePositions {
				count := 0
				for _, secondRule := range rules {
					if _, exists := secondRule.possiblePositions[i]; exists {
						count++
					}

				}
				if count == 1 {
					currentRule.possiblePositions = map[int]struct{}{i: struct{}{}}
					rulesAndPositions[name] = currentRule
					delete(rules, name)
					continue rules
				}
			}
		}
	}
	return rulesAndPositions
}

func parseInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return input
}

func prepInput(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), newLine)

	return lines
}

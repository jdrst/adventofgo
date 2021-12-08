package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	sublines := make([][]util.Line, len(lines))

	sum := 0

	for i, l := range lines {
		sublines[i] = l.SubSplitWith(" | ")
		digits := strings.Fields(string(sublines[i][1]))
		for _, d := range digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				sum++
			}
		}
	}
	return sum
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	sum := 0

	for _, l := range lines {
		split := l.SubSplitWith(" | ")
		sum += deduceValue(string(split[0]), string(split[1]))
	}
	return sum
}

func deduceValue(pattern, output string) int {
	patterns := strings.Fields(pattern)

	mapping := getPatternMapping(patterns)

	outputs := strings.Fields(output)

	mappedOutputs := mapOutputs(outputs, mapping)

	return getValue(mappedOutputs)
}

func getPatternMapping(patterns []string) map[int]string {
	patternMap := make(map[int]string)
	known := [10]int{}

	for _, p := range patterns {
		if len(p) == 2 {
			known[1] = sumOfCharacters(p)
		}
		if len(p) == 4 {
			known[4] = sumOfCharacters(p)
		}
		if len(p) == 3 {
			known[7] = sumOfCharacters(p)
		}
		if len(p) == 7 {
			known[8] = sumOfCharacters(p)
		}
	}

	for _, r := range "abcdefg" {
		count := 0
		for _, s := range patterns {
			if strings.ContainsRune(s, r) {
				count++
			}
		}
		if count == 4 {
			patternMap[int(r)] = "e"
		}
		if count == 9 {
			patternMap[int(r)] = "f"
			patternMap[known[1]-int(r)] = "c"
		}
		if count == 6 {
			patternMap[int(r)] = "b"
		}
	}

	patternMap[known[7]-known[1]] = "a"

	patternMap[known[8]-known[4]-keyForVal("e", patternMap)-keyForVal("a", patternMap)] = "g"

	patternMap[known[8]-sumOfKeys(patternMap)] = "d"

	return patternMap
}

func keyForVal(value string, m map[int]string) int {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return 0
}

func sumOfKeys(known map[int]string) int {
	sum := 0
	for k := range known {
		sum += k
	}
	return sum
}

func sumOfCharacters(str string) int {
	sum := 0
	for _, r := range str {
		sum += int(r)
	}
	return sum
}

func mapOutputs(outputs []string, mapping map[int]string) []string {
	mappedOutputs := make([]string, len(outputs))
	for i, s := range outputs {
		mappedOutputs[i] = getDigit(s, mapping)
	}
	return mappedOutputs
}

func getDigit(s string, mapping map[int]string) string {
	result := ""
	for _, c := range s {
		result += mapping[int(c)]
	}
	return result
}

func getValue(segments []string) int {
	var displayDigits map[string]string = map[string]string{"abcefg": "0", "cf": "1", "acdeg": "2", "acdfg": "3", "bcdf": "4", "abdfg": "5", "abdefg": "6", "acf": "7", "abcdefg": "8", "abcdfg": "9"}
	stringResult := ""
next:
	for _, s := range segments {
		for k, v := range displayDigits {
			if isInAnyOrder(s, k) {
				stringResult += v
				continue next
			}
		}
	}
	res, err := strconv.Atoi(stringResult)
	if err != nil {
		log.Fatal("getValue can't convert")
	}
	return res
}

func isInAnyOrder(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	result := true
next:
	for _, cA := range a {
		for _, cB := range b {
			if cA == cB {
				continue next
			}
		}
		result = false
	}
	return result
}

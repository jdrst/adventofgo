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
	one := 0
	seven := 0
	four := 0
	eight := 0

	/*
		we're producing a map from "mixed up" segment to correct segment,
		so we can then map each character (rune) of the pattern to the correct character,
		applying the following rules:
	*/

	//getting the values (sum of all characters) for known representations 1 ("cf"), 4 ("bcdf"), 7 ("acf") and 8 ("abcdefg")
	for _, p := range patterns {
		switch len(p) {
		case 2:
			one = sumOfCharacters(p)
		case 3:
			seven = sumOfCharacters(p)
		case 4:
			four = sumOfCharacters(p)
		case 7:
			eight = sumOfCharacters(p)
		}
	}

	//getting the values that can be derived from the count of occurence
	for _, r := range "abcdefg" {
		count := 0
		for _, s := range patterns {
			if strings.ContainsRune(s, r) {
				count++
			}
		}
		switch count {
		case 4:
			//four occurences gives us "e"
			patternMap[int(r)] = "e"
		case 9:
			//nine occurences gives us "f"
			//and also "c" because 1 is "cf" and "cf"-"f" = "c"
			patternMap[int(r)] = "f"
			patternMap[one-int(r)] = "c"
		case 6:
			//six occurences gives us "b"
			patternMap[int(r)] = "b"
		}
	}

	//difference between 7 ("acf") and 1 ("cf") is "a"
	patternMap[seven-one] = "a"

	//difference between 8 ("abcdefg") and 4 ("bcdf") = "aeg" then "aeg" - "a" - "e" (values we now know) is "g"
	patternMap[eight-four-keyForVal("e", patternMap)-keyForVal("a", patternMap)] = "g"

	//difference between 8 ("abcdefg") and the values of all characters known ("abcefg") is "d"
	patternMap[eight-sumOfKeys(patternMap)] = "d"

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

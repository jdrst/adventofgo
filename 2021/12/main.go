package main

import (
	"fmt"
	"unicode"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type cave struct {
	leadsTo []string
}

func partOne(file util.File) int {
	lines := file.AsLines()
	caveMap := map[string]cave{}
	visits := map[string]int{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		caveMap[string(caves[0])] = cave{leadsTo: append(caveMap[string(caves[0])].leadsTo, string(caves[1]))}
		caveMap[string(caves[1])] = cave{leadsTo: append(caveMap[string(caves[1])].leadsTo, string(caves[0]))}

	}
	return allPossiblePathsVisitOnce(caveMap, visits, "start")
}

func allPossiblePathsVisitOnce(caveMap map[string]cave, visits map[string]int, currentCave string) int {
	visits[currentCave] += 1
	sum := 0
	for _, c := range caveMap[currentCave].leadsTo {
		if c == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || visits[c] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsVisitOnce(caveMap, nextVisits, c)
		}
	}
	return sum
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	caveMap := map[string]cave{}
	visits := map[string]int{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		caveMap[string(caves[0])] = cave{leadsTo: append(caveMap[string(caves[0])].leadsTo, string(caves[1]))}
		caveMap[string(caves[1])] = cave{leadsTo: append(caveMap[string(caves[1])].leadsTo, string(caves[0]))}

	}
	return allPossiblePathsOneDoubleVisit(caveMap, "start", visits, false)
}

func allPossiblePathsOneDoubleVisit(caveMap map[string]cave, currentCave string, visits map[string]int, hasVisitedTwice bool) int {
	if visits[currentCave] > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	visits[currentCave] += 1
	sum := 0
	for _, c := range caveMap[currentCave].leadsTo {
		if c == "end" {
			sum++
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || visits[c] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(c[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsOneDoubleVisit(caveMap, c, nextVisits, hasVisitedTwice)
		}
	}
	return sum
}

func copyMap(visits map[string]int) map[string]int {
	res := make(map[string]int, len(visits))
	for k, v := range visits {
		res[k] = v
	}
	return res
}

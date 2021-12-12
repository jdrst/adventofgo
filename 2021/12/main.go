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
	visits  int
	leadsTo []string
}

func partOne(file util.File) int {
	lines := file.AsLines()
	caveMap := map[string]cave{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		caveMap[string(caves[0])] = cave{visits: 0, leadsTo: append(caveMap[string(caves[0])].leadsTo, string(caves[1]))}
		caveMap[string(caves[1])] = cave{visits: 0, leadsTo: append(caveMap[string(caves[1])].leadsTo, string(caves[0]))}

	}
	return allPossiblePaths(caveMap, "start")
}

func allPossiblePaths(caveMap map[string]cave, currentCave string) int {
	caveMap[currentCave] = cave{visits: 1, leadsTo: caveMap[currentCave].leadsTo}
	sum := 0
	for _, c := range caveMap[currentCave].leadsTo {
		if c == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(c[0])) || caveMap[c].visits == 0 {
			sum += allPossiblePaths(copyMap(caveMap), c)
		}
	}
	return sum
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	caveMap := map[string]cave{}

	for _, l := range lines {
		caves := l.SubSplitWith("-")
		caveMap[string(caves[0])] = cave{visits: 0, leadsTo: append(caveMap[string(caves[0])].leadsTo, string(caves[1]))}
		caveMap[string(caves[1])] = cave{visits: 0, leadsTo: append(caveMap[string(caves[1])].leadsTo, string(caves[0]))}

	}
	return allPossiblePathsTwo(caveMap, "start", false)
}

func allPossiblePathsTwo(caveMap map[string]cave, currentCave string, hasVisitedTwice bool) int {
	if caveMap[currentCave].visits > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	caveMap[currentCave] = cave{visits: caveMap[currentCave].visits + 1, leadsTo: caveMap[currentCave].leadsTo}
	sum := 0
	for _, c := range caveMap[currentCave].leadsTo {
		if c == "end" {
			sum++
			continue
		}
		if c != "start" && (!unicode.IsLower(rune(c[0])) || caveMap[c].visits < 1 || !hasVisitedTwice) {
			sum += allPossiblePathsTwo(copyMap(caveMap), c, hasVisitedTwice)
		}
	}
	return sum
}

func copyMap(caveMap map[string]cave) map[string]cave {
	res := make(map[string]cave, len(caveMap))
	for k, v := range caveMap {
		res[k] = v
	}
	return res
}

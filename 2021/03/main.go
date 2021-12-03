package main

import (
	"fmt"
	"strconv"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()
	gammaRate := 0
	for i := 0; i < len(lines[0]); i++ {
		gammaRate <<= 1
		if getMostCommonBitOnPos(i, lines) {
			gammaRate |= 1
		}
	}
	return int(gammaRate * (1<<len(lines[0]) - 1 ^ gammaRate))
}

func partTwo(file util.File) int {
	oxy := file.AsLines()
	co2 := file.AsLines()

	for i := 0; i < len(oxy[0]); i++ {
		mostCommon := getMostCommonBitOnPos(i, oxy)
		leastCommon := !getMostCommonBitOnPos(i, co2)

		for j := 0; j < len(oxy) && len(oxy) > 1; j++ {
			if string(oxy[j][i]) == "1" == mostCommon {
				continue
			}
			oxy[j] = oxy[len(oxy)-1]
			oxy = oxy[:len(oxy)-1]
			j--
		}

		for j := 0; j < len(co2) && len(co2) > 1; j++ {
			if string(co2[j][i]) == "1" == leastCommon {
				continue
			}
			co2[j] = co2[len(co2)-1]
			co2 = co2[:len(co2)-1]
			j--
		}
	}

	first, err := strconv.ParseInt(string(oxy[0]), 2, 64)
	util.Handle(err)
	second, err := strconv.ParseInt(string(co2[0]), 2, 64)
	util.Handle(err)
	return int(first * second)
}

func getMostCommonBitOnPos(pos int, lines util.Lines) bool {
	cnt := 0
	for _, l := range lines {
		if l[pos] == '1' {
			cnt++
		}
	}
	return float64(cnt) >= float64(len(lines))/2
}

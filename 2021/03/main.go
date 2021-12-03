package main

import (
	"fmt"
	"math/bits"
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
		if getMostCommonBitOnPos(i, lines) == '1' {
			gammaRate |= 1
		}
	}
	leadingZeroes := bits.LeadingZeros(uint(gammaRate))
	return int(gammaRate * (^gammaRate << leadingZeroes >> leadingZeroes))
}

func partTwo(file util.File) int {
	oxy := file.AsLines()
	co2 := make(util.Lines, len(oxy))
	copy(co2, oxy)

	for i := 0; i < len(oxy[0]); i++ {
		mostCommonOxy := getMostCommonBitOnPos(i, oxy)
		mostCommonCo2 := getMostCommonBitOnPos(i, co2)

		for j := 0; j < len(oxy) && len(oxy) > 1; j++ {
			if rune(oxy[j][i]) == mostCommonOxy {
				continue
			}
			oxy[j] = oxy[len(oxy)-1]
			oxy = oxy[:len(oxy)-1]
			j--
		}

		for j := 0; j < len(co2) && len(co2) > 1; j++ {
			if rune(co2[j][i]) != mostCommonCo2 {
				continue
			}
			co2[j] = co2[len(co2)-1]
			co2 = co2[:len(co2)-1]
			j--
		}
	}

	oxyGenRat, err := strconv.ParseInt(string(oxy[0]), 2, 64)
	util.Handle(err)
	co2ScrRat, err := strconv.ParseInt(string(co2[0]), 2, 64)
	util.Handle(err)
	return int(oxyGenRat * co2ScrRat)
}

func getMostCommonBitOnPos(pos int, lines util.Lines) rune {
	cnt := 0
	for _, l := range lines {
		if l[pos] == '1' {
			cnt++
		}
	}
	if float64(cnt) >= float64(len(lines))/2 {
		return '1'
	}
	return '0'
}

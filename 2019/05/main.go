package main

import (
	"fmt"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	// fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) string {
	strings := strings.Split(string(file), ",")
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i] = util.ToInt(s)
	}
	return fmt.Sprint(execute(ints, 0))
}

func execute(intcode []int, currPos int) string {
	currCode := intcode[currPos]
	if currCode == 99 {
		return "exited"
	}
	p1mode, p2mode := 0, 0
	p1, p2, p3 := intcode[currPos+1], intcode[currPos+2], intcode[currPos+3]

	if currCode > 99 {
		p1mode, p2mode = currCode/100%10, currCode/1000%10
		currCode = currCode % 100
	}

	if p1mode == 0 && currCode != 3 {
		p1 = intcode[p1]
	}
	if p2mode == 0 && currCode < 3 {
		p2 = intcode[p2]
	}

	switch currCode {
	case 1:
		intcode[p3] = p1 + p2
		return execute(intcode, currPos+4)
	case 2:
		intcode[p3] = p1 * p2
		return execute(intcode, currPos+4)
	case 3:
		intcode[p1] = 1
		return execute(intcode, currPos+2)
	case 4:
		fmt.Println(p1)
		return execute(intcode, currPos+2)
	default:
		return "fault"
	}
}

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

const (
	sum int64 = iota
	product
	minimum
	maximum
	literal
	greaterThan
	lessThan
	equal
)

type operation struct {
	tId         int64
	lastIndex   int
	subPkgCount int64
	values      []int64
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	binary := toBinaryString(string(lines[0]))

	var sum int64 = 0
	i := 0

	for i < len(binary)-10 {
		version, _ := strconv.ParseInt(binary[i:i+3], 2, 8)
		tId, _ := strconv.ParseInt(binary[i+3:i+6], 2, 8)
		sum += version
		i += 6
		if tId == literal {
			more := true
			for more {
				if binary[i] == '0' {
					more = false
				}
				i += 5
			}
			continue
		}
		switch binary[i] {
		case '0':
			i += 16
		case '1':
			i += 12
		}
	}

	return int(sum)
}

func toBinaryString(str string) string {
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	s := make([]byte, len(decoded)*8)
	for i, b := range decoded {
		for j := 0; j < 8; j++ {
			s[i*8+j] = b>>uint(7-j)&0x01 + '0'
		}
	}
	return string(s)
	// sb := strings.Builder{}
	// for _, r := range hex {
	// 	bits, err := strconv.ParseUint(string(r), 16, 4)
	// 	util.Handle(err)
	// 	sb.WriteString(fmt.Sprintf("%04b", bits))
	// }
	// return sb.String()
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	binary := toBinaryString(string(lines[0]))

	var operations []operation
	var done bool
	i := 0

	for {
		if !strings.ContainsRune(binary[i:], '1') {
			done = true
		}

		if len(operations) > 0 {
			op := operations[len(operations)-1]
			if i == op.lastIndex || op.subPkgCount == 0 || done {
				if len(operations) > 1 {
					val := solveOperation(op)
					operations = operations[:len(operations)-1]
					operations[len(operations)-1].values = append(operations[len(operations)-1].values, val)
					continue
				} else {
					return int(solveOperation(op))
				}
			}
			if op.subPkgCount > 0 {
				operations[len(operations)-1].subPkgCount--
			}
		}

		tId, _ := strconv.ParseInt(binary[i+3:i+6], 2, 8)
		i += 6
		if tId == literal {
			more := true
			var num int
			for more {
				if binary[i] == '0' {
					more = false
				}
				for j := 1; j < 5; j++ {
					num <<= 1
					num += int(binary[i+j] % 2)
				}
				i += 5
			}
			if len(operations) > 0 {
				operations[len(operations)-1].values = append(operations[len(operations)-1].values, int64(num))
			} else {
				return num
			}
			continue
		}
		switch binary[i] {
		case '0':
			length, _ := strconv.ParseInt(binary[i+1:i+16], 2, 16)
			i += 16
			operations = append(operations, operation{tId: tId, lastIndex: i + int(length), subPkgCount: -1})
		case '1':
			count, _ := strconv.ParseInt(binary[i+1:i+12], 2, 16)
			operations = append(operations, operation{tId: tId, lastIndex: -1, subPkgCount: count})
			i += 12
		}
	}
}

func solveOperation(o operation) int64 {
	var res int64 = 0
	switch o.tId {
	case sum:
		for _, v := range o.values {
			res += v
		}
	case product:
		res = 1
		for _, v := range o.values {
			res *= v
		}
	case minimum:
		var min int64 = math.MaxInt64
		for _, v := range o.values {
			if v < min {
				min = v
			}
		}
		res = min
	case maximum:
		var max int64 = math.MinInt64
		for _, v := range o.values {
			if v > max {
				max = v
			}
		}
		res = max
	case lessThan:
		if o.values[0] < o.values[1] {
			res = 1
		}
	case greaterThan:
		if o.values[0] > o.values[1] {
			res = 1
		}
	case equal:
		if o.values[0] == o.values[1] {
			res = 1
		}
	}
	return res
}

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

type sfnum struct {
	parent, left, right *sfnum
	value               *int
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	for i := 0; i < len(lines)-1; i++ {
		left := toSfnum(string(lines[i]))
		right := toSfnum(string(lines[i+1]))

		newLeft := &sfnum{left: left, right: right}
		newLeft.left.parent = newLeft
		newLeft.right.parent = newLeft
		newLeft.reduce()
		lines[i+1] = util.Line(newLeft.String())
	}

	last := toSfnum(string(lines[len(lines)-1]))
	last.reduce()
	return *last.magnitude()
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	max := math.MinInt
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if j == i {
				continue
			}
			left := toSfnum(string(lines[i]))
			right := toSfnum(string(lines[j]))
			result := &sfnum{left: left, right: right}
			result.left.parent = result
			result.right.parent = result
			result.reduce()
			mag := result.magnitude()
			if *mag > max {
				max = *mag
			}
		}
	}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if j == i {
				continue
			}
			left := toSfnum(string(lines[i]))
			right := toSfnum(string(lines[j]))
			result := &sfnum{left: right, right: left}
			result.left.parent = result
			result.right.parent = result
			result.reduce()
			mag := result.magnitude()
			if *mag > max {
				max = *mag
			}
		}
	}
	return max
}

func (num *sfnum) magnitude() *int {
	if num.value != nil {
		return num.value
	}
	var first, second *int
	if num.left != nil {
		first = num.left.magnitude()
	}
	if num.right != nil {
		second = num.right.magnitude()
	}
	res := (*first*3 + *second*2)
	num.value = &res
	return num.value
}

func (num *sfnum) explode() {
	// fmt.Printf("exploding [%v,%v]\n", *num.left.value, *num.right.value)
	setNum := func(num *sfnum, add *int) {
		if num != nil {
			new := *num.value + *add
			num.value = &new
		}
	}

	left, _ := findRegularLeft(num)
	setNum(left, num.left.value)

	right, _ := findRegularRight(num)
	setNum(right, num.right.value)

	null := 0
	num.value = &null
	num.left, num.right = nil, nil
}

func (num *sfnum) split() {
	// fmt.Printf("splitting %v\n", fmt.Sprint(num))
	lv := *num.value / 2
	rv := *num.value / 2
	if *num.value%2 == 1 {
		rv++
	}
	num.addLeftValue(lv)
	num.addRightValue(rv)
	num.value = nil
}

func (num *sfnum) reduce() {
	noResult := true
explode:
	for noResult {
		// fmt.Println(num)
		noResult = num.tryExplode(0)
	}
	noResult = true
	for noResult {
		// fmt.Println(num)
		noResult = num.trySplit()
		if noResult {
			goto explode
		}
	}
}

func (num *sfnum) tryExplode(depth int) bool {
	res := false
	if depth > 3 && num.value == nil {
		num.explode()
		return true
	}
	if num.left != nil {
		res = num.left.tryExplode(depth + 1)
	}
	if !res && num.right != nil {
		res = num.right.tryExplode(depth + 1)
	}
	return res
}

func (num *sfnum) trySplit() bool {
	res := false
	if num.left != nil {
		res = num.left.trySplit()
	}
	if !res && num.right != nil {
		res = num.right.trySplit()
	}
	if !res && num.value != nil && *num.value > 9 {
		num.split()
		return true
	}
	return res
}

func findRegularLeft(num *sfnum) (res *sfnum, isFirst bool) {
	for {
		if num.parent == nil {
			return nil, true
		}
		if num.parent.value != nil {
			return num.parent, true
		}
		if num.parent.left != num {
			res := findRightMost(num.parent.left)
			return res, false
		}
		num = num.parent
	}
}

func findRegularRight(num *sfnum) (res *sfnum, isFirst bool) {
	for {
		if num.parent == nil {
			return nil, false
		}
		if num.parent.value != nil {
			return num.parent, false
		}
		if num.parent.right != num {
			res := findLeftMost(num.parent.right)
			return res, true
		}
		num = num.parent
	}
}

func findLeftMost(num *sfnum) *sfnum {
	for {
		if num.left != nil {
			num = num.left
			continue
		}
		return num
	}
}

func findRightMost(num *sfnum) *sfnum {
	for {
		if num.right != nil {
			num = num.right
			continue
		}
		return num
	}
}

func (num *sfnum) addChild(left bool, child *sfnum) (*sfnum, bool) {
	if left {
		num.left = child
	} else {
		num.right = child
	}
	return child, true
}

func (num *sfnum) addValue(left bool, value int) {
	if left {
		num.addLeftValue(value)
	} else {
		num.addRightValue(value)
	}
}

func (num *sfnum) addLeftValue(value int) {
	num.left = &sfnum{value: &value, parent: num}
}

func (num *sfnum) addRightValue(value int) {
	num.right = &sfnum{value: &value, parent: num}
}

func toSfnum(s string) *sfnum {
	num := &sfnum{}
	isLeft := true
	for _, c := range s[1 : len(s)-1] {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num.addValue(isLeft, int(c-'0'))
		case ',':
			isLeft = false
		case '[':
			num, isLeft = num.addChild(isLeft, &sfnum{parent: num})
		case ']':
			num = num.parent
		}
	}
	return num
}

func (num *sfnum) String() string {
	if num.value != nil {
		return fmt.Sprint(*num.value)
	}
	sb := strings.Builder{}
	sb.WriteRune('[')
	if num.left != nil {
		sb.WriteString(fmt.Sprint(num.left))
	}
	sb.WriteRune(',')
	if num.right != nil {
		sb.WriteString(fmt.Sprint(num.right))
	}
	sb.WriteRune(']')
	return sb.String()
}

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

type sfnum struct {
	parent, left, right *sfnum
	first, second       *int
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
	if num.left != nil {
		num.first = num.left.magnitude()
	}
	if num.right != nil {
		num.second = num.right.magnitude()
	}
	res := *num.first*3 + *num.second*2
	return &res
}

func (num *sfnum) explode() {
	// fmt.Printf("exploding [%v,%v]\n", *num.first, *num.second)
	setNum := func(num *sfnum, add *int, isLeft bool) {
		if num != nil {
			if isLeft {
				new := *num.first + *add
				num.first = &new
			} else {
				new := *num.second + *add
				num.second = &new
			}
		}
	}

	left, isLeft := findRegularLeft(num)
	setNum(left, num.first, isLeft)

	right, isLeft := findRegularRight(num)
	setNum(right, num.second, isLeft)

	null := 0
	if num.parent.left == num {
		num.parent.first = &null
		num.parent.left = nil
	}
	if num.parent.right == num {
		num.parent.second = &null
		num.parent.right = nil
	}
}

func (num *sfnum) split(left bool) {
	split := func(from int) *sfnum {
		newLeft := from / 2
		newRight := from / 2
		if from%2 == 1 {
			newRight++
		}
		return &sfnum{first: &newLeft, second: &newRight, parent: num}
	}
	if left {
		// fmt.Printf("splitting [%v,]\n", *num.first)
		num.left = split(*num.first)
		num.first = nil
	} else {
		// fmt.Printf("splitting [,%v]\n", *num.second)
		num.right = split(*num.second)
		num.second = nil
	}
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
	if depth > 3 {
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
	if !res && num.first != nil && *num.first > 9 {
		num.split(true)
		return true
	}
	if !res && num.right != nil {
		res = num.right.trySplit()
	}
	if !res && num.second != nil && *num.second > 9 {
		num.split(false)
		return true
	}
	return res
}

func findRegularLeft(num *sfnum) (res *sfnum, isFirst bool) {
	if num.parent == nil {
		return nil, true
	}
	if num.parent.first != nil {
		return num.parent, true
	}
	if num.parent.left != num {
		res := findRightMost(num.parent.left)
		return res, false
	}
	return findRegularLeft(num.parent)
}

func findRegularRight(num *sfnum) (res *sfnum, isFirst bool) {
	if num.parent == nil {
		return nil, false
	}
	if num.parent.second != nil {
		return num.parent, false
	}
	if num.parent.right != num {
		res := findLeftMost(num.parent.right)
		return res, true
	}
	return findRegularRight(num.parent)
}

func findLeftMost(num *sfnum) *sfnum {
	if num.left != nil {
		return findLeftMost(num.left)
	}
	if num.first != nil {
		return num
	}
	return nil
}

func findRightMost(num *sfnum) *sfnum {
	if num.right != nil {
		return findRightMost(num.right)
	}
	if num.second != nil {
		return num
	}
	return nil
}

func toSfnum(s string) *sfnum {
	num := &sfnum{}
	isLeft := true
	for _, c := range s[1 : len(s)-1] {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if isLeft {
				fst := int(c - '0')
				num.first = &fst
			} else {
				snd := int(c - '0')
				num.second = &snd
				isLeft = true
			}
		case ',':
			isLeft = false
		case '[':
			child := &sfnum{parent: num}
			if isLeft {
				num.left = child
			} else {
				num.right = child
				isLeft = true
			}
			num = child
		case ']':
			num = num.parent
		}
	}
	return num
}

func (num *sfnum) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	if num.left != nil {
		sb.WriteString(fmt.Sprint(num.left))
	}
	if num.first != nil {
		sb.WriteString(fmt.Sprint(*num.first))
	}
	sb.WriteRune(',')
	if num.right != nil {
		sb.WriteString(fmt.Sprint(num.right))
	}
	if num.second != nil {
		sb.WriteString(fmt.Sprint(*num.second))
	}
	sb.WriteRune(']')
	return sb.String()
}

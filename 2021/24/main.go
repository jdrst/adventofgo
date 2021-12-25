package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

type ALU struct {
	instructions []string
	w, x, y, z   *int
}

func partOne(file util.File) int {
	lines := file.AsLines()

	// var w, x, y, z int
	// p := ALU{instructions: lines, w: &w, x: &x, y: &y, z: &z}

	// var res int
	// for i := 11111111111111; i <= 99999999999999; i++ {
	// 	if strings.ContainsAny(strconv.Itoa(i), "0") {
	// 		continue
	// 	}
	// 	rc := hardcoded(i)
	// 	if rc == 0 {
	// 		fmt.Printf("found valid: %d\n", i)
	// 		res = i
	// 	}
	// 	// rc := p.process(i)
	// 	// if rc == 0 {
	// 	// 	return i
	// 	// }
	// }
	// return res
	// return findPossibleSolutions(true)
	return solve(parseInstructions(lines), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func parseInstructions(lines util.Lines) [14]instruction {
	res := [14]instruction{}
	aLine, bLine, cLine := 4, 5, 15
	for i := 0; i < 14; i++ {
		aVals := lines[aLine].SubSplitWith(" ")
		bVals := lines[bLine].SubSplitWith(" ")
		cVals := lines[cLine].SubSplitWith(" ")
		res[i] = instruction{a: aVals[2].AsInt(), b: bVals[2].AsInt(), c: cVals[2].AsInt()}
		aLine += 18
		bLine += 18
		cLine += 18
	}
	return res
}

func (alu *ALU) process(num int) int {
	str := strconv.Itoa(num)
	var idx int
	null := 0
	alu.w = &null
	alu.x = &null
	alu.y = &null
	alu.z = &null
	for _, l := range alu.instructions {
		instruction := strings.Split(l, " ")

		var variable **int
		var res int

		switch instruction[1] {
		case "w":
			variable = &alu.w
		case "x":
			variable = &alu.x
		case "y":
			variable = &alu.y
		case "z":
			variable = &alu.z
		}
		switch instruction[0] {
		case "inp":
			res = int(str[idx] - '0')
			*variable = &res
			idx++
		case "add":
			res = **variable + alu.IntOrRegister(instruction[2])
			*variable = &res
		case "mul":
			res = **variable * alu.IntOrRegister(instruction[2])
			*variable = &res
		case "div":
			res = **variable / alu.IntOrRegister(instruction[2])
			*variable = &res
		case "mod":
			res = **variable % alu.IntOrRegister(instruction[2])
			*variable = &res
		case "eql":
			if **variable == alu.IntOrRegister(instruction[2]) {
				res = 1
			} else {
				res = 0
			}
			*variable = &res
		default:
			panic("i don't know what to do")
		}
	}
	return *alu.z
}

func (alu *ALU) IntOrRegister(c string) int {
	if c[0] > '9' {
		switch c {
		case "w":
			return *alu.w
		case "x":
			return *alu.x
		case "y":
			return *alu.y
		case "z":
			return *alu.z
		}
	}
	res := util.ToInt(c)
	return res
}

func partTwo(file util.File) int {
	lines := file.AsLines()
	return solve(parseInstructions(lines), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func hardcoded(num int) int {
	str := strconv.Itoa(num)

	var z int
	stacksize := 0
	getW := func(idx int) int {
		return int(str[idx] - '0')
	}
	instr := func(w, a, b, c int) {
		if z%26+b == w {
			z /= a
			if a > 1 {
				stacksize--
			}
			return
		}
		z /= a
		if a > 1 {
			stacksize--
		}
		z = 26*z + w + c
		stacksize++
	}

	//vorhandener wert%26 + b == neue ziffer dann /26

	// w0 := int(str[0] - '0')
	// w1 := int(str[1] - '0')
	// w2 := int(str[2] - '0')
	// w3 := int(str[3] - '0')

	instr(getW(0), 1, 13, 15) //push w0+15
	instr(getW(1), 1, 13, 16) //push w1+16
	instr(getW(2), 1, 10, 4)  //push w2+4
	instr(getW(3), 1, 15, 14) //push w3+14
	// z = 17576*w0 + 676*w1 + 26*w2 + w3 + 274574 // 4 x w(n)+c(n) auf dem stack
	instr(getW(4), 26, -8, 1)    //pop (w3+14); if (w3+14)%26-8 == w4 -> z/26 (pop) else push (w3+14)+(w4+1)
	instr(getW(5), 26, -10, 5)   //pop ..
	instr(getW(6), 1, 11, 1)     //push
	instr(getW(7), 26, -3, 3)    //pop
	instr(getW(8), 1, 14, 3)     //push
	instr(getW(9), 26, -4, 7)    //pop
	instr(getW(10), 1, 14, 5)    //push
	instr(getW(11), 26, -5, 13)  //pop
	instr(getW(12), 26, -8, 3)   //pop
	instr(getW(13), 26, -11, 10) //pop
	return z
}

type instruction struct {
	a, b, c int
}

func findPossibleSolutions(max bool) int {
	var z int
	// stacksize := 0

	// calc := func(w, a, b, c int) {
	// 	if a > 1 {
	// 		stacksize--
	// 	}
	// 	if (z%26)+b == w {
	// 		z /= a
	// 		return
	// 	}
	// 	z /= a
	// 	z = 26*z + w + c
	// 	stacksize++
	// }

	// sol := [14]int{9, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sol := [14]int{}
	instructions := []instruction{
		{1, 13, 15},
		{1, 13, 16},
		{1, 10, 4},
		{1, 15, 14},
		{26, -8, 1},
		{26, -10, 5},
		{1, 11, 1},
		{26, -3, 3},
		{1, 14, 3},
		{26, -4, 7},
		{1, 14, 5},
		{26, -5, 13},
		{26, -8, 3},
		{26, -11, 10},
	}
	for i := range sol {
		digit := 10
		for j := 9; j > 0; j-- {
			in := instructions[i]
			//we NEED to pop and we only pop if (z%26)+b == w
			if in.a == 26 {
				if (z%26)+in.b == j {
					z /= in.a
					digit = j
					break
				}
				continue
			}
			z = 26*z + j + in.c
			digit = j
			break
			// if in.a == 26 && j-a {
			// 	break
			// }
			// stacksize = currentSize
			// z = currentZ
		}
		// in := instructions[i]
		// calc(j-1, in.a, in.b, in.c)
		sol[i] = digit
	}
	fmt.Println(sol)
	res := 0
	for _, v := range sol {
		res *= 10
		res += v
	}
	return res
}

func divideBy26times(z, times int) int {
	for times > 0 {
		z /= 26
		times--
	}
	return z
}

func solve(instructions [14]instruction, order []int) int {
	sol := [14]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	stack := []int{}
	functions := [14](func(int) int){}
	controlFunctionIdxs := []int{}
	for i, in := range instructions {
		switch in.a {
		case 1:
			c := in.c
			fn := func(n int) int { return n + c }
			functions[i] = fn
			stack = append(stack, i)
		case 26:
			idx := stack[len(stack)-1]
			current := i
			b := in.b
			fn := func(n int) int {
				sol[idx] = n
				sol[current] = functions[idx](n) + b
				return 0
			}
			controlFunctionIdxs = append(controlFunctionIdxs, current)
			stack = stack[:len(stack)-1]
			functions[i] = fn
		}
	}

	for _, idx := range controlFunctionIdxs {
		for _, i := range order {
			functions[idx](i)
			if isValid(sol) {
				break
			}
		}
	}

	return toInt(sol)
}

func isValid(sol [14]int) bool {
	for _, v := range sol {
		if v < 1 || v > 9 {
			return false
		}
	}
	return true
}

func toInt(sol [14]int) int {
	res := 0
	for _, v := range sol {
		res *= 10
		res += v
	}
	return res
}

// type stack []int

// func (s *stack) push(n int) {
// 	new := stack(append([]int(*s), n))
// 	s = &new
// }

// func (s stack) pop() int {
// 	res := s[len(s)-1]
// 	s = s[:len(s)-1]
// 	return res
// }

// type stack [](*func(int) int)

// func (s stack) push(n func(int) int) {
// 	s = append(s, &n)
// }

// func (s stack) pop() *func(int) int {
// 	res := s[len(s)-1]
// 	s = s[:len(s)-1]
// 	return res
// }

func foo(instructions [14]instruction) int {
	// instructions := []instruction{
	// 	{1, 13, 15}, d1= w1+15 / 5 1
	// 	{1, 13, 16}, d2= w2+16 / 1 1
	// 	{1, 10, 4},  d3= w3+4 / 9 7
	// 	{1, 15, 14}, d4= w4+14 / 3 1
	// 	{26, -8, 1}, d5= w4+14+-8 / 9 7
	// 	{26, -10, 5}, d6= w3+4+-10 / 3 1
	// 	{1, 11, 1}, d7= w7+1 / 9 3
	// 	{26, -3, 3}, d8= w7+1+-3 / 7 1
	// 	{1, 14, 3}, d9= w8+3 / 9 2
	// 	{26, -4, 7}, d10= w8+3+-4 / 8 1
	// 	{1, 14, 5}, d11= w11+5 / 9 1
	// 	{26, -5, 13}, d12= w11+5-5 / 9 1
	// 	{26, -8, 3}, d13= w2+16-8 / 9 9
	// 	{26, -11, 10}, d14= w1+15-11 / 9 5
	// }
	var solve func([14]int, int, int, int) int
	solve = func(sol [14]int, i int, z int, w int) int {
		if i > 13 {
			if z == 0 {
				res := 0
				for _, v := range sol {
					res *= 10
					res += v
				}
				return res
			}
			return 0
		}

		in := instructions[i]
		//we NEED to pop and we only pop if (z%26)+b == w
		if in.a == 26 {
			if (z%26)+in.b != w {
				return 0
			}
			z /= in.a //POP
			new := cpArray(sol)
			new[i] = w
			res := math.MinInt
			for j := 9; j > 0; j-- {
				res = max(res, solve(new, i+1, z, j))
			}
			return res
		} else {
			res := math.MinInt
			for j := 9; j > 0; j-- {
				newZ := 26*z + j + in.c //PUSH
				new := cpArray(sol)
				new[i] = w
				res = max(res, solve(new, i+1, newZ, j))
			}
			return res
		}
	}

	res := math.MinInt
	for j := 9; j > 0; j-- {
		res = max(res, solve([14]int{}, 0, 0, j))
	}
	return res
}

func cpArray(old [14]int) [14]int {
	c := [14]int{}
	for i, v := range old {
		c[i] = v
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
z = 0

w = 0
x = eql((z%26)+13, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+15)*x

w = 1
x = eql((z%26)+13, 1)
x = eql(x, w)
z /= 1
z = z*((25*x)+1) + (w+16)*x

w = 2
x = eql((z%26)+10, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+4)*x

w = 3
x = eql((z%26)+15, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+14)*x

w = 4
x = eql((z%26)+-8, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+1)*x

w = 5
x = eql((z%26)+-10, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+5)*x

w = 6
x = eql((z%26)+11, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+1)*x

w = 7
x = eql((z%26)+-3, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+3)*x

w = 8
x = eql((z%26)+14, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+3)*x

w = 9
x = eql((z%26)+-4, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+7)*x

w = 10
x = eql((z%26)+14, w)
x = eql(x, 0)
z /= 1
z = z*((25*x)+1) + (w+5)*x

w = 11
x = eql((z%26)+-5, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+13)*x

w = 12
x = eql((z%26)+-8, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+3)*x

w = 13
x = eql((z%26)+-11, w)
x = eql(x, 0)
z /= 26
z = z*((25*x)+1) + (w+10)*x

return z
*/

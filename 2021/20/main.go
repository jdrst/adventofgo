package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
	fmt.Printf("Second part: %v\n", partTwo(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	algorithm := string(lines[0])

	image := makeImage(lines[2:])

	return enhanceXTimes(2, image, algorithm)
}

func partTwo(file util.File) int {
	lines := file.AsLines()

	algorithm := string(lines[0])

	image := makeImage(lines[2:])

	return enhanceXTimes(50, image, algorithm)
}

func enhanceXTimes(x int, image [][]bool, algorithm string) int {
	isAlternating := algorithm[0] == '#' && algorithm[len(algorithm)-1] == '.'

	for i := 0; i < x; i++ {
		image = enhance(image, algorithm, isAlternating && i%2 == 1)
	}

	cnt := 0
	for _, l := range image {
		for _, b := range l {
			if b {
				cnt++
			}
		}
	}

	return cnt
}

func enhance(image [][]bool, algorithm string, outerPx bool) [][]bool {
	image = expand(image, outerPx)

	new := make([][]bool, len(image))
	for i, l := range image {
		new[i] = make([]bool, len(l))
	}

	for i, l := range image {
		for j := range l {
			idx := getNewPixel(image, i, j, outerPx)
			if algorithm[idx] == '#' {
				new[i][j] = true
			}
		}
	}
	return new
}

func expand(image [][]bool, with bool) [][]bool {
	//todo, expanding with append?
	new := make([][]bool, len(image)+2)
	new[0] = make([]bool, len(image[0])+2)
	new[len(new)-1] = make([]bool, len(image[0])+2)
	for i := range new[0] {
		new[0][i] = with
		new[len(new)-1][i] = with
	}
	for i := range image {
		new[i+1] = append([]bool{with}, append(image[i], with)...)
	}
	return new
}

func getNewPixel(image [][]bool, x, y int, outerPx bool) int {
	res := 0
	points := gridFor(x, y)
	isLit := func(x, y int) bool {
		if x < 0 || x > len(image)-1 || y < 0 || y > len(image)-1 {
			return outerPx
		}
		return image[x][y]
	}
	for _, p := range points {
		res <<= 1
		if isLit(p.X, p.Y) {
			res += 1
		}
	}
	return res
}

func gridFor(x, y int) []util.Point {
	res := make([]util.Point, 0)
	res = append(res, util.Point{X: x - 1, Y: y - 1})
	res = append(res, util.Point{X: x - 1, Y: y})
	res = append(res, util.Point{X: x - 1, Y: y + 1})
	res = append(res, util.Point{X: x, Y: y - 1})
	res = append(res, util.Point{X: x, Y: y})
	res = append(res, util.Point{X: x, Y: y + 1})
	res = append(res, util.Point{X: x + 1, Y: y - 1})
	res = append(res, util.Point{X: x + 1, Y: y})
	res = append(res, util.Point{X: x + 1, Y: y + 1})
	return res
}

func makeImage(lines util.Lines) [][]bool {
	image := make([][]bool, len(lines))

	for i, l := range lines {
		image[i] = make([]bool, len(l))
		for j, c := range l {
			switch c {
			case '.':
				image[i][j] = false
			case '#':
				image[i][j] = true
			}
		}
	}
	return image
}

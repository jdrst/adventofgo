package main

import (
	"fmt"

	"github.com/jdrst/adventofgo/util"
)

type floortile rune

// type marina [][]floortile

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
}

func partOne(file util.File) int {
	lines := file.AsLines()

	width := len(lines[0])
	height := len(lines)

	floor := make([][]floortile, height)

	for i, l := range lines {
		floor[i] = make([]floortile, width)
		for j, c := range l {
			floor[i][j] = floortile(c)
		}
	}

	moved, cnt := true, 0
	for moved {
		cnt++
		moved = false
		for i, l := range floor {
			firstOccupied := l[0] != '.'
			for j := 0; j < len(l); j++ {
				c := floor[i][j]
				if c == '.' || j == len(l)-1 && firstOccupied {
					continue
				}
				if c == '>' {
					next := (j + 1) % len(floor[i])
					if floor[i][next] == '.' {
						floor[i][j] = '.'
						floor[i][next] = '>'
						j++
						moved = true
					}
				}
			}
		}
		for j := 0; j < len(floor[0]); j++ {
			firstOccupied := floor[0][j] != '.'
			for i := 0; i < len(floor); i++ {
				c := floor[i][j]
				if c == '.' || i == len(floor)-1 && firstOccupied {
					continue
				}
				if c == 'v' {
					next := (i + 1) % len(floor)
					if floor[next][j] == '.' {
						floor[i][j] = '.'
						floor[next][j] = 'v'
						i++
						moved = true
					}
				}
			}
		}
		// fmt.Println(marina(floor))
	}

	return cnt
}

// func (f marina) String() string {
// 	b := strings.Builder{}
// 	for _, l := range f {
// 		for _, c := range l {
// 			b.WriteRune(rune(c))
// 		}
// 		b.WriteRune('\n')
// 	}
// 	return b.String()
// }

package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	newLine = "\n"
	lines := prepareInput([]byte(testinput1))
	currentPos := ship{point{0, 0}, point{1, 0}}
	currentPos.move(lines[0], false)
	actualDir, actualX, actualY := currentPos.direction, currentPos.position.x, currentPos.position.y
	expectedDir, expectedX, expectedY := point{1, 0}, 10, 0
	if actualX != expectedX {
		t.Errorf("X-Coord not as expected, got: %d, want: %d.", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("Y-Coord not as expected, got: %d, want: %d.", actualY, expectedY)
	}
	if actualDir != expectedDir {
		t.Errorf("Direction not as expected, got: %v, want: %v.", actualDir, expectedDir)
	}
	currentPos.move(lines[1], false)
	actualDir, actualX, actualY = currentPos.direction, currentPos.position.x, currentPos.position.y
	expectedDir, expectedX, expectedY = point{1, 0}, 10, 3
	if actualX != expectedX {
		t.Errorf("X-Coord not as expected, got: %d, want: %d.", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("Y-Coord not as expected, got: %d, want: %d.", actualY, expectedY)
	}
	if actualDir != expectedDir {
		t.Errorf("Direction not as expected, got: %v, want: %v.", actualDir, expectedDir)
	}
	currentPos.move(lines[2], false)
	actualDir, actualX, actualY = currentPos.direction, currentPos.position.x, currentPos.position.y
	expectedDir, expectedX, expectedY = point{1, 0}, 17, 3
	if actualX != expectedX {
		t.Errorf("X-Coord not as expected, got: %d, want: %d.", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("Y-Coord not as expected, got: %d, want: %d.", actualY, expectedY)
	}
	if actualDir != expectedDir {
		t.Errorf("Direction not as expected, got: %v, want: %v.", actualDir, expectedDir)
	}
	currentPos.move(lines[3], false)
	actualDir, actualX, actualY = currentPos.direction, currentPos.position.x, currentPos.position.y
	expectedDir, expectedX, expectedY = point{0, -1}, 17, 3
	if actualX != expectedX {
		t.Errorf("X-Coord not as expected, got: %d, want: %d.", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("Y-Coord not as expected, got: %d, want: %d.", actualY, expectedY)
	}
	if actualDir != expectedDir {
		t.Errorf("Direction not as expected, got: %v, want: %v.", actualDir, expectedDir)
	}
	currentPos.move(lines[4], false)
	actualDir, actualX, actualY = currentPos.direction, currentPos.position.x, currentPos.position.y
	expectedDir, expectedX, expectedY = point{0, -1}, 17, -8
	if actualX != expectedX {
		t.Errorf("X-Coord not as expected, got: %d, want: %d.", actualX, expectedX)
	}
	if actualY != expectedY {
		t.Errorf("Y-Coord not as expected, got: %d, want: %d.", actualY, expectedY)
	}
	if actualDir != expectedDir {
		t.Errorf("Direction not as expected, got: %v, want: %v.", actualDir, expectedDir)
	}
}

func BenchmarkMain(b *testing.B) {
	newLine = "\r\n"
	for n := 0; n < b.N; n++ {
		main()
	}
}

const testinput1 = `F10
N3
F7
R90
F11`

package main

import (
	"testing"
)

func TestExercise5(t *testing.T) {
	expected := 5
	actual := solveExercise5("file05_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise5b(t *testing.T) {
	expected := 12
	actual := solveExercise5b("file05_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestIsPointOnLine(t *testing.T) {
	var line LineOfVents
	line.x1 = 4
	line.y1 = 0
	line.x2 = 0
	line.y2 = 4

	if !isPointOnLine(line, 4, 0) {
		t.Fatalf("Was not the right solution. Expected %d,%d is on line.", 4, 0)
	}
	if !isPointOnLine(line, 3, 1) {
		t.Fatalf("Was not the right solution. Expected %d,%d is on line.", 3, 1)
	}
	if !isPointOnLine(line, 2, 2) {
		t.Fatalf("Was not the right solution. Expected %d,%d is on line.", 2, 2)
	}
	if !isPointOnLine(line, 1, 3) {
		t.Fatalf("Was not the right solution. Expected %d,%d is on line.", 1, 3)
	}
	if !isPointOnLine(line, 0, 4) {
		t.Fatalf("Was not the right solution. Expected %d,%d is on line.", 0, 4)
	}

	if isPointOnLine(line, 0, 0) {
		t.Fatalf("Was not the right solution. Expected %d,%d is not on line.", 0, 0)
	}
	if isPointOnLine(line, 4, 1) {
		t.Fatalf("Was not the right solution. Expected %d,%d is not on line.", 4, 1)
	}
}

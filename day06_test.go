package main

import (
	"testing"
)

func TestExercise6(t *testing.T) {
	expected := 5934
	actual := solveExercise6("file06_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise6b(t *testing.T) {
	expected := 26984457539
	actual := solveExercise6b("file06_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

package main

import (
	"testing"
)

func TestExercise9(t *testing.T) {
	expected := 0
	actual := solveExercise9("file09_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise9b(t *testing.T) {
	expected := 0
	actual := solveExercise9b("file09_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

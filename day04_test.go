package main

import (
	"testing"
)

func TestExercise4(t *testing.T) {
	expected := 4512
	actual := solveExercise4("file04_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise4b(t *testing.T) {
	expected := 1924
	actual := solveExercise4b("file04_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

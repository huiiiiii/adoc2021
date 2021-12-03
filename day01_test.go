package main

import (
	"testing"
)

func TestExercise1(t *testing.T) {
	expected := 7
	actual := solveExercise1("file01_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise1b(t *testing.T) {
	expected := 5
	actual := solveExercise1b("file01_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

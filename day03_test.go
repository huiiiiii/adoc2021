package main

import (
	"testing"
)

func TestExercise3(t *testing.T) {
	expected := 198
	actual := solveExercise3("file03_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise3b(t *testing.T) {
	expected := 230
	actual := solveExercise3b("file03_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

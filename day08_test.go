package main

import (
	"testing"
)

func TestExercise8(t *testing.T) {
	expected := 0
	actual := solveExercise8("file08_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise8b(t *testing.T) {
	expected := 0
	actual := solveExercise8b("file08_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

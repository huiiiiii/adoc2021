package main

import (
	"testing"
)

func TestExercise4(t *testing.T) {
	expected := 0
	actual := solveExercise4("fileß4_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise4b(t *testing.T) {
	expected := 0
	actual := solveExercise4b("file04_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

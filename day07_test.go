package main

import (
	"testing"
)

func TestExercise7(t *testing.T) {
	expected := 37
	actual := solveExercise7("file07_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise7b(t *testing.T) {
	expected := 168
	actual := solveExercise7b("file07_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

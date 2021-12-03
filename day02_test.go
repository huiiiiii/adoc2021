package main

import (
	"testing"
)

func TestExercise2(t *testing.T) {
	expected := 150
	actual := solveExercise2("file02_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise2a(t *testing.T) {
	expected := 150
	actual := solveExercise2a("file02_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

func TestExercise2b(t *testing.T) {
	expected := 900
	actual := solveExercise2b("file02_test.txt")
	if expected != int(actual) {
		t.Fatalf("Was not the right solution. Expected %d, but was actual %d", expected, actual)
	}
}

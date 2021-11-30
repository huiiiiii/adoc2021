package main

import "strings"

func solveExercise1() (solution string) {
	var file = readFile("file1.txt")
	return "not yet solved file: " + strings.Join(file, ",")
}

package main

import (
	"strings"
)

type NavStep struct {
	direction string
	number    int
}

func solveExercise2() (solution int) {
	lines := readFile("file2.txt")
	horizontalPos := 0
	depth := 0
	for _, line := range lines {
		navStep := stringToNavStep(line)
		switch navStep.direction {
		case "forward":
			horizontalPos += navStep.number
		case "down":
			depth += navStep.number
		case "up":
			depth -= navStep.number
		}
	}
	solution = horizontalPos * depth
	return
}

func solveExercise2b() (solution int) {
	lines := readFile("file2.txt")
	horizontalPos := 0
	aim := 0
	depth := 0
	for _, line := range lines {
		navStep := stringToNavStep(line)
		switch navStep.direction {
		case "forward":
			horizontalPos += navStep.number
			depth += aim * navStep.number
		case "down":
			aim += navStep.number
		case "up":
			aim -= navStep.number
		}
	}
	solution = horizontalPos * depth
	return
}

func stringToNavStep(line string) (navStep NavStep) {
	lineArray := strings.Split(line, " ")
	navStep.direction = lineArray[0]
	navStep.number = strToInt(lineArray[1])
	return
}

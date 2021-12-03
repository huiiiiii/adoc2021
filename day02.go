package main

import (
	"reflect"
	"strings"
)

type NavStep struct {
	direction string
	number    int
}

func solveExercise2(fileName string) (solution int) {
	lines := readFile(fileName)
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

func solveExercise2b(fileName string) (solution int) {
	lines := readFile(fileName)
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

//---------- alternative solution ----------

type stubMapping map[string]interface{}

var horizontalPos int
var depth int
var stubStorage = stubMapping{}

func solveExercise2a(fileName string) (solution int) {
	lines := readFile(fileName)
	horizontalPos = 0
	depth = 0
	stubStorage = map[string]interface{}{
		"forward": forward,
		"up":      up,
		"down":    down,
	}
	for _, line := range lines {
		navStep := stringToNavStep(line)
		callFunctionByName(stubStorage[navStep.direction], navStep.number)
	}
	solution = horizontalPos * depth
	return
}

func forward(number int) {
	horizontalPos += number
}

func down(number int) {
	depth += number
}

func up(number int) {
	depth -= number
}

func callFunctionByName(funcName interface{}, params ...interface{}) {
	f := reflect.ValueOf(funcName)
	if len(params) != f.Type().NumIn() {
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	f.Call(in)
}

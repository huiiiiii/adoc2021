package main

import "reflect"

type stubMapping map[string]interface{}

var horizontalPos int
var depth int
var stubStorage = stubMapping{}

func solveExercise2a() (solution int) {
	lines := readFile("file2.txt")
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

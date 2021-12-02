package main

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
		callFunction(stubStorage[navStep.direction], navStep.number)
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

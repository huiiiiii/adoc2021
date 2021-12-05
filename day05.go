package main

import (
	"math"
	"strings"
)

type LineOfVents struct {
	x1, y1, x2, y2 int
}

func solveExercise5(fileName string) (solution int) {
	lines := readFile(fileName)
	linesOfVents := linesToLinesOfVents(lines)
	horAndVertLines := getHorizontalAndVerticalLines(linesOfVents)
	ventsDiagram := calculateVentsDiagram(horAndVertLines)
	solution = countDangerousAreas(ventsDiagram)
	return
}

func solveExercise5b(fileName string) (solution int) {
	lines := readFile(fileName)
	linesOfVents := linesToLinesOfVents(lines)
	ventsDiagram := calculateVentsDiagram(linesOfVents)
	solution = countDangerousAreas(ventsDiagram)
	return
}

func linesToLinesOfVents(lines []string) (linesOfVents []LineOfVents) {
	for _, line := range lines {
		linesOfVents = append(linesOfVents, lineToLineOfVents(line))
	}
	return
}

func lineToLineOfVents(line string) (lineOfVents LineOfVents) {
	points := strings.Split(line, " -> ")
	point1 := strToInts(points[0], ",")
	point2 := strToInts(points[1], ",")
	lineOfVents.x1 = point1[0]
	lineOfVents.y1 = point1[1]
	lineOfVents.x2 = point2[0]
	lineOfVents.y2 = point2[1]
	return
}

func getHorizontalAndVerticalLines(linesOfVents []LineOfVents) (horizontalAndVerticalLines []LineOfVents) {
	for _, line := range linesOfVents {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			horizontalAndVerticalLines = append(horizontalAndVerticalLines, line)
		}
	}
	return
}

func calculateVentsDiagram(linesOfVents []LineOfVents) (ventsDiagram [][]int) {
	ventsDiagram = initEmptyVentsDiagram(linesOfVents)
	for _, line := range linesOfVents {
		for x := getMin(line.x1, line.x2); x <= getMax(line.x1, line.x2); x++ {
			for y := getMin(line.y1, line.y2); y <= getMax(line.y1, line.y2); y++ {
				if isPointOnLine(line, x, y) {
					ventsDiagram[x][y] += 1
				}
			}
		}
	}
	return ventsDiagram
}

func initEmptyVentsDiagram(linesOfVents []LineOfVents) (ventsDiagram [][]int) {
	xMax, yMax := getMaxCoordinates(linesOfVents)
	for x := 0; x <= xMax; x++ {
		ventsDiagram = append(ventsDiagram, getEmptyIntArray(yMax))
	}
	return
}

func getMaxCoordinates(linesOfVents []LineOfVents) (xMax, yMax int) {
	for _, line := range linesOfVents {
		xMax = getMax(getMax(line.x1, line.x2), xMax)
		yMax = getMax(getMax(line.y1, line.y2), yMax)
	}
	return
}

func isPointOnLine(line LineOfVents, x, y int) bool {
	distanceAB := distance(line.x1, line.y1, x, y)
	distanceBC := distance(x, y, line.x2, line.y2)
	distanceAC := distance(line.x1, line.y1, line.x2, line.y2)
	return almostEqual(distanceAB+distanceBC, distanceAC, 0.00000000001)
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0))
}

func countDangerousAreas(ventsDiagram [][]int) (solution int) {
	for _, xLine := range ventsDiagram {
		for _, ventCount := range xLine {
			if ventCount >= 2 {
				solution++
			}
		}
	}
	return
}

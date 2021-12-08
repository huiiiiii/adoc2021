package main

import (
	"math"
	"sort"
)

func solveExercise7(fileName string) (solution int) {
	lines := readFile(fileName)
	positions := strToInts(lines[0], ",")
	sort.Ints(positions)
	middlePosition := positions[len(positions)/2]
	solution = calculateFuelLin(positions, middlePosition)
	return
}

func calculateFuelLin(positions []int, middlePosition int) (fuel int) {
	for _, position := range positions {
		fuel += int(math.Abs(float64(position) - float64(middlePosition)))
	}
	return
}

func solveExercise7b(fileName string) (solution int) {
	lines := readFile(fileName)
	positions := strToInts(lines[0], ",")
	sort.Ints(positions)
	solution = calculateFuelb(positions, 0)
	for i := positions[0]; i <= positions[len(positions)-1]; i++ {
		solution = getMin(solution, calculateFuelb(positions, i))
	}
	return
}

func calculateFuelb(positions []int, middlePosition int) (fuel int) {
	for _, position := range positions {
		distance := int(math.Abs(float64(position) - float64(middlePosition)))

		fuel += (distance * (distance + 1)) / 2
	}
	return
}

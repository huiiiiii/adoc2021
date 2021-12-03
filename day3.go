package main

import (
	"math"
	"strconv"
)

func solveExercise3() (solution int64) {
	lines := readFile("file3.txt")
	digitCount := len(lines[0])
	gammaInBin := ""
	for i := 0; i < digitCount; i++ {
		gammaInBin += getMostCommonAtPosition(lines, i)
	}
	gamma, _ := strconv.ParseInt(gammaInBin, 2, 64)
	epsilon := int(math.Pow(2, float64(digitCount))) - 1 - int(gamma)
	solution = gamma * int64(epsilon)
	return
}

func solveExercise3b() int64 {
	oxygenGeneratorRating := calculateRating(getMostCommonAtPosition)
	co2ScrubberRating := calculateRating(getLessCommonAtPosition)
	return oxygenGeneratorRating * co2ScrubberRating
}

func calculateRating(getCommonAtPosition func([]string, int) string) (solution int64) {
	lines := readFile("file3.txt")
	digitCount := len(lines[0])
	for i := 0; i < digitCount; i++ {
		if len(lines) == 1 {
			break
		}
		digitAtPosition := getCommonAtPosition(lines, i)
		lines = reduceLinesByDigitAtPosition(lines, strToInt(digitAtPosition), i)
	}
	solution, _ = strconv.ParseInt(lines[0], 2, 64)
	return
}

func reduceLinesByDigitAtPosition(lines []string, digit int, position int) []string {
	for i := len(lines) - 1; i >= 0; i-- {
		if getDigitAtPlace(lines[i], position) != digit {
			lines = removeFromArray(lines, i)
		}
	}
	return lines
}

func removeFromArray(lines []string, i int) []string {
	lines[i] = lines[len(lines)-1]
	return lines[:len(lines)-1]
}

func getMostCommonAtPosition(lines []string, index int) (solution string) {
	count := getCountAtPosition(lines, index)
	countHalfLines := float32(len(lines)) / 2.0
	solution = "0"
	if float32(count) >= countHalfLines {
		solution = "1"
	}
	return
}

func getLessCommonAtPosition(lines []string, index int) (solution string) {
	count := getCountAtPosition(lines, index)
	countHalfLines := float32(len(lines)) / 2.0
	solution = "0"
	if float32(count) < countHalfLines {
		solution = "1"
	}
	return
}

func getCountAtPosition(lines []string, index int) (countAtIndex int) {
	for _, line := range lines {
		countAtIndex += getDigitAtPlace(line, index)
	}
	return
}

func getDigitAtPlace(num string, place int) int {
	return strToInt(string(num[place]))
}

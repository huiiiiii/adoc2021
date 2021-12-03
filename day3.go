package main

import (
	"math"
	"strconv"
)

//type getCommonAtPosition func(lines []string, index int) string

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

func solveExercise3b() (solution int64) {
	gammaLines := readFile("file3.txt")
	digitCount := len(gammaLines[0])
	for i := 0; i < digitCount; i++ {
		mostCommon := getMostCommonAtPosition(gammaLines, i)
		gammaLines = reduceLinesByDigitAtPosition(gammaLines, strToInt(mostCommon), i)
	}
	epsilonLines := readFile("file3.txt")
	for i := 0; i < digitCount; i++ {
		if len(epsilonLines) == 1 {
			break
		}
		lessCommon := getLessCommonAtPosition(epsilonLines, i)
		epsilonLines = reduceLinesByDigitAtPosition(epsilonLines, strToInt(lessCommon), i)
	}
	gamma, _ := strconv.ParseInt(gammaLines[0], 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonLines[0], 2, 64)
	solution = gamma * int64(epsilon)
	return
}

/*ToDo
func a(fn getCommonAtPosition) string {
	gammaLines := readFile("file3.txt")
	digitCount := len(gammaLines[0])
	for i := 0; i < digitCount; i++ {
		mostCommon := getCommonAtPosition(gammaLines, i)
		gammaLines = reduceLinesByDigitAtPosition(gammaLines, strToInt(mostCommon), i)
	}
}*/

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

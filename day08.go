package main

import (
	"strings"
)

type DisplayDigit struct {
	digit                 int
	displayRepresentation string
}

type Display struct {
	uniqueSignalPatterns []DisplayDigit
	outputValues         []DisplayDigit
}

//{{0, "abcefg"}, {1, "cf"}, {2, "acdeg"}, {3, "acdfg"}, {4, "bcdf"}, {5, "abdfg"}, {6, "abdefg"}, {7, "acf"}, {8, "abcdefg"}, {9, "abcdfg"}}

// count 1, 4, 7, or 8
func solveExercise8(fileName string) (solution int) {
	lines := readFile(fileName)
	displays := linesToDisplays(lines)
	displays = translateDisplays(displays)
	solution = count1478(displays)
	return
}

func solveExercise8b(fileName string) (solution int) {
	lines := readFile(fileName)
	displays := linesToDisplays(lines)
	displays = translateDisplays(displays)
	solution = sumOutputValues(displays)
	return
}

func translateDisplays(displays []Display) (translatedDisplays []Display) {
	for _, display := range displays {
		display = translateDigits1478(display)
		display = translateDigits069(display)
		display = translateDigits235(display)
		display = setOutputDigits(display)
		translatedDisplays = append(translatedDisplays, display)
	}
	return
}

func translateDigits1478(display Display) Display {
	for i, pattern := range display.uniqueSignalPatterns {
		if len(pattern.displayRepresentation) == 2 {
			pattern.digit = 1
		}
		if len(pattern.displayRepresentation) == 3 {
			pattern.digit = 7
		}
		if len(pattern.displayRepresentation) == 4 {
			pattern.digit = 4
		}
		if len(pattern.displayRepresentation) == 7 {
			pattern.digit = 8
		}
		display.uniqueSignalPatterns[i] = pattern
	}
	return display
}

func translateDigits069(display Display) Display {
	for i, pattern := range display.uniqueSignalPatterns {
		if len(pattern.displayRepresentation) == 6 {
			pattern1 := getPatternWithDigit(display, 1)
			if containsPattern(pattern.displayRepresentation, pattern1.displayRepresentation) {
				pattern4 := getPatternWithDigit(display, 4)
				if containsPattern(pattern.displayRepresentation, pattern4.displayRepresentation) {
					pattern.digit = 9
				} else {
					pattern.digit = 0
				}
			} else {
				pattern.digit = 6
			}
		}
		display.uniqueSignalPatterns[i] = pattern
	}
	return display
}

func getPatternWithDigit(display Display, digit int) (digitPattern DisplayDigit) {
	for _, pattern := range display.uniqueSignalPatterns {
		if pattern.digit == digit {
			digitPattern = pattern
		}
	}
	return
}

func containsPattern(pattern, patternToContain string) bool {
	contains := true
	for _, char := range patternToContain {
		if !strings.Contains(pattern, string(char)) {
			contains = false
		}
	}
	return contains
}

func translateDigits235(display Display) Display {
	for i, pattern := range display.uniqueSignalPatterns {
		if len(pattern.displayRepresentation) == 5 {
			pattern1 := getPatternWithDigit(display, 1)
			pattern6 := getPatternWithDigit(display, 6)
			if containsPattern(pattern6.displayRepresentation, pattern.displayRepresentation) {
				pattern.digit = 5
			} else if containsPattern(pattern.displayRepresentation, pattern1.displayRepresentation) {
				pattern.digit = 3
			} else {
				pattern.digit = 2
			}
		}
		display.uniqueSignalPatterns[i] = pattern
	}
	return display
}

func setOutputDigits(display Display) Display {
	for i, outputValue := range display.outputValues {
		outputValue.digit = getDigitOfOutputValue(outputValue.displayRepresentation, display.uniqueSignalPatterns)
		display.outputValues[i] = outputValue
	}
	return display
}

func getDigitOfOutputValue(outputValue string, uniqueSignalPatterns []DisplayDigit) int {
	for _, pattern := range uniqueSignalPatterns {
		if pattern.displayRepresentation == outputValue {
			return pattern.digit
		}
	}
	return -1
}

func linesToDisplays(lines []string) (displays []Display) {
	for _, line := range lines {
		patternAndOutput := strings.Split(line, " | ")
		var display Display
		display.uniqueSignalPatterns = strToDisplayDigits(patternAndOutput[0])
		display.outputValues = strToDisplayDigits(patternAndOutput[1])
		displays = append(displays, display)
	}
	return
}

func strToDisplayDigits(input string) (displayDigits []DisplayDigit) {
	values := strings.Split(input, " ")
	for _, value := range values {
		displayDigit := DisplayDigit{-1, SortString(value)}
		displayDigits = append(displayDigits, displayDigit)
	}
	return
}

func count1478(displays []Display) (count int) {
	for _, display := range displays {
		for _, outputValue := range display.outputValues {
			if outputValue.digit == 1 || outputValue.digit == 4 || outputValue.digit == 7 || outputValue.digit == 8 {
				count++
			}
		}
	}
	return
}

func sumOutputValues(displays []Display) (sum int) {
	for _, display := range displays {
		sum += getOutputValuesAsNumber(display.outputValues)
	}
	return
}

func getOutputValuesAsNumber(displayDigits []DisplayDigit) (number int) {
	for _, digit := range displayDigits {
		number = number*10 + digit.digit
	}
	return
}

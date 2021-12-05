package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) (fileLines []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileLines
}

func readNumbersOfFile(path string) (numbers []int) {
	var fileStrings = readFile(path)
	for _, value := range fileStrings {
		number := strToInt(value)
		if number != -1 {
			numbers = append(numbers, number)
		}
	}
	return
}

func strToInt(numberAsText string) int {
	if i, err := strconv.Atoi(numberAsText); err == nil {
		return i
	}
	return -1
}

func strToInts(line string, sep string) (drawnNumbers []int) {
	lineArray := strings.Split(line, sep)
	for _, number := range lineArray {
		if number == "" {
			continue
		}
		drawnNumbers = append(drawnNumbers, strToInt(number))
	}
	return
}

func getMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getEmptyIntArray(length int) (array []int) {
	for i := 0; i <= length; i++ {
		array = append(array, 0)
	}
	return
}

func almostEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

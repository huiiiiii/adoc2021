package main

import (
	"sort"
	"strings"
)

const bingoGridLength int = 5

type Bingo struct {
	grid          [][]int
	isnumberDrawn [bingoGridLength][bingoGridLength]bool
}

func solveExercise4(fileName string) (solution int) {
	lines := readFile(fileName)
	drawnNumbers, bingos := linesToDrawnNumbersAndBingoGrids(lines)
	var winner Bingo
	hasWon := false
	lastDrawnNumber := 0

	for _, number := range drawnNumbers {
		bingos = drawNumber(number, bingos)
		hasWon, winner, _ = hasAnyBingoWon(bingos)
		if hasWon {
			lastDrawnNumber = number
			break
		}
	}
	return calculateResult(winner, lastDrawnNumber)
}

func solveExercise4b(fileName string) (solution int) {
	lines := readFile(fileName)
	drawnNumbers, bingos := linesToDrawnNumbersAndBingoGrids(lines)
	var looser Bingo
	lastDrawnNumber := 0

	for _, number := range drawnNumbers {
		bingos = drawNumber(number, bingos)
		hasWon, _, winnerIndexes := hasAnyBingoWon(bingos)
		if hasWon {
			if len(bingos) == 1 {
				looser = bingos[0]
				lastDrawnNumber = number
				break
			}
			sort.Sort(sort.Reverse(sort.IntSlice(winnerIndexes)))
			for _, winnerIndex := range winnerIndexes {
				bingos = removeFromBingoArray(bingos, winnerIndex)
			}
		}
	}
	return calculateResult(looser, lastDrawnNumber)
}

func linesToDrawnNumbersAndBingoGrids(lines []string) (drawnNumbers []int, bingos []Bingo) {
	drawnNumbers = stringToNumbers(lines[0], ",")
	for i := 2 + bingoGridLength; i <= len(lines); i += bingoGridLength + 1 {
		linesForBingo := lines[i-bingoGridLength : i]
		bingos = append(bingos, linesToBingo(linesForBingo))
	}
	return
}

func stringToNumbers(line string, sep string) (drawnNumbers []int) {
	lineArray := strings.Split(line, sep)
	for _, number := range lineArray {
		if number == "" {
			continue
		}
		drawnNumbers = append(drawnNumbers, strToInt(number))
	}
	return
}

func linesToBingo(lines []string) (bingo Bingo) {
	for _, line := range lines {
		lineNumbers := stringToNumbers(line, " ")
		bingo.grid = append(bingo.grid, lineNumbers)
	}
	return
}

func drawNumber(number int, bingos []Bingo) []Bingo {
	for i, bingo := range bingos {
		bingos[i] = setNumberForBingo(number, bingo)
	}
	return bingos
}

func setNumberForBingo(number int, bingo Bingo) Bingo {
	for x := 0; x < bingoGridLength; x++ {
		for y := 0; y < bingoGridLength; y++ {
			if bingo.grid[x][y] == number {
				bingo.isnumberDrawn[x][y] = true
			}
		}
	}
	return bingo
}

func hasAnyBingoWon(bingos []Bingo) (hasWon bool, winner Bingo, index []int) {
	for i, bingo := range bingos {
		if hasBingoWon(bingo) {
			hasWon = true
			winner = bingo
			index = append(index, i)
		}
	}
	return
}

func hasBingoWon(bingo Bingo) (hasWon bool) {
	for x := 0; x < bingoGridLength; x++ {
		thisRowHasWon := true
		thisCollumnHasWon := true
		for y := 0; y < bingoGridLength; y++ {
			if !bingo.isnumberDrawn[x][y] {
				thisRowHasWon = false
			}
			if !bingo.isnumberDrawn[y][x] {
				thisCollumnHasWon = false
			}
		}
		if thisRowHasWon || thisCollumnHasWon {
			return true
		}
	}
	return
}

func calculateResult(winner Bingo, lastDrawnNumber int) (solution int) {
	return sumUnmarkedNumbers(winner) * lastDrawnNumber
}

func sumUnmarkedNumbers(bingo Bingo) (sum int) {
	for x := 0; x < bingoGridLength; x++ {
		for y := 0; y < bingoGridLength; y++ {
			if !bingo.isnumberDrawn[x][y] {
				sum += bingo.grid[x][y]
			}
		}
	}
	return
}

//changes order of array!
func removeFromBingoArray(bingos []Bingo, i int) []Bingo {
	bingos[i] = bingos[len(bingos)-1]
	return bingos[:len(bingos)-1]
}

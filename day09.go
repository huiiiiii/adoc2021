package main

func solveExercise9(fileName string) (solution int) {
	lines := readFile(fileName)
	heightmap := linesToHeightmap(lines)
	solution = calculateSumOfLowestPoints(heightmap)
	return
}

//ToDo
func solveExercise9b(fileName string) (solution int) {
	lines := readFile(fileName)
	return len(lines)
}

func linesToHeightmap(lines []string) (heightmap [][]int) {
	heightmap = append(heightmap, getEmptyIntArray(len(lines[0])+1, 10))
	for _, line := range lines {
		heightmap = append(heightmap, lineToHeights(line))
	}
	heightmap = append(heightmap, getEmptyIntArray(len(lines[0])+1, 10))
	return
}

func lineToHeights(line string) (heights []int) {
	heights = append(heights, 10)
	numbers := strToInts(line, "")
	heights = append(heights, numbers...)
	heights = append(heights, 10)
	return
}

func calculateSumOfLowestPoints(heightmap [][]int) (solution int) {
	for x := 1; x < len(heightmap)-1; x++ {
		for y := 1; y < len(heightmap[x])-1; y++ {
			if heightmap[x][y] < heightmap[x-1][y] && heightmap[x][y] < heightmap[x+1][y] && heightmap[x][y] < heightmap[x][y-1] && heightmap[x][y] < heightmap[x][y+1] {
				solution += heightmap[x][y] + 1
			}
		}
	}
	return
}

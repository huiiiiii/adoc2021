package main

func solveExercise6(fileName string) (solution int) {
	const daysToSimulate int = 80
	lines := readFile(fileName)
	daysToCountFish := getInitialDaysToCountFish(lines[0])
	for i := 0; i < daysToSimulate; i++ {
		daysToCountFish = calculateNewFishSpawnDays(daysToCountFish)
	}
	solution = countAllFish(daysToCountFish)
	return
}

//ToDo needs improvement, runs too long
func solveExercise6b(fileName string) (solution int) {
	const daysToSimulate int = 256
	lines := readFile(fileName)
	daysToCountFish := getInitialDaysToCountFish(lines[0])
	for i := 0; i < daysToSimulate; i++ {
		daysToCountFish = calculateNewFishSpawnDays(daysToCountFish)
	}
	solution = countAllFish(daysToCountFish)
	return
}

func getInitialDaysToCountFish(input string) map[int]int {
	fishSpawnDays := strToInts(input, ",")
	daysToCountFish := getEmptyMap()
	for _, fish := range fishSpawnDays {
		daysToCountFish[fish] += 1
	}
	return daysToCountFish
}

func getEmptyMap() map[int]int {
	return map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0}
}

func calculateNewFishSpawnDays(fishSpawnDays map[int]int) (newFishSpawnDays map[int]int) {
	newFishSpawnDays = getEmptyMap()
	for days, countFish := range fishSpawnDays {
		if days == 0 {
			newFishSpawnDays[6] += countFish
			newFishSpawnDays[8] += countFish
		} else {
			newFishSpawnDays[days-1] += countFish
		}
	}
	return
}

func countAllFish(daysToCountFish map[int]int) (count int) {
	for _, countFish := range daysToCountFish {
		count += countFish
	}
	return
}

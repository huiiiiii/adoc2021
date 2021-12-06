package main

func solveExercise6(fileName string) int {
	lines := readFile(fileName)
	fishSpawnDays := strToInts(lines[0], ",")
	const daysToSimulate int = 80
	for i := 0; i < daysToSimulate; i++ {
		fishSpawnDays = calculateNewFishSpawnDays(fishSpawnDays)
	}
	return len(fishSpawnDays)
}

//ToDo needs improvement, runs too long
func solveExercise6b(fileName string) int {
	lines := readFile(fileName)
	fishSpawnDays := strToInts(lines[0], ",")
	const daysToSimulate int = 256
	for i := 0; i < daysToSimulate; i++ {
		fishSpawnDays = calculateNewFishSpawnDays(fishSpawnDays)
	}
	return len(fishSpawnDays)
}

func calculateNewFishSpawnDays(fishSpawnDays []int) (newFishSpawnDays []int) {
	for _, fish := range fishSpawnDays {
		if fish == 0 {
			newFishSpawnDays = append(newFishSpawnDays, 6)
			newFishSpawnDays = append(newFishSpawnDays, 8)
		} else {
			fish--
			newFishSpawnDays = append(newFishSpawnDays, fish)
		}
	}
	return
}

package main

func solveExercise1() (solution int) {
	numbers := readNumbersOfFile("file1.txt")
	numberSlice := numbers[1:]
	for index, number := range numberSlice {
		difference := numbers[index] - number
		if difference < 0 {
			solution++
		}
	}
	return
}

func solveExercise1b() (solution int) {
	numbers := readNumbersOfFile("file1.txt")
	numberSlice := numbers[3:]
	for index, number := range numberSlice {
		a := numbers[index] + numbers[index+1] + numbers[index+2]
		b := numbers[index+1] + numbers[index+2] + number
		difference := a - b

		if difference < 0 {
			solution++
		}
	}
	return
}

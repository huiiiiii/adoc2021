package main

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"strconv"
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

func callFunction(funcName interface{}, params ...interface{}) {
	f := reflect.ValueOf(funcName)
	if len(params) != f.Type().NumIn() {
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	f.Call(in)
}

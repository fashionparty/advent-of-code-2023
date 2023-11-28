package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintAnswers(dayNumber int, answer1 string, answer2 string) {
	fmt.Println("\nDay", dayNumber)
	fmt.Println("Answer 1:", answer1)
	fmt.Println("Answer 2:", answer2)
}

func ReadInput(path string) []string {
	input, err := os.Open(path)
	LogError(err)
	scanner := bufio.NewScanner(input)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Filter[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for i := 0; i < len(collection)-1; i++ {
		if predicate(collection[i]) {
			result = append(result, collection[i])
		}
	}
	return result
}

func Update[T any](collection []T, transformation func(T) T) []T {
	var result []T
	for i := 0; i < len(collection)-1; i++ {
		result = append(result, transformation(collection[i]))
	}
	return result
}

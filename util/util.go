package util

import (
	"bufio"
	"errors"
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

func WindowString(input string, windowSize int) []string {
	result := []string{}
	loopLimit := len(input) - windowSize + 1
	for i := 0; i < loopLimit; i++ {
		result = append(result, input[i:i+windowSize])
	}
	return result
}

func FindSubstringIndex(input string, substring string) (int, error) {
	loopLimit := len(input) - len(substring) + 1
	for i := 0; i < loopLimit; i++ {
		if substring == input[i:i+len(substring)] {
			return i, nil
		}
	}
	return 0, errors.New("substring not found")
}

func FindBiggestNumber[T numbers](input []T) T {
	var result T = input[0]
	for _, v := range input {
		if v > result {
			result = v
		}
	}
	return result
}

func FindSmallestNumber[T numbers](input []T) T {
	var result T = input[0]
	{
		for _, v := range input {
			if v < result {
				result = v
			}
		}
	}
	return result
}

func FindBiggestUint(values []uint) uint {
	var result uint
	for _, v := range values {
		if v > result {
			result = v
		}
	}
	return result
}

type numbers interface {
	int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		uint | uint8 | uint16 | uint32 | uint64
}

var Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func StringContains(input string, collection []string) bool {
	for _, v := range collection {
		if v == input {
			return true
		}
	}
	return false
}

package day_1

import (
	"advent-of-code-2023/util"
	"strconv"
)

type digitWithIndex struct {
	value string
	index int
}

var spelledDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func Run() {
	input := util.ReadInput("day_1\\input.txt")

	answer1 := sumCalibrationValuesPart1(input)
	answer2 := sumCalibrationValuesPart2(input)

	util.PrintAnswers(1, answer1, answer2)
}

func findFirstDigit(input string) digitWithIndex {
	result := ""
	resultIndex := 0
	for i := 0; i < len(input); i++ {
		for k := 0; k < 10; k++ {
			digit := strconv.Itoa(k)
			if digit == string(input[i]) {
				result = string(input[i])
				resultIndex = i
			}
		}
		if result != "" {
			break
		}
	}
	return digitWithIndex{result, resultIndex}
}

func findLastDigit(input string) digitWithIndex {
	result := ""
	resultIndex := 0
	for i := len(input) - 1; i >= 0; i-- {
		for k := 0; k < 10; k++ {
			digit := strconv.Itoa(k)
			if digit == string(input[i]) {
				result = string(input[i])
				resultIndex = i
			}
		}
		if result != "" {
			break
		}
	}
	return digitWithIndex{result, resultIndex}
}

func getCalibrationValue(input string) int {
	firstDigit := findFirstDigit(input)
	secondDigit := findLastDigit(input)
	merged := firstDigit.value + secondDigit.value
	result, err := strconv.Atoi(merged)
	util.LogError(err)
	return result
}

func sumCalibrationValuesPart1(input []string) string {
	result := 0
	for _, v := range input {
		result += getCalibrationValue(v)
	}
	return strconv.Itoa(result)
}

func parseDigit(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}

func findFirstSpelledDigit(input string) digitWithIndex {
	var windowedDigits []digitWithIndex
	for _, spelledDigit := range spelledDigits {
		windows := util.WindowString(input, len(spelledDigit))
		for i, window := range windows {
			if spelledDigit == window {
				windowedDigits = append(windowedDigits, digitWithIndex{
					value: parseDigit(spelledDigit),
					index: i,
				})
			}
		}
	}
	if len(windowedDigits) == 0 {
		return digitWithIndex{"", 999}
	}
	leftmostDigit := windowedDigits[0]
	for _, v := range windowedDigits {
		if v.index <= leftmostDigit.index {
			leftmostDigit = v
		}
	}
	return leftmostDigit
}

func findLastSpelledDigit(input string) digitWithIndex {
	var windowedDigits []digitWithIndex
	for _, spelledDigit := range spelledDigits {
		windows := util.WindowString(input, len(spelledDigit))
		for i, window := range windows {
			if spelledDigit == window {
				windowedDigits = append(windowedDigits, digitWithIndex{
					value: parseDigit(spelledDigit),
					index: i,
				})
			}
		}
	}
	if len(windowedDigits) == 0 {
		return digitWithIndex{"", -999}
	}
	rightmostDigit := windowedDigits[0]
	for _, v := range windowedDigits {
		if v.index >= rightmostDigit.index {
			rightmostDigit = v
		}
	}
	return rightmostDigit
}

func getNewCalibrationValue(input string) int {
	firstDigit := findFirstDigit(input)
	firstSpelledDigit := findFirstSpelledDigit(input)
	secondDigit := findLastDigit(input)
	secondSpelledDigit := findLastSpelledDigit(input)

	firstDigitResult := ""

	if firstDigit.index < firstSpelledDigit.index {
		firstDigitResult = firstDigit.value
	} else {
		firstDigitResult = firstSpelledDigit.value
	}

	secondDigitResult := ""

	if secondDigit.index > secondSpelledDigit.index {
		secondDigitResult = secondDigit.value
	} else {
		secondDigitResult = secondSpelledDigit.value
	}

	merged := firstDigitResult + secondDigitResult

	result, err := strconv.Atoi(merged)
	util.LogError(err)

	return result
}

func sumCalibrationValuesPart2(input []string) string {
	result := 0
	for _, v := range input {
		result += getNewCalibrationValue(v)
	}
	return strconv.Itoa(result)
}

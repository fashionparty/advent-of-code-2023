package day_3

import (
	"advent-of-code-2023/util"
	"strconv"
)

type number struct {
	value        int
	length       int
	line         int
	index        int
	isPartNumber bool
}

// todo works only on example input :(
func Run() {
	input := util.ReadInput("day_3\\input.txt")

	numbers := parseNumbers(input)
	partNumbers := determinePartNumbers(input, numbers)
	sum := strconv.Itoa(sumPartNumbers(partNumbers))

	util.PrintAnswers(3, sum, "n/a")
}

func parseNumbers(input []string) []number {
	var result []number
	for i, v := range input {
		numbersFromLine := findNumbersInLine(v, i)
		for _, k := range numbersFromLine {
			result = append(result, k)
		}
	}
	return result
}

func findNumbersInLine(input string, lineNumber int) []number {
	var result []number
	flagIsDigit := false
	currentNumber := ""
	for i, v := range input {
		if util.StringContains(string(v), util.Digits) {
			flagIsDigit = true
			currentNumber += string(v)
		} else {
			if flagIsDigit {
				newNumberValue, err := strconv.Atoi(currentNumber)
				util.LogError(err)
				newNumber := number{
					value:  newNumberValue,
					length: len(currentNumber),
					line:   lineNumber,
					index:  i - len(currentNumber),
				}
				result = append(result, newNumber)
				currentNumber = ""
			}
			flagIsDigit = false
		}
		if currentNumber != "" {
			newNumberValue, err := strconv.Atoi(currentNumber)
			util.LogError(err)
			newNumber := number{
				value:  newNumberValue,
				length: len(currentNumber),
				line:   lineNumber,
				index:  i - len(currentNumber),
			}
			result = append(result, newNumber)
		}

	}

	return result
}

func determinePartNumbers(input []string, numbers []number) []number {
	result := numbers
	for i, v := range result {
		isPartNumber := false
		hasSpaceLeft := v.index > 0
		hasSpaceRight := v.index+v.length < len(input[v.line])
		hasSpaceAbove := v.line > 0
		hasSpaceBelow := v.line < len(input)-1
		//check left side
		if hasSpaceLeft {
			if input[v.line][v.index-1] != '.' {
				isPartNumber = true
			}
		}
		//check right side
		if hasSpaceRight {
			if input[v.line][v.index+v.length] != '.' {
				isPartNumber = true
			}
		}
		if hasSpaceAbove {
			//upper left corner
			if hasSpaceLeft {
				if input[v.line-1][v.index-1] != '.' {
					isPartNumber = true
				}
			}
			//upper right corner
			if hasSpaceRight {
				if input[v.line-1][v.index+v.length] != '.' {
					isPartNumber = true
				}
			}
			//above
			for i := 0; i < v.length; i++ {
				if input[v.line-1][v.index+i] != '.' {
					isPartNumber = true
				}
			}
		}

		if hasSpaceBelow {
			//bottom left corner
			if hasSpaceLeft {
				if input[v.line+1][v.index-1] != '.' {
					isPartNumber = true
				}
			}
			//bottom right corner
			if hasSpaceRight {
				if input[v.line+1][v.index+v.length] != '.' {
					isPartNumber = true
				}
			}
			//below
			for i := 0; i < v.length; i++ {
				if input[v.line+1][v.index+i] != '.' {
					isPartNumber = true
				}
			}
		}
		result[i].isPartNumber = isPartNumber
	}
	return result
}

func sumPartNumbers(numbers []number) int {
	result := 0
	for _, v := range numbers {
		if v.isPartNumber {
			result += v.value
		}
	}
	return result
}

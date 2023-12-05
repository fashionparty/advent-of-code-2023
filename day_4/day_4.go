package day_4

import "advent-of-code-2023/util"

func Run() {
	//input := util.ReadInput("day_4\\input.txt")
	util.PrintAnswers(4, "n/a", "n/a")
}

func separateLine(input string) (string, string) {
	dividerIndex, err := util.FindSubstringIndex(input, "|")
	util.LogError(err)
	semicolonIndex, err := util.FindSubstringIndex(input, ":")
	util.LogError(err)
	leftSide := input[semicolonIndex+2 : dividerIndex-1]
	rightSide := input[dividerIndex+2:]
	return leftSide, rightSide
}

func parseNumbers(input string) []string {
	currentNumber := ""
	var result []string

	for _, v := range input {
		if util.StringContains(string(v), util.Digits) {
			currentNumber += string(v)
		} else {
			result = append(result, currentNumber)
			currentNumber = ""
		}
		//todo doda pusty string przy dwoch spacjach
		//nie zadziała prawidłowo jeśli cyfra jest na końcu linii
	}
	return result
}

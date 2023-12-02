package day_2

import (
	"advent-of-code-2023/util"
	"slices"
	"strconv"
)

var redCubesLimit uint = 12
var greenCubesLimit uint = 13
var blueCubesLimit uint = 14

type game struct {
	gameId   uint
	maxRed   uint
	maxGreen uint
	maxBlue  uint
}

func Run() {
	input := util.ReadInput("day_2\\input.txt")
	games := parseGames(input)
	possibleGames := strconv.Itoa(int(countPossibleGamesIdSum(games)))
	sumOfPowers := strconv.Itoa(int(countSumOfPowers(games)))
	util.PrintAnswers(2, possibleGames, sumOfPowers)
}

func parseGames(input []string) []game {
	var result []game
	for _, v := range input {
		id := getGameId(v)
		maxRed := getGameMaxColor(v, "red")
		maxGreen := getGameMaxColor(v, "green")
		maxBlue := getGameMaxColor(v, "blue")
		game := game{
			gameId:   id,
			maxRed:   maxRed,
			maxGreen: maxGreen,
			maxBlue:  maxBlue,
		}
		result = append(result, game)
	}
	return result
}

func getGameId(input string) uint {
	indexStart, err := util.FindSubstringIndex(input, " ")
	util.LogError(err)
	indexEnd, err := util.FindSubstringIndex(input, ":")
	util.LogError(err)
	result, err := strconv.Atoi(input[indexStart+1 : indexEnd])
	util.LogError(err)
	return uint(result)
}

func getGameMaxColor(input string, color string) uint {
	var colorValues []uint
	movingIndex := 0
	for {
		currentIndex, err := util.FindSubstringIndex(input[movingIndex:], color)
		currentIndex += movingIndex
		if err != nil {
			break
		} else {
			var digits []string
			movingSubIndex := currentIndex - 2
			digits = append(digits, input[movingSubIndex:movingSubIndex+1])
			for {
				if movingSubIndex > 0 {
					movingSubIndex--
				} else {
					break
				}
				if string(input[movingSubIndex]) == " " {
					break
				} else {
					digits = append(digits, input[movingSubIndex:movingSubIndex+1])
				}
			}
			slices.Reverse(digits)
			digitsMerged := ""
			for _, v := range digits {
				digitsMerged += v
			}
			colorValue, err := strconv.Atoi(digitsMerged)
			util.LogError(err)
			colorValues = append(colorValues, uint(colorValue))
			movingIndex = currentIndex + 1
		}
	}
	return util.FindBiggestUint(colorValues)
}

func countPossibleGamesIdSum(games []game) uint {
	var result uint
	for _, v := range games {
		if v.maxRed <= redCubesLimit && v.maxGreen <= greenCubesLimit && v.maxBlue <= blueCubesLimit {
			result += v.gameId
		}
	}
	return result
}

func countSumOfPowers(input []game) uint {
	var result uint
	for _, v := range input {
		result += v.maxRed * v.maxGreen * v.maxBlue
	}
	return result
}

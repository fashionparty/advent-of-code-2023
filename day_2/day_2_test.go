package day_2

import (
	"strconv"
	"testing"
)

var exampleInput = []string{
	"game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

var firstExampleAnswer uint = 8

var exampleGames = []game{
	{1, 4, 2, 6},
	{2, 1, 3, 4},
	{3, 20, 13, 6},
	{4, 14, 3, 15},
	{5, 6, 3, 2},
}

var secondExampleAnswer = "2286"

func TestParseGames(t *testing.T) {
	input := exampleInput
	want := exampleGames
	got := parseGames(input)
	for i, v := range got {
		if v != want[i] {
			t.Errorf("parseGames = %v, expected = %v", v, want[i])
		}
	}
}

func TestGetGameId(t *testing.T) {
	input := exampleInput
	wantCollection := []uint{1, 2, 3, 4, 5}
	for i, v := range input {
		got := getGameId(v)
		want := wantCollection[i]
		if got != want {
			t.Errorf("getGameId = %v, expected = %v", got, want)
		}
	}
}

func TestGetGameMaxColor(t *testing.T) {
	input := exampleInput
	wantCollection := []uint{4, 1, 20, 14, 6}
	for i, v := range input {
		got := getGameMaxColor(v, "red")
		want := wantCollection[i]
		if got != want {
			t.Errorf("getGameMaxColor = %v, expected = %v", got, want)
		}
	}
}

func TestCountPossibleGamesIdSum(t *testing.T) {
	input := exampleGames
	var want = firstExampleAnswer
	got := countPossibleGamesIdSum(input)
	if want != got {
		t.Errorf("countPossibleGamesIdSum = %v, expected = %v", got, want)
	}
}

func TestCountSumOfPowers(t *testing.T) {
	input := exampleGames
	want := secondExampleAnswer
	got := strconv.Itoa(int(countSumOfPowers(input)))
	if got != want {
		t.Errorf("countSumOfPowers = %v, expected = %v", got, want)
	}
}

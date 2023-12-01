package day_1

import "testing"

var firstExampleInput = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var firstExampleAnswer = "142"

var secondExampleInput = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

var secondExampleAnswer = "281"

func TestFindFirstDigit(t *testing.T) {
	input := firstExampleInput
	wantCollection := []digitWithIndex{
		{"1", 0},
		{"3", 3},
		{"1", 1},
		{"7", 4},
	}
	for i, v := range input {
		want := wantCollection[i]
		got := findFirstDigit(v)
		if got != want {
			t.Errorf("findFirstDigit = %v, expected = %v", got, want)
		}
	}
}

func TestFindLastDigit(t *testing.T) {
	input := firstExampleInput
	wantCollection := []digitWithIndex{
		{"2", 4},
		{"8", 7},
		{"5", 9},
		{"7", 4},
	}
	for i, v := range input {
		want := wantCollection[i]
		got := findLastDigit(v)
		if got != want {
			t.Errorf("findLastDigit = %v, expected = %v", got, want)
		}
	}
}

func TestGetCalibrationValue(t *testing.T) {
	input := firstExampleInput
	wantCollection := []int{12, 38, 15, 77}
	for i, v := range input {
		want := wantCollection[i]
		got := getCalibrationValue(v)
		if got != want {
			t.Errorf("getCalibrationValue = %v, expected = %v", got, want)
		}
	}
}

func TestSumCalibrationValuesPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	want := firstExampleAnswer
	got := sumCalibrationValuesPart1(input)
	if got != want {
		t.Errorf("sumCalibrationValuesPart1 = %v, expected = %v", got, want)
	}
}

func TestParseDigit(t *testing.T) {
	input := spelledDigits
	wantCollection := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	for i, v := range input {
		want := wantCollection[i]
		got := parseDigit(v)
		if got != want {
			t.Errorf("parseDigits = %v, expected = %v", got, want)
		}
	}
}

func TestFindFirstSpelledDigit(t *testing.T) {
	input := secondExampleInput
	wantCollection := []digitWithIndex{
		{"2", 0},
		{"8", 0},
		{"1", 3},
		{"2", 1},
		{"9", 1},
		{"1", 1},
		{"6", 6},
	}
	for i, v := range input {
		wanted := wantCollection[i]
		got := findFirstSpelledDigit(v)
		if got != wanted {
			t.Errorf("findFirstSpelledDigit = %v, expected %v", got, wanted)
		}
	}
}

func TestFindLastSpelledDigit(t *testing.T) {
	input := secondExampleInput
	wantCollection := []digitWithIndex{
		{"9", 4},
		{"3", 7},
		{"3", 7},
		{"4", 7},
		{"7", 10},
		{"8", 3},
		{"6", 6},
	}
	for i, v := range input {
		want := wantCollection[i]
		got := findLastSpelledDigit(v)
		if want != got {
			t.Errorf("findLastSpelledDigit = %v, expected = %v", got, want)
		}
	}
}

func TestGetNewCalibrationValue(t *testing.T) {
	input := secondExampleInput
	wantCollection := []int{
		29, 83, 13, 24, 42, 14, 76,
	}
	for i, v := range input {
		want := wantCollection[i]
		got := getNewCalibrationValue(v)
		if got != want {
			t.Errorf("getNewCalibrationValue = %v, expected = %v", got, want)
		}
	}
}

func TestSumCalibrationValuesPart2(t *testing.T) {
	input := secondExampleInput
	want := secondExampleAnswer
	got := sumCalibrationValuesPart2(input)
	if got != want {
		t.Errorf("sumCalibrationValuesPart2 = %v, expected %v", got, want)
	}
}

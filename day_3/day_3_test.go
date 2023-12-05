package day_3

import "testing"

var firstExampleInput = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

var numbers = []number{
	{value: 467, length: 3, line: 0, index: 0},
	{value: 114, length: 3, line: 0, index: 5},
	{value: 35, length: 2, line: 2, index: 2},
	{value: 633, length: 3, line: 2, index: 6},
	{value: 617, length: 3, line: 4, index: 0},
	{value: 58, length: 2, line: 5, index: 7},
	{value: 592, length: 3, line: 6, index: 2},
	{value: 755, length: 3, line: 7, index: 6},
	{value: 664, length: 3, line: 9, index: 1},
	{value: 598, length: 3, line: 9, index: 5},
}

var numbersChecked = []number{
	{value: 467, length: 3, line: 0, index: 0, isPartNumber: true},
	{value: 114, length: 3, line: 0, index: 5, isPartNumber: false},
	{value: 35, length: 2, line: 2, index: 2, isPartNumber: true},
	{value: 633, length: 3, line: 2, index: 6, isPartNumber: true},
	{value: 617, length: 3, line: 4, index: 0, isPartNumber: true},
	{value: 58, length: 2, line: 5, index: 7, isPartNumber: false},
	{value: 592, length: 3, line: 6, index: 2, isPartNumber: true},
	{value: 755, length: 3, line: 7, index: 6, isPartNumber: true},
	{value: 664, length: 3, line: 9, index: 1, isPartNumber: true},
	{value: 598, length: 3, line: 9, index: 5, isPartNumber: true},
}

func TestSumPartNumbers(t *testing.T) {
	input := numbersChecked
	want := 4361
	got := sumPartNumbers(input)
	if got != want {
		t.Errorf("sumPartNumbers = %v, expected := %v", got, want)
	}
}

func TestFindNumbersInLine(t *testing.T) {
	input := firstExampleInput[0]
	want := []number{numbers[0], numbers[1]}
	got := findNumbersInLine(input, 0)
	for i, v := range want {
		if v != got[i] {
			t.Errorf("findNumbersInLine = %v, expected = %v", got, want)
		}
	}
}

func TestParseNumbers(t *testing.T) {
	input := firstExampleInput
	want := numbers
	got := parseNumbers(input)
	for i, v := range want {
		if v != got[i] {
			t.Errorf("parseNumbers = %v, expected = %v", got, want)
		}
	}
}

func TestDeterminePartNumbers(t *testing.T) {
	input_1 := firstExampleInput
	input_2 := numbers
	want := numbersChecked
	got := determinePartNumbers(input_1, input_2)

	for i, v := range want {
		if v != got[i] {
			t.Errorf("determinePartNumbers = %v, expected = %v", got, want)
		}
	}

}

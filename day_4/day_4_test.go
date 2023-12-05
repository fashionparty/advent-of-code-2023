package day_4

import "testing"

var firstExampleInput = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestSeparateLine(t *testing.T) {
	input := firstExampleInput[0]
	gotLeft, gotRight := separateLine(input)
	wantLeft := "41 48 83 86 17"
	wantRight := "83 86  6 31 17  9 48 53"

	if wantLeft != gotLeft {
		t.Errorf("separateLine left = %v, expected left = %v", gotLeft, wantLeft)
	}
	if wantRight != gotRight {
		t.Errorf("separate right = %v, expected right = %v", gotRight, wantRight)
	}
}

func TestParseNumbers(t *testing.T) {
	input := "41 48 83 86 17"
	want := []string{"41", "48", "83", "86", "17"}
	got := parseNumbers(input)
	for i, v := range want {
		if v != got[i] {
			t.Errorf("parseNumbers = %v, expected = %v", got[i], want)
		}
	}
}

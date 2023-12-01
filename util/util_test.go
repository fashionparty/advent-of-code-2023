package util

import "testing"

func TestWindowString(t *testing.T) {
	input := "abcdefg"
	want := []string{"abc", "bcd", "cde", "def", "efg"}
	got := WindowString(input, 3)
	for i := 0; i < len(want)-1; i++ {
		if got[i] != want[i] {
			t.Errorf("WindowString = %v, expected = %v", got, want)
		}
	}
}

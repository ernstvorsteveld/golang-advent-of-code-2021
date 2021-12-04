package main

import "testing"

func TestExample1a(t *testing.T) {
	const expected = 7
	var actual = Challenge1a([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestChallenge1a(t *testing.T) {
	const expected = 1696
	var actual = Challenge1aFile("puzzle-input/1a-depth-measurements.txt")
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}

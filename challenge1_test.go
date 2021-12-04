package main

import "testing"

func TestExample1(t *testing.T) {
	const expected = 7
	var actual = Challenge1([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestChallenge1(t *testing.T) {
	const expected = 1696
	var actual = Challenge1File("puzzle-input/puzzle-input-depth-2-measurements.txt")
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}

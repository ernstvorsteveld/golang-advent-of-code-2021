package main

import "testing"

func TestExample2(t *testing.T) {
	const expected = 5
	var actual = Challenge2([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}


func TestChallenge2(t *testing.T) {
	const expected = 1737
	var actual = Challenge2File("2/2-measurements.txt")
	if actual != expected {
		t.Errorf("Number larger than previous incorrect, got: %d, want: %d.", actual, expected)
	}
}

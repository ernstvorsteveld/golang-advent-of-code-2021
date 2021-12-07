package main

import "testing"

func TestExample4aSimple(t *testing.T) {
	actualFinalScore := Challenge4aSimple("puzzle-input/4a-simple.txt", 5)
	expectedFinalScore := 4512

	if actualFinalScore != expectedFinalScore {
		t.Errorf("Epsilon Rate incorrect, got: %d, want: %d.", actualFinalScore, expectedFinalScore)
	}
}

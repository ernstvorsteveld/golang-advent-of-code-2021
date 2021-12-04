package main

import "testing"

func TestExample2a(t *testing.T) {
	const expectedHorizontal = 15
	const expectedDepth = 10
	var actualHorizontal, actualDepth = Challenge2aFile("puzzle-input/3a-simple.txt")
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Horizontal position not incorrect, got: %d, want: %d.", actualHorizontal, expectedHorizontal)
	}
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Depth not incorrect, got: %d, want: %d.", actualDepth, expectedDepth)
	}
}

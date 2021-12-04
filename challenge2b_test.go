package main

import (
	"fmt"
	"testing"
)

func TestExample2bChallenge(t *testing.T) {
	const expectedHorizontal = 2050
	const expectedDepth = 826
	var actualHorizontal, actualDepth = Challenge2bFile("puzzle-input/2b-challenge.txt")
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Horizontal position not incorrect, got: %d, want: %d.", actualHorizontal, expectedHorizontal)
	}
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Depth not incorrect, got: %d, want: %d.", actualDepth, expectedDepth)
	}
	fmt.Println(actualHorizontal * actualDepth)
}

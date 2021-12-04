package main

import (
	"fmt"
	"testing"
)

func TestExample2aSimple(t *testing.T) {
	const expectedHorizontal = 15
	const expectedDepth = 10
	var actualHorizontal, actualDepth = Challenge2aFile("puzzle-input/2a-simple.txt")
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Horizontal position not incorrect, got: %d, want: %d.", actualHorizontal, expectedHorizontal)
	}
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Depth not incorrect, got: %d, want: %d.", actualDepth, expectedDepth)
	}
}

func TestExample2aChallenge(t *testing.T) {
	const expectedHorizontal = 2050
	const expectedDepth = 826
	var actualHorizontal, actualDepth = Challenge2aFile("puzzle-input/2a-challenge.txt")
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Horizontal position not incorrect, got: %d, want: %d.", actualHorizontal, expectedHorizontal)
	}
	if actualHorizontal != expectedHorizontal {
		t.Errorf("Depth not incorrect, got: %d, want: %d.", actualDepth, expectedDepth)
	}
	fmt.Println(expectedHorizontal * expectedDepth)
}

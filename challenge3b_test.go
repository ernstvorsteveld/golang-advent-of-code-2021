package main

import (
	"fmt"
	"testing"
)

func TestExample3bSimple(t *testing.T) {
	expectedOxygenGeneratorRating := int64(23)
	expectedCO2ScrubberRating := int64(10)
	actualOxygenGeneratorRating,actualCO2ScrubberRating := Challenge3bFile("puzzle-input/3a-simple.txt")
	if actualOxygenGeneratorRating != expectedOxygenGeneratorRating {
		t.Errorf("Gamma Rate incorrect, got: %d, want: %d.", actualOxygenGeneratorRating, expectedOxygenGeneratorRating)
	}
	if actualCO2ScrubberRating != expectedCO2ScrubberRating {
		t.Errorf("Epsilon Rate incorrect, got: %d, want: %d.", actualCO2ScrubberRating, expectedCO2ScrubberRating)
	}
	fmt.Println(actualCO2ScrubberRating * actualOxygenGeneratorRating)
}

func TestExample3bChallenge(t *testing.T) {
	expectedOxygenGeneratorRating := int64(2781)
	expectedCO2ScrubberRating := int64(919)
	actualOxygenGeneratorRating,actualCO2ScrubberRating := Challenge3bFile("puzzle-input/3b-challenge.txt")
	if actualOxygenGeneratorRating != expectedOxygenGeneratorRating {
		t.Errorf("Oxygen Generator Rating incorrect, got: %d, want: %d.", actualOxygenGeneratorRating, expectedOxygenGeneratorRating)
	}
	if actualCO2ScrubberRating != expectedCO2ScrubberRating {
		t.Errorf("CO2 Rubber Rating	 incorrect, got: %d, want: %d.", actualCO2ScrubberRating, expectedCO2ScrubberRating)
	}
	fmt.Println(actualCO2ScrubberRating * actualOxygenGeneratorRating)
}


package main

import (
	"fmt"
	"testing"
)

func TestExample3aSimple(t *testing.T) {
	expectedGammaRate := int64(22)
	expectedEpsilonRate := int64(9)
	actualGammaRate, actualEpsilonRate := Challenge3aFile("puzzle-input/3a-simple.txt")
	if actualGammaRate != expectedGammaRate {
		t.Errorf("Gamma Rate incorrect, got: %d, want: %d.", actualGammaRate, expectedGammaRate)
	}
	if actualGammaRate != expectedGammaRate {
		t.Errorf("Epsilon Rate incorrect, got: %d, want: %d.", actualEpsilonRate, expectedEpsilonRate)
	}
	fmt.Println(actualEpsilonRate * actualGammaRate)
}

func TestExample3aChallenge(t *testing.T) {
	expectedGammaRate := int64(2502)
	expectedEpsilonRate := int64(1593)
	actualGammaRate, actualEpsilonRate := Challenge3aFile("puzzle-input/3a-challenge.txt")
	if actualGammaRate != expectedGammaRate {
		t.Errorf("Gamma Rate incorrect, got: %d, want: %d.", actualGammaRate, expectedGammaRate)
	}
	if actualGammaRate != expectedGammaRate {
		t.Errorf("Epsilon Rate incorrect, got: %d, want: %d.", actualEpsilonRate, expectedEpsilonRate)
	}
	fmt.Println(actualEpsilonRate * actualGammaRate)
	expectedConsumption := int(3985686)
	actualConsumption := int(actualEpsilonRate * actualGammaRate)
	if actualConsumption != expectedConsumption {
		t.Errorf("Consumption incorrect, got: %d, want: %d.", actualConsumption, expectedConsumption)
	}
}

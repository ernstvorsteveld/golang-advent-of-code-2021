package main

// How many measurements are larger than the previous measurement?

func Challenge1aFile(s string) int {
	var depths = StringToIn(ReadMeasurements(s))
	return Challenge1a(depths)
}

func Challenge1a(depths []int) int {
	var numLarger = 0
	var previous = depths[0]
	for i := 0; i < len(depths); i++ {
		if previous < depths[i] {
			numLarger++
		}
		previous = depths[i]
	}
	return numLarger
}

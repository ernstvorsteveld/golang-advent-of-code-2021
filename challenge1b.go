package main

func Challenge1bFile(s string) int {
	var depths = StringToIn(ReadMeasurements(s))
	return Challenge1b(depths)
}

func Challenge1b(depths []int) int {
	var numLarger = 0
	var previous = getMultiMeasure(depths, 3, 0)
	for i := 0; i < len(depths); i++ {
		var next = getMultiMeasure(depths, 3, i)
		if previous < next {
			numLarger++
		}
		previous = next
	}
	return numLarger
}

func getMultiMeasure(d []int, length int, start int) int {
	var value = 0
	for i := 0; i < length; i++ {
		if i + start < len(d) {
			value += d[i + start]
		} else {
			return -1
		}
	}
	return value
}

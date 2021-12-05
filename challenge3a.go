package main

import "strconv"

func Challenge3aFile(s string) (int64, int64) {
	var commands = ReadMeasurements(s)
	return Challenge3a(commands)
}

func Challenge3a(commands []string) (int64, int64) {
	length := len(commands[0])
	numberOnes := make([]int, length)
	for i := 0; i < length; i++ {
		numberOnes[i] = Count(commands, i, "1")
	}
	g := Rate(numberOnes, len(commands), "1", "0")
	e := Rate(numberOnes, len(commands), "0", "1")
	return g, e
}

func Count(commands []string, p int, ch string) int {
	numberOnes := 0
	for i := 0; i < len(commands); i++ {
		c := commands[i][p]
		if string(c) == ch {
			numberOnes++
		}
	}
	return numberOnes
}

func Rate(n []int, length int, posVal string, negVal string) int64 {
	count := len(n)
	result := ""
	for i := 0; i < count; i++ {
		if n[i] > length/2 {
			result += posVal
		} else {
			result += negVal
		}
	}
	rate, _ := strconv.ParseInt(result, 2, 64)
	return rate
}

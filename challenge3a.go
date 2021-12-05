package main

import "strconv"

func Challenge3aFile(s string) (int64,int64) {
	var commands = ReadMeasurements(s)
	return Challenge3a(commands)
}

func Challenge3a(commands []string) (int64, int64) {
	length := len(commands[0])
	numberOnes := make([]int, length)
	for i := 0; i < len(commands); i++ {
		for j, c := range commands[i] {
			if c == '1' {
				numberOnes[j]++
			}
		}
	}
	g := GammaRate(numberOnes, len(commands))
	e := EpsilonRate(numberOnes, len(commands))
	return g, e
}

func GammaRate(n []int, length int) int64 {
	count := len(n)
	result := ""
	for i := 0; i < count; i++ {
		if n[i] > length/2 {
			result += "1"
		} else {
			result += "0"
		}
	}
	gammaRate, _ := strconv.ParseInt(result, 2, 64)
	return gammaRate
}

func EpsilonRate(n []int, length int) int64 {
	count := len(n)
	result := ""
	for i := 0; i < count; i++ {
		if n[i] > length/2 {
			result += "0"
		} else {
			result += "1"
		}
	}
	gammaRate, _ := strconv.ParseInt(result, 2, 64)
	return gammaRate
}
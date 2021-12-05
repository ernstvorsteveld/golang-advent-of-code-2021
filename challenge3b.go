package main

import (
	"strconv"
	"strings"
)

func Challenge3bFile(s string) (int64, int64) {
	var commands = ReadMeasurements(s)
	return Challenge3b(commands)
}

func Challenge3b(commands []string) (int64, int64) {
	oxygenGeneratorRateBits := ByteRate(commands, "1", "0")
	co2ScrubberRateBits := ByteRate(commands, "0", "1")
	oxygenGeneratorRate, _ := strconv.ParseInt(oxygenGeneratorRateBits, 2, 64)
	co2ScrubberRate, _ := strconv.ParseInt(co2ScrubberRateBits, 2, 64)
	return oxygenGeneratorRate, co2ScrubberRate
}

func ByteRate(commands []string, maxChar string, minChar string) string {
	count := len(commands[0])
	leftOver := commands
	starts := ""
	for i := 0; i < count; i++ {
		count := Count(leftOver, i, "1")
		if (2 * count) >= len(leftOver) {
			starts += maxChar
		} else {
			starts += minChar
		}
		leftOver = Select(leftOver, starts)
		if 1 == len(leftOver) {
			return leftOver[0]
		}
	}
	return leftOver[0]
}

func Select(commands []string, startsWith string) []string {
	return SelectStartsWith(commands, CountStartsWith(commands, startsWith), startsWith)
}

func SelectStartsWith(commands []string, l int, with string) []string {
	result := make([]string, l)
	j := 0
	for i := 0; i < len(commands); i++ {
		t := commands[i]
		if strings.HasPrefix(t, with) {
			result[j] = t
			j++
		}
	}
	return result
}

func CountStartsWith(commands []string, with string) int {
	count := 0
	for i := 0; i < len(commands); i++ {
		t := commands[i]
		if strings.HasPrefix(t, with) {
			count++
		}
	}
	return count
}

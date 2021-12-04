package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadMeasurements(f string) []string {
	file, err := os.Open(f)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	return text
}

func StringToIn(s []string) []int {
	var output = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		output[i], _ = strconv.Atoi(s[i])
	}
	return output
}

func StringToCommand(s []string) []Command {
	var output = make([]Command, len(s))
	for i := 0; i < len(s); i++ {
		s := strings.Split(s[i], " ")
		var command = s[0]
		var amount, _ = strconv.Atoi(s[1])
		output[i] = Command{command, amount}
	}
	return output
}

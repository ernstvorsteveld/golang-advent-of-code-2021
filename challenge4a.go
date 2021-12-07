package main

import (
	"strconv"
	"strings"
)

type Board struct {
	grid   [][]int
	scores [][]bool
}

type Bingo struct {
	draws  []int
	boards []Board
}

func Challenge4aSimple(f string, size int) int {
	lines := ReadMeasurements(f)
	game := Bingo{ReadDraws(lines[0]), ReadBoards(lines, size)}
	return PlayBingo(game)
}

func PlayBingo(game Bingo) int {
	return 0
}

func ReadDraws(s string) []int {
	drawsString := strings.Split(s, ",")
	length := len(drawsString)
	draws := ReadLineAsInteger(length, drawsString)
	return draws
}

func ReadLineAsInteger(length int, drawsString []string) []int {
	draws := make([]int, length)
	for i := 0; i < length; i++ {
		if drawsString[i] == " " {
			continue
		}
		value, _ := strconv.Atoi(drawsString[i])
		draws[i] = value
	}
	return draws
}

func ReadBoards(lines []string, size int) []Board {
	boards := make([]Board, 0)
	remainingBoardLines := lines[2:]
	for len(remainingBoardLines) >= size {
		board := ReadBoard(remainingBoardLines[:size], size)
		boards = append(boards, board)

		if len(remainingBoardLines) > size + 1 {
			remainingBoardLines = remainingBoardLines[size+1:]
		} else {
			remainingBoardLines = remainingBoardLines[size:]
		}
	}
	return boards
}

func ReadBoard(lines []string, size int) Board {
	var gridValues [][]int = make([][]int, size)
	var gridBools [][]bool = make([][]bool, size)
	for k := 0; k < size; k++ {
		gridValues[k] = make([]int, size)
		gridBools[k] = make([]bool, size)
	}

	for i := 0; i < len(lines); i++ {
		lineAsStrings := strings.Split(lines[i], " ")
		line:= ReadLineAsInteger(size, lineAsStrings)
		for j := 0; j < size; j++ {
			gridValues[i][j] = line[j]
			gridBools[i][j] = false
		}
	}
	return Board{grid: gridValues, scores: gridBools}
}

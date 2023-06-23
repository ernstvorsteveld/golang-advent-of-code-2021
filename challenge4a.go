package main

import (
	"strconv"
	"strings"
)

type BingoBoard struct {
	grid [][]BingoField
}

type BingoField struct {
	value int
	draw  bool
}

type Bingo struct {
	size   int
	draws  []int
	boards []BingoBoard
}

type BingoPlayer interface {
	play()
	draw(draw int) (bool, BingoBoard, int, int)
	check() (bool, BingoBoard, int, int)
	value() int
}

func (b Bingo) draw(draw int) (bool, BingoBoard, int, int) {
	for _, board := range b.boards {
		for _, h := range board.grid {
			for _, v := range h {
				v.draw = v.draw || v.value == draw
			}
		}
	}
	return b.check()
}

func (b Bingo) check() (bool, BingoBoard, int, int) {
	for _, board := range b.boards {
		for i, h := range board.grid {
			allHorizontals := false
			for _, v := range h {
				allHorizontals = allHorizontals && v.draw
			}
			if allHorizontals {
				return true, board, i, 0
			}
		}
	}
	for _, board := range b.boards {
		for i := 0; i < b.size; i++ {
			allVerticals := false
			j := 0
			for ; j < b.size; j++ {
				allVerticals = allVerticals && board.grid[j][i].draw
			}
			if allVerticals {
				return true, board, j, i
			}
		}
	}
	return false, BingoBoard{}, 0, 0
}

func (b Bingo) value() int {
	return 0
}

func Challenge4aSimple(f string, size int) int {
	lines := ReadMeasurements(f)
	game := Bingo{size, ReadDraws(lines[0]), ReadBoards(lines, size)}
	return PlayBingo(game)
}

func PlayBingo(play BingoPlayer) int {
	play.play()
	return play.value()
}

func (b Bingo) play() {
	for _, s := range b.draws {
		b.draw(s)
	}
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

func ReadBoards(lines []string, size int) []BingoBoard {
	boards := make([]BingoBoard, 0)
	remainingBoardLines := lines[2:]
	for len(remainingBoardLines) >= size {
		board := ReadBoard(remainingBoardLines[:size], size)
		boards = append(boards, board)

		if len(remainingBoardLines) > size+1 {
			remainingBoardLines = remainingBoardLines[size+1:]
		} else {
			remainingBoardLines = remainingBoardLines[size:]
		}
	}
	return boards
}

func ReadBoard(lines []string, size int) BingoBoard {
	var bingoFields = make([][]BingoField, size)
	for k := 0; k < size; k++ {
		bingoFields[k] = make([]BingoField, size)
	}

	for i := 0; i < len(lines); i++ {
		lineAsStrings := strings.Split(lines[i], " ")
		line := ReadLineAsInteger(size, lineAsStrings)
		for j := 0; j < size; j++ {
			bingoFields[i][j].value = line[j]
			bingoFields[i][j].draw = false
		}
	}
	return BingoBoard{grid: bingoFields}
}

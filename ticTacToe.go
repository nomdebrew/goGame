package main

import (
	"fmt"
	"strconv"
	"time"
)

func PlayTicTacToe() {
	BannerScreen("Tic-Tac-Toe")
	board := buildBoard()
	PrintBoard(board)
	playersMark := "X"
	for isBoardNotFull(board) && !hasPlayerWon(board, playersMark) {
		board, playersMark = locationSelection(board, playersMark)
		PrintBoard(board)
		if hasPlayerWon(board, switchPlayersMark(playersMark)) {
			BannerScreen(switchPlayersMark(playersMark + " Won"))
			time.Sleep(3 * time.Second)
			break
		} else if !isBoardNotFull(board) {
			BannerScreen("Board Full")
			time.Sleep(3 * time.Second)
			break
		}
	}
}

// builds the game board and fills it with empty strings
func buildBoard() [][]string {
	rows := 3
	columns := 3
	board := make([][]string, rows)
	for row := range board {
		board[row] = make([]string, columns)
		for column := range board[row] {
			board[row][column] = " "
		}
	}
	return board
}

// check board for an empty space if found returns false
func isBoardNotFull(board [][]string) bool {
	emptySpaceFound := false
	for row := range board {
		for column := range board[row] {
			if board[row][column] == " " {
				emptySpaceFound = true
			}
		}
	}
	return emptySpaceFound
}

// checks all possible options for a player win
func hasPlayerWon(board [][]string, playersMark string) bool {
	if hasPlayerWonCheckDiagonalBackward(board, playersMark) || hasPlayerWonCheckDiagonalForward(board, playersMark) {
		return true
	}
	for i := 0; i < len(board); i++ {
		if hasPlayerWonCheckRow(board, i, playersMark) ||
			hasPlayerWonCheckColumn(board, i, playersMark) {
			return true
		}
	}
	return false
}

// detects a win horizontially
func hasPlayerWonCheckRow(board [][]string, row int, playersMark string) bool {
	winnerFound := true
	for column := 0; column < len(board); column++ {
		if board[row][column] != playersMark {
			winnerFound = false
		}
	}
	return winnerFound
}

// detects a win vertically
func hasPlayerWonCheckColumn(board [][]string, column int, playersMark string) bool {
	winnerFound := true
	for row := 0; row < len(board); row++ {
		if board[row][column] != playersMark {
			winnerFound = false
		}
	}
	return winnerFound
}

// detects a win with a backward diagonal
func hasPlayerWonCheckDiagonalBackward(board [][]string, playersMark string) bool {
	winnerFound := true
	for i := 0; i < len(board); i++ {
		if board[i][i] != playersMark {
			winnerFound = false
		}
	}
	return winnerFound
}

// detects a win with a forward diagonal
func hasPlayerWonCheckDiagonalForward(board [][]string, playersMark string) bool {
	winnerFound := true
	for i := 0; i < len(board); i++ {
		if board[i][len(board)-(1+i)] != playersMark {
			winnerFound = false
		}
	}
	return winnerFound
}

func locationSelection(board [][]string, playersMark string) ([][]string, string) {
	location := InputReaderToString(playersMark + ", Where would you like to place your mark? ")
	if string(location[0]) == "" || string(location[1]) == "" || string(location[2]) != "" {
		fmt.Println("Invalid Choice")
		return locationSelection(board, playersMark)
	}
	column, row := string(location[0]), string(location[1])

	switch column {
	case "a":
		column = "0"
	case "b":
		column = "1"
	case "c":
		column = "2"
	case "d":
		column = "3"
	case "e":
		column = "4"
	case "f":
		column = "5"
	case "g":
		column = "6"
	default:
		column = "0"
	}

	rowIndex, _ := strconv.Atoi(row)
	columnIndex, _ := strconv.Atoi(column)

	if board[rowIndex][columnIndex] == " " {
		board[rowIndex][columnIndex] = playersMark
	} else {
		fmt.Print("Invalid selection, choose again.")
		locationSelection(board, playersMark)
	}
	return board, switchPlayersMark(playersMark)
}

func switchPlayersMark(playersMark string) string {
	if playersMark == "X" {
		return "O"
	} else {
		return "X"
	}
}

func PrintBoard(board [][]string) {
	CallClear()
	fmt.Println("\n     A        B        C    ")
	fmt.Println(" ===========================")
	rowIndex := 0
	for row := range board {
		fmt.Println(" |       ||       ||       |")
		fmt.Print(rowIndex)
		for column := range board[row] {
			fmt.Print("|   ", board[row][column], "   |")
		}
		fmt.Println("\n |       ||       ||       |")
		fmt.Println(" ===========================")
		rowIndex++
	}
}

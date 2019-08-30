package main

import (
	"fmt"
	"strconv"
)

func TicTacToe() {
	rows := 3
	columns := 3
	board := make([][]string, rows)
	for row := range board {
		board[row] = make([]string, columns)
		for column := range board[row] {
			board[row][column] = " "
		}
	}
	PrintBoard(board)
	playTicTacToe(board)
}

func playTicTacToe(board [][]string) {
	board = locationSelection(board)
	PrintBoard(board)
}

func locationSelection(board [][]string) [][]string {
	location := InputReaderToString("Where would you like to place your mark? ")
	column, row := string(location[0]), string(location[1])
	fmt.Println(column, row)
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
	}
	fmt.Println(column, row)

	rowIndex, _ := strconv.Atoi(row)
	columnIndex, _ := strconv.Atoi(column)

	if board[rowIndex][columnIndex] == " " {
		board[rowIndex][columnIndex] = "X"
	} else {
		fmt.Print("Invalid selection, choose again.")
		locationSelection(board)
	}
	return board
}

func PrintBoard(board [][]string) {
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

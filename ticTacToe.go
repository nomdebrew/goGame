package main

import "fmt"

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
}

func PrintBoard(board [][]string) {
	fmt.Println("\n===========================")
	for row := range board {
		fmt.Println("|       ||       ||       |")
		for column := range board[row] {
			fmt.Print("|   ", board[row][column], "   |")
		}
		fmt.Println("\n|       ||       ||       |")
		fmt.Println("===========================")
	}
}

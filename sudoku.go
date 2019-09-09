package main

import (
	"fmt"
	"math/rand"
)

func PlaySudoku() {
	BannerScreen("Sudoku")
	board := buildSudokuBoard()
	setBoard(board)
	printSudokuBoard(board)
}

// builds the game board and fills it with empty strings
func buildSudokuBoard() [][]int {
	rows := 9
	columns := 9
	board := make([][]int, rows)
	for row := range board {
		board[row] = make([]int, columns)
		// for column := range board[row] {
		// 	board[row][column] = " "
		// }
	}
	return board
}

func setBoard(board [][]int) [][]int {
	row := rand.Intn(10)
	column := rand.Intn(10)
	cellValue := rand.Intn(10)

	board[row][column] = cellValue

	return board
}

func printSudokuBoard(board [][]int) {
	CallClear()
	fmt.Println("\n     A        B        C        D        E        F        G        H        I    ")
	fmt.Println(" =================================================================================")
	rowIndex := 0
	for row := range board {
		fmt.Println(" |       ||       ||       ||       ||       ||       ||       ||       ||       |")
		fmt.Print(rowIndex)
		for column := range board[row] {
			fmt.Print("|   ", board[row][column], "   |")
		}
		fmt.Println("\n |       ||       ||       ||       ||       ||       ||       ||       ||       |")
		fmt.Println(" =================================================================================")
		rowIndex++
	}
}

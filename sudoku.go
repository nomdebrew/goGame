package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PlaySudoku() {
	BannerScreen("Sudoku")
	board := buildSudokuBoard()
	selectGameDifficulty(board)
	printSudokuBoard(board)
}

// select dificulty
func selectGameDifficulty(board [][]string) {
	difficultyLevel := InputReaderToInt("What difficulty would you like to play? (0-10): ")
	for i := 0; i <= difficultyLevel*2; i++ {
		setBoard(board)
	}
}

// builds the game board and fills it with empty strings
func buildSudokuBoard() [][]string {
	rows := 9
	columns := 9
	board := make([][]string, rows)
	for row := range board {
		board[row] = make([]string, columns)
		for column := range board[row] {
			board[row][column] = " "
		}
	}
	return board
}

func checkVerticalRepeat(board [][]string, testValue string, column int) bool {
	checkMap := map[string]bool{
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
	}
	for row := range board {
		checkMap[board[row][column]] = true
	}
	if checkMap[testValue] == true {
		return true
	}
	return false
}

func checkHorizontalRepeat(board [][]string, testValue string, row int) bool {
	checkMap := map[string]bool{
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
	}
	// temporaryTransposedBoard := board.T()
	for aRow := range board {
		if aRow == row {
			for column := range board[row] {
				// fmt.Println("row: ", row, " column: ", column)
				checkMap[board[row][column]] = true
			}
		}
	}
	if checkMap[testValue] == true {
		return true
	}
	return false
}

func checkBlockRepeat(board [][]string, testValue string, row int, column int) bool {
	checkMap := map[string]bool{
		"1": false,
		"2": false,
		"3": false,
		"4": false,
		"5": false,
		"6": false,
		"7": false,
		"8": false,
		"9": false,
	}
	localRow := row % 3
	localColumn := column % 3

	for i := row - localRow; i <= row-localRow+2; i++ {
		for j := column - localColumn; i <= column-localColumn+2; i++ {
			checkMap[board[i][j]] = true
		}
	}
	if checkMap[testValue] == true {
		return true
	}
	return false
}

func setBoard(board [][]string) [][]string {

	rand.Seed(time.Now().UTC().UnixNano())
	row := rand.Intn(9)
	column := rand.Intn(9)
	cellValue := strconv.Itoa(rand.Intn(9) + 1)

	if checkVerticalRepeat(board, cellValue, column) || checkHorizontalRepeat(board, cellValue, row) || checkBlockRepeat(board, cellValue, row, column) {
		return setBoard(board)
	} else {
		board[row][column] = cellValue
		return board
	}
}

func printSudokuBoard(board [][]string) {
	CallClear()
	fmt.Println("\n     A       B       C       D       E       F       G       H       I    ")
	fmt.Println(" =============================================================================")
	rowIndex := 0
	for row := range board {
		fmt.Println(" ||       |       |       ||       |       |       ||       |       |       ||")
		fmt.Print(rowIndex, "||")
		for column := range board[row] {
			fmt.Print("   ", board[row][column], "   ")
			if (column+1)%3 == 0 {
				fmt.Print("||")
			} else {
				fmt.Print("|")
			}
		}
		fmt.Println("\n ||       |       |       ||       |       |       ||       |       |       ||")
		if (row+1)%3 == 0 {
			fmt.Println(" =============================================================================")
		} else {
			fmt.Println(" -----------------------------------------------------------------------------")
		}

		rowIndex++
	}
}

package main

import (
	"fmt"
	"strings"
)

// you shold know what a main function does... get out of here!
func main() {
	CallClear()
	MainMenu()
	// GameOver("Thank you for playing")
}

// main menu
func MainMenu() {
	selectedMenu := "m"
	for selectedMenu != "q" {
		BannerScreen("goGame")
		selectedMenu = menuSelect()
		PlayGame(selectedMenu)
	}
}

// gets user input for game selections in menu
func menuSelect() string {
	selection := strings.ToLower(InputReaderToString("What game would you like to play:  \n(B)  BlackJack\n(P)  Texas Holdem Poker\n(T)  TicTacToe\n(S)  Sudoku\n(Q)  Quit\n"))
	return selection
}

// intializes individual games
func PlayGame(selectedMenu string) {
	switch selectedMenu {
	case "b":
		PlayBlackjack()
	case "p":
		PlayTexasHoldem()
	case "t":
		PlayTicTacToe()
	case "s":
		PlaySudoku()
	case "q":
		GameOver("Thank you for playing")
	default:
		fmt.Println("please select a valid option")
		menuSelect()
	}

}

package main

import (
	"fmt"
	"strings"
)

// you shold know what a main function does... get out of here!
func main() {
	MainMenu()
	// GameOver("Thank you for playing")
}

// main menu
func MainMenu() {
	selectedMenu := "m"
	for selectedMenu != "q" {
		CallClear()
		BannerScreen("goGame")
		selectedMenu = menuSelect()
		PlayGame(selectedMenu)
	}
}

// gets user input for game selections in menu
func menuSelect() string {
	selection := strings.ToLower(InputReaderToString("What game would you like to play:  \n(B)  BlackJack\n(P)  Texas Holdem Poker\n(T)  TicTacToe\n(Q)  Quit\n"))
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
		BannerScreen("Tic Tac Toe")
		TicTacToe()
	case "q":
		GameOver("Thank you for playing")
	default:
		fmt.Println("please select a valid option")
		menuSelect()
	}

}

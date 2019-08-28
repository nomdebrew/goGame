package main

import (
	"fmt"
	"strings"
)

// you shold know what a main function does... get out of here!
func main() {
	CallClear()
	BannerScreen("goGame")
	gameSelect()
	GameOver("Thank you for playing")
}

func gameSelect() {
	selection := strings.ToLower(InputReaderToString("What game would you like to play:  \n(B)  BlackJack\n(P)  Texas Holdem Poker\n(T)  TicTacToe\n(Q)  Quit\n"))
	switch selection {
	case "b":
		numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings("BLACKJACK")
		deck := MakeDeck(numberOfDecksInShoot)
		allPlayers := GeneratePlayers(numberOfPlayers)
		allHands := DrawAllHands(deck, allPlayers)
		SetCardAsUnseen(allPlayers, "Dealer", 2)
		PrintAllHands(allHands)
		PlayGame(allHands, deck)
		DidPlayerWin(allHands)
	case "p":
		numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings("Texas Hold'em")
		deck := MakeDeck(numberOfDecksInShoot)
		allPlayers := GeneratePlayers(numberOfPlayers)
		allHands := DrawAllHands(deck, allPlayers)
		PrintAllHands(allHands)
	case "t":
		BannerScreen("Tic Tac Toe")
		TicTacToe()
	case "q":
		GameOver("Thank you for playing")
	default:
		fmt.Println("please select a valid option")
		gameSelect()
	}

}

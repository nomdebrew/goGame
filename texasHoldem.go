package main

import (
	"fmt"
	"time"
)

func PlayTexasHoldem() {
	BannerScreen("Texas Hold'em")
	numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings()
	deck := MakeDeck(numberOfDecksInShoot)
	allPlayers := GeneratePlayers(numberOfPlayers)
	allHands := DrawAllHands(deck, allPlayers)
	PrintAllHands(allHands)
	fmt.Println("\nStill woking on this part of the project")
	time.Sleep(3 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func PlayTexasHoldem() {
	numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings("Texas Hold'em")
	deck := MakeDeck(numberOfDecksInShoot)
	allPlayers := GeneratePlayers(numberOfPlayers)
	allHands := DrawAllHands(deck, allPlayers)
	PrintAllHands(allHands)
	fmt.Println("\nStill woking on this part of the project")
	time.Sleep(3 * time.Second)
}

package main

// you shold know what a main function does... get out of here!
func main() {
	numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings("BLACKJACK")
	deck := MakeDeck(numberOfDecksInShoot)
	allPlayers := GeneratePlayers(numberOfPlayers)
	allHands := DrawAllHands(deck, allPlayers)
	SetCardAsUnseen(allPlayers, "Dealer", 2)
	PrintAllHands(allHands)
	PlayGame(allHands, deck)
	DidPlayerWin(allHands)
	GameOver("Thank you for playing")
}

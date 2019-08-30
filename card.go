package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// makes deck
func MakeDeck(numberOfDecksInShoot int) map[string]int {
	deck := map[string]int{}
	for i := 0; i < 52; i++ {
		deck[PrettyCard(i%13+1, i%4)] = numberOfDecksInShoot
	}
	return deck
}

// drwas inital cards for all players
func DrawAllHands(deck map[string]int, allPlayers map[string]map[string]int) map[string]map[string]int {
	for player := range allPlayers {
		allPlayers[player] = DrawHand(deck)
	}
	return allPlayers
}

// draws two cards to make intial hand
func DrawHand(deck map[string]int) map[string]int {
	hand := map[string]int{}
	for i := 0; i < 2; i++ {
		hand[DrawCard(deck)] = 1
	}
	return hand
}

// to draw the card form the deck a check is performed to see if the card has already been drawn in not the deck is decremented by 1 if it has the function calls itself recursively until a unique card is found
func DrawCard(deck map[string]int) string {
	card := rand.Intn(52)
	if IsDeckEmpty(deck) == true {
		fmt.Println("You are out of cards")
		time.Sleep(3 * time.Second)
		MainMenu()
	}
	if deck[PrettyCard(card%13+1, card%4)] != 0 {
		deck[PrettyCard(card%13+1, card%4)] = deck[PrettyCard(card%13+1, card%4)] - 1
		return PrettyCard(card%13+1, card%4)
	} else {
		DrawCard(deck)
		// make escape case for when the deck is empty
	}
	return "error"
}

//makes sure the deck isn't empty
func IsDeckEmpty(deck map[string]int) bool {
	totalInDeck := 0
	for _, cardCount := range deck {
		totalInDeck += cardCount
	}
	if totalInDeck > 0 {
		return false
	} else {
		return true
	}
}

// convers two numbers representing card number and suit to a single string ex: 11, 1 => "J♥"
func PrettyCard(card, suit int) string {
	var suitIcon, cardName = "", ""
	switch suit {
	default:
		suitIcon = strconv.Itoa(suit)
	case 0:
		suitIcon = "♠"
	case 1:
		suitIcon = "♥"
	case 2:
		suitIcon = "♦"
	case 3:
		suitIcon = "♣"
	}

	switch card {
	default:
		cardName = strconv.Itoa(card)
	case 11:
		cardName = "J"
	case 12:
		cardName = "Q"
	case 13:
		cardName = "K"
	case 1:
		cardName = "A"
	}

	return cardName + suitIcon
}

// prints hands of all players
func PrintAllHands(allHands map[string]map[string]int) {
	for player, playerHand := range allHands {
		PrintHand(player, playerHand)
	}
}

// prints the players hand omiting dealers second card and if dealer hit blackjack
func PrintHand(player string, hand map[string]int) {
	fmt.Print("\n" + player + ": ")
	for cards := range hand {
		if hand[cards] == 2 {
			fmt.Print("  ")
		} else {
			fmt.Print(cards + " ")
		}
	}
	if player != "Dealer" {
		PrintValueOfHand(hand)
		if ValueOfHand(hand) == 21 {
			fmt.Print("  *BlackJack*  ")
		}
	}
}

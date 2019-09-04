package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// strats Blackjack game
func PlayBlackjack() {
	BannerScreen("Blackjack")
	numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings()
	deck := MakeDeck(numberOfDecksInShoot)
	allPlayers := GeneratePlayers(numberOfPlayers)
	allHands := DrawAllHands(deck, allPlayers)
	SetCardAsUnseen(allPlayers, "Dealer", 2)
	PrintAllHands(allHands)
	InitializeBlackjack(allHands, deck)
	DidPlayerWin(allHands)
	time.Sleep(3 * time.Second)
}

// initializes game logic for each player to play the game
// ** update to play in order **
func InitializeBlackjack(allPlayers map[string]map[string]int, deck map[string]int) {
	for player, _ := range allPlayers {
		allPlayers[player] = PlayerPlayHand(player, allPlayers[player], deck)
	}
}

// takes player's name, hand, deck and returns an updated hand
func PlayerPlayHand(player string, hand map[string]int, deck map[string]int) map[string]int {
	if HasHit21(hand) == true {
		return hand
	}
	PrintHand(player, hand)
	if HasPlayerBusted(hand) == true {
		return hand
	}
	if player == "Dealer" {
		for (HasPlayerBusted(hand) || ValueOfHand(hand) >= 16) == false {
			hand[DrawCard(deck)] = 1
			PlayerPlayHand(player, hand, deck)
		}
		return hand
	}
	response := strings.ToLower(InputReaderToString("Would you like to hit (H) or stay (S)?  "))
	switch response {
	case "s":
		return hand
	case "h":
		hand[DrawCard(deck)] = 1
		PlayerPlayHand(player, hand, deck)
	default:
		fmt.Println("Invalid selection, please choose again. ")
	}
	return hand
}

// checks if hand is equal to 21 and retruns a boolean
func HasHit21(hand map[string]int) bool {
	if ValueOfHand(hand) == 21 {
		return true
	}
	return false
}

// check see if the player has busted
func HasPlayerBusted(hand map[string]int) bool {
	return ValueOver21(ValueOfHand(hand))
}

// check value for being greater than 21
func ValueOver21(totalOfHand int) bool {
	if totalOfHand > 21 {
		//fmt.Print(" BUSTED ")
		return true
	}
	// } else if totalOfHand == 21 {
	// 	fmt.Print(" ***BlackJack*** ")
	// 	return false
	// }
	return false
}

// cauculates the running total of the current hand and returns the sum as an int
func ValueOfHand(hand map[string]int) int {
	totalOfHand := 0
	for card := range hand {
		//fmt.Println(string(card[0]))
		totalOfHand += CardToNumericValue(string(card[0]))
		// if ValueOver21(totalOfHand) && string(card[0]) == "A" {
		// 	card[0] = rune("1")
		// }
	}
	if ValueOver21(totalOfHand) {
		for card := range hand {
			if string(card[0]) == "A" {
				totalOfHand -= 10
			}
		}
	}
	return totalOfHand
}

// takes a string containg the card number(2, 4, J, K, etc) an returns the int value
func CardToNumericValue(card string) int {
	switch card {
	case "J", "Q", "K", "1":
		return 10
	case "A":
		return 11
	default:
		number, _ := strconv.Atoi(card)
		return number
	}
}

// check to see if player won
func DidPlayerWin(allHands map[string]map[string]int) {
	SetAllOfHandAsSeen(allHands["Dealer"])
	dealersHand := ValueOfHand(allHands["Dealer"])
	if ValueOver21(dealersHand) {
		dealersHand = 0
	}
	for player, hand := range allHands {
		if player != "Dealer" && ValueOver21(ValueOfHand(hand)) == false {
			if ValueOfHand(hand) > dealersHand {
				PrintHand(player, hand)
				fmt.Print(" **WINNER** ")
			} else {
				PrintHand(player, hand)
				fmt.Print(" **LOSER** ")
			}
		} else if player == "Dealer" {
			PrintHand(player, hand)
		} else {
			PrintHand(player, hand)
			fmt.Print(" **LOSER** ")
		}
	}

}

// print value of hand
func PrintValueOfHand(hand map[string]int) {
	fmt.Print(" You have: ", ValueOfHand(hand), " ")
}

// gernerate players
func GeneratePlayers(numberOfPlayers int) map[string]map[string]int {
	allHands := map[string]map[string]int{}
	for i := 0; i < numberOfPlayers; i++ {
		allHands["Player"+strconv.Itoa(i)] = map[string]int{}
	}
	allHands["Dealer"] = map[string]int{}
	return allHands
}

// makes on of the dealers cars unseen, usually the second
func SetCardAsUnseen(allHands map[string]map[string]int, playerName string, cardNumber int) map[string]map[string]int {
	i := 0
	for dealer := range allHands[playerName] {
		i++
		if i == cardNumber {
			allHands[playerName][dealer] = 2
		}
	}
	return allHands
}

// makes all cards visible
func SetAllOfHandAsSeen(hand map[string]int) {
	for card := range hand {
		hand[card] = 1
	}
}

// waits for user input, then proceeds to game
func GetCardGameSettings() (int, int) {
	numberOfPlayers := InputReaderToInt("How many players are there? ")
	numberOfDecksInShoot := InputReaderToInt("How many decks would you like the shoot to have? ")
	CallClear()
	return numberOfPlayers, numberOfDecksInShoot
}

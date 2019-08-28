package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	figure "github.com/common-nighthawk/go-figure"
)

// // you shold know what a main function does... get out of here!
// func main() {
// 	numberOfPlayers, numberOfDecksInShoot := GetCardGameSettings("BLACKJACK")
// 	deck := MakeDeck(numberOfDecksInShoot)
// 	allPlayers := GeneratePlayers(numberOfPlayers)
// 	allHands := DrawAllHands(deck, allPlayers)
// 	SetCardAsUnseen(allPlayers, "Dealer", 2)
// 	PrintAllHands(allHands)
// 	PlayGame(allHands, deck)
// 	DidPlayerWin(allHands)
// 	GameOver("Thank you for playing")
// }

// initializes game logic for each player to play the game
// ** update to play in order **
func PlayGame(allPlayers map[string]map[string]int, deck map[string]int) {
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
		// while HasPlayerBusted(hand) == false && totalOfHand <16
		// hand[DrawCard(deck)] = 1
		// PlayerPlayHand(player, hand, deck)
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

// makes deck
func MakeDeck(numberOfDecksInShoot int) map[string]int {
	deck := map[string]int{}
	for i := 0; i < 52; i++ {
		deck[PrettyCard(i%13+1, i%4)] = numberOfDecksInShoot
	}
	return deck
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

// drwas inital cards for all players
func DrawAllHands(deck map[string]int, allPlayers map[string]map[string]int) map[string]map[string]int {
	for player := range allPlayers {
		allPlayers[player] = DrawHand(deck)
	}
	return allPlayers
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
		GameOver("You are out of cards")
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

// clears screen prints welcome message
func WelcomeScreenAndSettings(gameName string) {
	CallClear()
	gameOverBanner := figure.NewFigure("WELCOME", "slant", true)
	gameOverBanner.Print()
	gameOverBanner = figure.NewFigure(gameName, "slant", true)
	gameOverBanner.Print()
}

// waits for user input, then proceeds to game
func GetCardGameSettings(gameName string) (int, int) {
	WelcomeScreenAndSettings(gameName)
	numberOfPlayers := InputReaderToInt("How many players are there? ")
	numberOfDecksInShoot := InputReaderToInt("How many decks would you like the shoot to have? ")
	CallClear()
	return numberOfPlayers, numberOfDecksInShoot
}

// listens for input, prints message to user, converts string to int
func InputReaderToInt(messageToUser string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(messageToUser)
	stringRead, _ := reader.ReadString('\n')
	stringRead = strings.Replace(stringRead, "\r\n", "", -1) // change "\r\n" to "\n" for uninx, end of line termination different
	intRead, anErrorMessage := strconv.Atoi(stringRead)
	if anErrorMessage != nil || intRead == 0 {
		fmt.Println("You entered "+stringRead+", which was read as ", intRead, anErrorMessage)
		fmt.Println("You must enter an integer greater than 0\n")
		InputReaderToInt(messageToUser)
	}
	return intRead
}

// takes string that will be pritned to the user records the resopnse and returns it as a string
func InputReaderToString(messageToUser string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(messageToUser)
	stringRead, _ := reader.ReadString('\n')
	stringRead = strings.Replace(stringRead, "\r\n", "", -1) // change "\r\n" to "\n" for uninx, end of line termination different
	return stringRead
}

// closes game
func GameOver(messageToUser string) {
	fmt.Println("\n", messageToUser)
	gameOverBanner := figure.NewFigure("Game   Over", "slant", true)
	gameOverBanner.Print()
	time.Sleep(3 * time.Second)
	os.Exit(3)
}

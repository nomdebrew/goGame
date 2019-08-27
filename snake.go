package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// you shold know what a main function does... get out of here!
func main() {
	numberOfPlayers, numberOfDecksInShoot := welcomeScreenAndSettings()
	deck := makeDeck(numberOfDecksInShoot)
	allPlayers := generatePlayers(numberOfPlayers)
	allHands := drawAllHands(deck, allPlayers)
	setCardAsUnseen(allPlayers, "Dealer", 2)
	printAllHands(allHands)
	playGame(allHands, deck)
	didPlayerWin(allHands)
	//printAllHands(allHands)
}

// initializes game logic for each player to play the game
// ** update to play in order **
func playGame(allPlayers map[string]map[string]int, deck map[string]int) {
	for player, _ := range allPlayers {
		allPlayers[player] = playerPlayHand(player, allPlayers[player], deck)
	}
}

// takes player's name, hand, deck and returns an updated hand
func playerPlayHand(player string, hand map[string]int, deck map[string]int) map[string]int {
	if hasHit21(hand) == true {
		return hand
	}
	printHand(player, hand)
	if hasPlayerBusted(hand) == true {
		return hand
	}
	if player == "Dealer" {
		for (hasPlayerBusted(hand) || valueOfHand(hand) >= 16) == false {
			hand[drawCard(deck)] = 1
			playerPlayHand(player, hand, deck)
		}
		// while hasPlayerBusted(hand) == false && totalOfHand <16
		// hand[drawCard(deck)] = 1
		// playerPlayHand(player, hand, deck)
		return hand
	}
	response := strings.ToLower(inputReaderToString("Would you like to hit (H) or stay (S)?  "))
	switch response {
	case "s":
		return hand
	case "h":
		hand[drawCard(deck)] = 1
		playerPlayHand(player, hand, deck)
	default:
		fmt.Println("Invalid selection, please choose again. ")
	}
	return hand
}

// checks if hand is equal to 21 and retruns a boolean
func hasHit21(hand map[string]int) bool {
	if valueOfHand(hand) == 21 {

		return true
	}
	return false
}

// check see if the player has busted
func hasPlayerBusted(hand map[string]int) bool {
	return valueOver21(valueOfHand(hand))
}

// check value for being greater than 21
func valueOver21(totalOfHand int) bool {
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
func valueOfHand(hand map[string]int) int {
	totalOfHand := 0
	for card := range hand {
		//fmt.Println(string(card[0]))
		totalOfHand += cardToNumericValue(string(card[0]))
		// if valueOver21(totalOfHand) && string(card[0]) == "A" {
		// 	card[0] = rune("1")
		// }
	}
	if valueOver21(totalOfHand) {
		for card := range hand {
			if string(card[0]) == "A" {
				totalOfHand -= 10
			}
		}
	}
	return totalOfHand
}

// takes a string containg the card number(2, 4, J, K, etc) an returns the int value
func cardToNumericValue(card string) int {
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
func didPlayerWin(allHands map[string]map[string]int) {
	setAllOfHandAsSeen(allHands["Dealer"])
	dealersHand := valueOfHand(allHands["Dealer"])
	if valueOver21(dealersHand) {
		dealersHand = 0
	}
	for player, hand := range allHands {
		if player != "Dealer" && valueOver21(valueOfHand(hand)) == false {
			if valueOfHand(hand) > dealersHand {
				printHand(player, hand)
				fmt.Print(" **WINNER** ")
			} else {
				printHand(player, hand)
				fmt.Print(" **LOSER** ")
			}
		} else if player == "Dealer" {
			printHand(player, hand)
		} else {
			printHand(player, hand)
			fmt.Print(" **LOSER** ")
		}
	}

}

// makes deck
func makeDeck(numberOfDecksInShoot int) map[string]int {
	deck := map[string]int{}
	for i := 0; i < 52; i++ {
		deck[prettyCard(i%13+1, i%4)] = numberOfDecksInShoot
	}
	return deck
}

// prints hands of all players
func printAllHands(allHands map[string]map[string]int) {
	for player, playerHand := range allHands {
		printHand(player, playerHand)
	}
}

// prints the players hand omiting dealers second card and if dealer hit blackjack
func printHand(player string, hand map[string]int) {
	fmt.Print("\n" + player + ": ")
	for cards := range hand {
		if hand[cards] == 2 {
			fmt.Print("  ")
		} else {
			fmt.Print(cards + " ")
		}
	}
	if player != "Dealer" {
		printValueOfHand(hand)
		if valueOfHand(hand) == 21 {
			fmt.Print("  *BlackJack*  ")
		}
	}
}

// print value of hand
func printValueOfHand(hand map[string]int) {
	fmt.Print(" You have: ", valueOfHand(hand), " ")
}

// gernerate players
func generatePlayers(numberOfPlayers int) map[string]map[string]int {
	allHands := map[string]map[string]int{}
	for i := 0; i < numberOfPlayers; i++ {
		allHands["Player"+strconv.Itoa(i)] = map[string]int{}
	}
	allHands["Dealer"] = map[string]int{}
	return allHands
}

// drwas inital cards for all players
func drawAllHands(deck map[string]int, allPlayers map[string]map[string]int) map[string]map[string]int {
	for player := range allPlayers {
		allPlayers[player] = drawHand(deck)
	}
	return allPlayers
}

// makes on of the dealers cars unseen, usually the second
func setCardAsUnseen(allHands map[string]map[string]int, playerName string, cardNumber int) map[string]map[string]int {
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
func setAllOfHandAsSeen(hand map[string]int) {
	for card := range hand {
		hand[card] = 1
	}
}

// draws two cards to make intial hand
func drawHand(deck map[string]int) map[string]int {
	hand := map[string]int{}
	for i := 0; i < 2; i++ {
		hand[drawCard(deck)] = 1
	}
	return hand
}

// to draw the card form the deck a check is performed to see if the card has already been
// drawn in not the deck is decremented by 1 if it has the function calls itself recursively
// until a unique card is found
func drawCard(deck map[string]int) string {
	card := rand.Intn(52)
	if deck[prettyCard(card%13+1, card%4)] != 0 {
		deck[prettyCard(card%13+1, card%4)] = deck[prettyCard(card%13+1, card%4)] - 1
		return prettyCard(card%13+1, card%4)
	} else {
		drawCard(deck)
		// make escape case for when the deck is empty
	}
	return "error"
}

// convers two numbers representing card number and suit to a single string ex: 11, 1 => "J♥"
func prettyCard(card, suit int) string {
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

// clears screen prints welcome message, waits for user input, then proceeds to game
func welcomeScreenAndSettings() (int, int) {
	CallClear()
	fmt.Println("                      Welcome, WELCOME, WELCOME!!!!")
	fmt.Println("....................._____......................____..........................")
	fmt.Println("..................../ __  \\..................../   |..........................")
	fmt.Println(".................../_/..\\  \\................./_/.| |..........................")
	fmt.Println(".........................| |.....................| |..........................")
	fmt.Println("......................../ /......................| |..........................")
	fmt.Println("......................./ /.......................| |..........................")
	fmt.Println("...................../ /.........................| |..........................")
	fmt.Println("..................../ /______.................___| |___.......................")
	fmt.Println("...................|_________|................|________|......................\n")

	numberOfPlayers := inputReaderToInt("How many players are there? ")
	numberOfDecksInShoot := inputReaderToInt("How many decks would you like the shoot to have? ")
	//time.Sleep(3 * time.Second)
	CallClear()
	// return numberOfPlayers, numberOfDecksInShoot
	return numberOfPlayers, numberOfDecksInShoot
}

// listens for input, prints message to user, converts string to int
func inputReaderToInt(messageToUser string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(messageToUser)
	stringRead, _ := reader.ReadString('\n')
	stringRead = strings.Replace(stringRead, "\r\n", "", -1) // change "\r\n" to "\n" for uninx, end of line termination different
	intRead, anErrorMessage := strconv.Atoi(stringRead)
	if anErrorMessage != nil || intRead == 0 {
		fmt.Println("You entered "+stringRead+", which was read as ", intRead, anErrorMessage)
		fmt.Println("You must enter an integer greater than 0\n")
		inputReaderToInt(messageToUser)
	}
	return intRead
}

// takes string that will be pritned to the user records the resopnse and returns it as a string
func inputReaderToString(messageToUser string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(messageToUser)
	stringRead, _ := reader.ReadString('\n')
	stringRead = strings.Replace(stringRead, "\r\n", "", -1) // change "\r\n" to "\n" for uninx, end of line termination different
	return stringRead
}

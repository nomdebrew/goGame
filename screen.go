package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	figure "github.com/common-nighthawk/go-figure"
)

// claer screen code from https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/*CallClear clears the screen*/
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// clears screen prints message as ASCII banner
func BannerScreen(message string) {
	welcomeBanner := figure.NewFigure(message, "slant", true)
	welcomeBanner.Print()
}

// closes game
func GameOver(messageToUser string) {
	fmt.Println("\n", messageToUser)
	gameOverBanner := figure.NewFigure("Game   Over", "slant", true)
	gameOverBanner.Print()
	time.Sleep(3 * time.Second)
	os.Exit(3)
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

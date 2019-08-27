package main

import (
	"os"
	"os/exec"
	"runtime"
)

// code from https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go

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

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// func main() {
// 	fmt.Println("I will clean the screen in 2 seconds!")
// 	time.Sleep(2 * time.Second)
// 	CallClear()
// 	fmt.Println("I'm alone...")
// }

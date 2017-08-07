package main

import (
	"fmt"
	"net/http"
	"os"

	// importing the required libraries
	"github.com/almakinah/gobot/bot"
	"github.com/almakinah/gobot/calculator"
)

// Our program entry point
func main() {
	// We are building a non-blocking chat calculator bot so we need to build
	// channels to be able to pass data between our sub-routines (threads)
	digits := make(chan string)               // Channel to pass numbers
	calculator.OutputChan = make(chan string) // Channel to pass output

	// Read the required tokens from the environment variables and
	// set them on the bot package; this is safer than hard-coding them
	bot.VerificationToken = os.Getenv("GOBOT_VERIFICATION_TOKEN")
	bot.PageAccessToken = os.Getenv("GOBOT_PAGE_ACCESS_TOKEN")

	// Initialize the bot server passing the channels
	// since `initServer` is blocking we launch it in its own
	// go sub-routine
	go initServer(digits, calculator.OutputChan)

	// Print a nice welcome message!
	fmt.Println("Hi welcome to chat bot")

	// `WaitForDigits` is blocking however since this is the only other
	// blocking routine in our app we don't need to run it on a separate
	// sub-routine, since doing so will terminate our main function and
	// terminate the bot
	calculator.WaitForDigits(digits)
}

// Initialize the server using http's `ListenAndServe`
// in the function definition we define our channel parameters
// as send only or recieve only; with digits we will only be sending
// and with output we are only going to revieve
func initServer(digits chan<- string, output <-chan string) {
	bot.DigitsChan = digits
	bot.OutputChan = output

	// Use port 8080
	// handling errors in Go is done using return values, if the returned
	// err variable is not empty then an error has occured and we need to
	// handle it properly.
	err := http.ListenAndServe(":8080", initRouter())
	if err != nil {
		fmt.Println(err)
		panic("Couldn't serve gogo bot")
	}
}

// This package defines the required functions and methods to communicate
// with the Facebook bot API
package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Export some channels so we are able to send and recieve data from other
// Sub-routines; idealy this should be done using `http#Context`
var DigitsChan chan<- string
var OutputChan <-chan string

// Export a couple of variables to hold the verification token and page
// access token
var VerificationToken, PageAccessToken string

// Where we handle the verification process by Facebook
func VerifyToken(writer http.ResponseWriter, request *http.Request) {
	// Facebook sends a request to our server including the `verify_token` we set
	// its expected that our server replies back with the challenge sent within
	// the request
	if request.FormValue("hub.mode") == "subscribe" && request.FormValue("hub.verify_token") == VerificationToken {
		fmt.Fprintf(writer, request.FormValue("hub.challenge"))
	}
}

// Where we handle recieving a message from a user on our page
func RecieveMessage(writer http.ResponseWriter, request *http.Request) {
	// first we decode the recieved json message
	var bMessage botEvent
	err := json.NewDecoder(request.Body).Decode(&bMessage)
	if err != nil {
		log.Println(err)
	}

	// since we are building a calculator, we expect the user to send either
	// a digit or an operator, in all cases we push the message to our DigitsChan
	DigitsChan <- bMessage.String()

	// on a separate go routine we wait for an output using an anonymous function
	// we do this on a separate routine because `<-OutputChan` will block until a
	// value is recieved on the channel
	go func() {
		output := <-OutputChan

		// if the output is an empty string then we need to terminate the function
		// this happens when the calculator hasn't yet recieved an operator
		if output == "" {
			return
		} else {
			// otherwise we send the output back to the user
			SendMessage(bMessage.Entries[0].Messaging[0].Sender.Id, output)
		}
	}()
}

// Sends a message to a specific user id on facebook
func SendMessage(recipientId string, message string) {
	// format the request url
	url := fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%v", PageAccessToken)
	// we create buffer to hold our request which will be generated from JSON
	// templates
	buf := bytes.NewBufferString("")

	// generate the response template
	genResponseTemplate(recipientId, message, buf)

	// log the request being sent
	log.Println(buf.String())
	// send our request
	resp, err := http.Post(url, "application/json", buf)

	if err != nil {
		log.Println(err)
	}

	// log the response status to make sure everything went well
	log.Println(resp.Status)
}

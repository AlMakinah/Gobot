// this package handles the calculator part of our bot
package calculator

import (
	"fmt"
	"strconv"
)

// heap is where we store the digits that we will perform the calculation
// upon
var heap []int

// Where we are going to send the output
var OutputChan chan<- string

// Infinitly wait for digits sent by the user
func WaitForDigits(digits <-chan string) {
	for {
		st := <-digits             // wait for digits to be sent on the digits channel
		d, err := strconv.Atoi(st) // try to convert the string into an int
		if err != nil {            // if an error occurs then the string is not a number
			if st == "+" || st == "-" || st == "*" || st == "/" { // check for operators
				calculateHeap(st) // if its an operator, perform the calculation
			} else {
				OutputChan <- "" // its not an operator, just ignore it and don't wait
			}
		} else {
			heap = append(heap, d) // its a digit, add it to the heap
			OutputChan <- ""       // don't wait
		}
	}
}

// Calculate the operation
func calculateHeap(operator string) {
	// Setup the output default value
	var output int
	if operator == "/" || operator == "*" {
		output = 1
	} else {
		output = 0
	}

	// Loop over the heap and perform the operation
	for _, digit := range heap {
		switch operator {
		case "+":
			output += digit
		case "-":
			output -= digit
		case "*":
			output *= digit
		case "/":
			output /= digit
		}
	}

	// send the output to the Output channel
	OutputChan <- fmt.Sprintf("%d", output)
	// clear the heap!
	heap = heap[:0]
}

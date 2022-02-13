package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	fmt.Println("Hello World")
	fmt.Println(quote.Go())

	// Get a greeting message and print it
	// message := greetings.Hello("Taro")
	// fmt.Println(message)

	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an error was retruned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message to the console.
	fmt.Println(message)
}

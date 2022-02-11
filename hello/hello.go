package main

import (
	"fmt"

	"../greetings"
)

func main() {
	fmt.Println("Hello World")
	// fmt.Println(quote.Go())

	// Get a greeting message and print it
	message := greetings.Hello("Taro")
	fmt.Println(message)

}

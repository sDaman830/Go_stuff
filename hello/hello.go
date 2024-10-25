package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("HDYD")
	fmt.Println(message)
}

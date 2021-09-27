package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set up logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// request a greeting messages for the name
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// if no error was returned, print the returned map of messages
	fmt.Println(messages)
}

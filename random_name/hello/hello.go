package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main () {
	//set properties of predefined Logger, including the 
	// log entry prefix and a flag to disable printing the
	// time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request greeting message
	message, err := greetings.Hello("Gladys")

	// if error, print and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error, print message
	fmt.Println(message)
}

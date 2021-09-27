package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	//set properties of the predifined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, the source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request greeting message
	message, err := greetings.Hello("")

	// if an error is returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error is returned, print the returned message    
	// to the console
	fmt.Println(message)
}

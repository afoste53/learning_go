package greetings

import "fmt"


/**
func Hello(name string) string
 	^ func  ^name   ^var    ^return type
        name   of var  type
 function that starts with capital letter can be called by a function
 not in the same package - an exported name **/

// Hello returns a greeting for the named person
func Hello(name string) string {
   // return a greeting that embeds the name in a message
   message := fmt.Sprintf("Hi, %v. Welcome!", name)
   /** 
	:= decalres and initializes variable in same line
	Sprintf replaces %v with specified variable
    **/
    return message
}

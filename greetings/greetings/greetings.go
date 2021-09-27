package greetings

import(
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is provided, return an error
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return valie that embeds name
	// in a greeting message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

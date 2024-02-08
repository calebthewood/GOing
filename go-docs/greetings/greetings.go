package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Common Commands
// % go run
// % go build
// % go install
// % go test

// Hellow returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name, return err
	if name == "" {
		return "", errors.New("empty name")
	}
	// if name, return greeting
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messaages.
	messages := make(map[string]string)
	// loop thru the rec'd slice of names, calling Hello to get a msg for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate msg with the name
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// running  ' % go mod tidy ' in cl will find and install imports
func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// request greeting messafes for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		// log.Fatal will print err and then exit program.
		log.Fatal(err)
	}
	fmt.Println(messages)
}

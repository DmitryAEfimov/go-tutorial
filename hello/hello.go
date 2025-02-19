package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including the log entry prefix and a flag to disable printing
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

    // Request a greeting message
	// If an error was returned, print it to the console and exit the program.
	message, err := greetings.Hello("Genifer")
	checkError(err)
	
	// If no error was returned, print the returned message to the console.
	fmt.Println(message)

	// A slice of names.
    names := []string {"Gladys", "Samantha", "Darrin"}
	// Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
	checkError(err)

	// If no error was returned, print the returned map of messages to the console.
    fmt.Println(messages)
}

func checkError(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}
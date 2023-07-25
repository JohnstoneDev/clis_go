package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"example.com/greetings" // import the greetings package that will print a greeting
)

func main(){
		// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// To illustrate randomness, some names & empty string
	names := []string{
		"Din Djarin",
		"General Kenobi",
		"Anakin Skywalker",
		"",
	}
	// seed value so that the random number changes on every run
	seed := time.Now().UnixNano()
	name := names[rand.Intn(len(names))]

	rand.Seed(seed)
	message, err := greetings.Hello(name)

	// log an error to the console if there's one
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
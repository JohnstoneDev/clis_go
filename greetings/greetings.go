package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello function that takes in a name & prints a greeting || error
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	return fmt.Sprintf(randomFormat(), name), nil
}

func randomFormat() string {
	// formats slice with different forms of greeting
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hello there, %v!",
		"Hail, %v! Well met!",
	}

	// seed valueu so thagt the random number changes on every run
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	return formats[rand.Intn(len(formats))]
}
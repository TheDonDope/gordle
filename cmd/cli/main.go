package main

import (
	"fmt"
	"os"

	"github.com/erikgeiser/promptkit/textinput"
)

const (
	correct = "🟩"
	wrong = "🟨"
	nospot = "⬛"
)

type Guess struct {
	Next *Guess
	Value interface{}
}

func main() {
	//var first *Guess
	//guesses := [6]*Guess{}
	//\033[0;0H

	input := textinput.New("Enter 5 characters")
	input.InitialValue = "⬛⬛⬛⬛⬛"
	input.Placeholder = "Your need exactly 5 characters"

	guess, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	// do something with the result
	_ = guess
}

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
	gotd = "ALTER"
	cr = "\033[0;0H"
)

// Guess ...
type Guess struct {
	Value string
	Evaluation string
}

func main() {
	guesses := [6]*Guess{}

	input := textinput.New("Enter 5 characters")
	input.InitialValue = ""
	input.Placeholder = "You need exactly 5 characters"
	input.Validate = func(value string) bool {
		 return len(value) <=5
	}

	for try:=0; try<6; try++ {
		val, err := input.RunPrompt()
		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}
		g := &Guess {
			Value: val,
			Evaluation: "⬛⬛⬛⬛⬛",
		}
		fmt.Printf("%v %v\n", g.Value, g.Evaluation)
		guesses[try] = g
	}

}

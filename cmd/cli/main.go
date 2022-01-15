package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/erikgeiser/promptkit/textinput"
)

const (
	correct = "🟩"
	wrong   = "🟨"
	nospot  = "⬛"
	gotd    = "ALTER"
	max     = 6
	cr      = "\033[0;0H"
)

// Guess ...
type Guess struct {
	Value      string
	Evaluation string
}

func evaluate(guess string, solution string) string {
	evaluation := []string{"⬛", "⬛", "⬛", "⬛", "⬛"}
	if guess == solution {
		for i := range evaluation {
			evaluation[i] = "🟩"
		}
	}
	return strings.Join(evaluation, "")
}

func won(guess *Guess) bool {
	return guess.Evaluation == "🟩🟩🟩🟩🟩"
}

func newGuess(guess string) *Guess {
	return &Guess{
		Value:      guess,
		Evaluation: evaluate(guess, gotd),
	}
}

func main() {
	guesses := []*Guess{}

	r := 1

	fmt.Println("Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩")
	fmt.Println("You have 6 trys to guess the word of the day.")
	fmt.Println("🟩 means, the letter is in the word and in the correct spot.")
	fmt.Println("🟨 means, that the letter is in the word but in the wrong spot.")
	fmt.Println("⬛ means, that the letter is in not in the word in any spot.")

	input := textinput.New("Enter 5 characters")
	input.InitialValue = ""
	input.Placeholder = "You need exactly 5 characters"
	input.Validate = func(value string) bool {
		return len(value) == 5
	}

	for try := 0; try < 6; try++ {
		val, err := input.RunPrompt()
		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}
		g := newGuess(val)
		fmt.Printf("%v %v (Try %v/%v)\n", g.Value, g.Evaluation, r, max)
		guesses = append(guesses, g)
		if won(g) {
			fmt.Println("\nCongratulations, you won! 🥳🥳")
			break
		}
		r++
	}
	fmt.Printf("\nYour Gordle results:\n")
	for i, v := range guesses {
		fmt.Printf("%v %v (%v/%v)\n", v.Value, v.Evaluation, i+1, max)
	}
	fmt.Printf("\nThe solution was: %s", gotd)
}

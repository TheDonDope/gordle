package main

import (
	"fmt"
	"os"
	"time"

	"github.com/erikgeiser/promptkit/textinput"

	val "github.com/TheDonDope/gordle/pkg/validation"
)

const (
	wotd = "ALTER"
	max  = 6
	cr   = "\033[0;0H"
)

func main() {
	guesses := []*val.Guess{}

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
		prmpt, err := input.RunPrompt()
		if err != nil {
			fmt.Printf("Error: %v\n", err)

			os.Exit(1)
		}
		g := val.NewGuess(prmpt, wotd)
		fmt.Printf("%v %v (Try %v/%v)\n", g.Value, g.Evaluation, r, max)
		guesses = append(guesses, g)
		if g.Won() {
			fmt.Println("\nCongratulations, you won! 🥳🥳")
			break
		}
		r++
	}
	fmt.Printf("\nYour Gordle results (%v):\n", time.Now().Format("2006-01-02"))
	for i, v := range guesses {
		fmt.Printf("%v %v (%v/%v)\n", v.Value, v.Evaluation, i+1, max)
	}
	fmt.Printf("\nThe solution was: %s", wotd)
}

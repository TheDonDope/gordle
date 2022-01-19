package game

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/erikgeiser/promptkit/textinput"

	storage "github.com/TheDonDope/gordle/pkg/storage"
)

const (
	maxTrys   = 6
	rightSpot = "🟩"
	wrongSpot = "🟨"
	noSpot    = "⬛"
	winGuess  = "🟩🟩🟩🟩🟩"
	failGuess = "⬛⬛⬛⬛⬛"
)

// Game encapsulates the state of the guessing game.
type Game struct {
	round   int
	guesses []*Guess
	dict    *storage.Dictionary
	wotd    string
}

// NewGame returns a fresh round of gordle!
func NewGame(dict io.Reader) *Game {
	g := &Game{
		round:   1,
		guesses: []*Guess{},
		dict:    storage.NewDictionary(maxTrys, dict),
	}
	g.wotd = g.dict.NewWotd()
	return g
}

// GreetPlayers says hello to the players.
func GreetPlayers() {
	fmt.Println("Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩")
	fmt.Println("You have 6 trys to guess the word of the day.")
	fmt.Println("NOTE: The current implementation will pick a new word on every run!")
	fmt.Println("🟩 means, the letter is in the word and in the correct spot.")
	fmt.Println("🟨 means, that the letter is in the word but in the wrong spot.")
	fmt.Println("⬛ means, that the letter is in not in the word in any spot.")
}

// Play runs the main game loop.
func (g *Game) Play() {
	input := textinput.New("Enter 5 characters")
	input.InitialValue = ""
	input.Placeholder = "You need exactly 5 characters"
	input.Validate = func(value string) bool {
		return len(value) == 5
	}

	for try := 0; try < maxTrys; try++ {
		prmpt, err := input.RunPrompt()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		input := NewGuess(prmpt, g)
		fmt.Printf("%v (Try %v/%v): %v\n", input.Rating, g.round, maxTrys, input.Prompt)
		g.guesses = append(g.guesses, input)
		if g.GuessWon(input) {
			fmt.Println("\nCongratulations, you won! 🥳")
			break
		}
		g.round++
	}
}

// PrintResults shows the success or failure result to the player.
func (g *Game) PrintResults() {
	fmt.Printf("\nYour Gordle results (%v):\n", time.Now().Format("2006-01-02"))
	for i, v := range g.guesses {
		fmt.Printf("%v (%v/%v): %v\n", v.Rating, i+1, maxTrys, v.Prompt)
	}
	fmt.Printf("\nThe solution was: %s\n", g.wotd)
}

// RateGuess returns the calculated emoji square result of the given guess.
func (g *Game) RateGuess(p *Guess) string {
	return RateGuess(p.Prompt, g.wotd)
}

// GuessWon returns true if the given guess matches the word of the day.
func (g *Game) GuessWon(p *Guess) bool {
	return g.wotd == p.Prompt
}

// Guess encapsulates the user input and its emoji-square-rating.
type Guess struct {
	Prompt string
	Rating string
}

// NewGuess creates a new guess and matches it to the solution.
func NewGuess(prompt string, game *Game) *Guess {
	g := &Guess{
		Prompt: prompt,
	}
	g.Rating = RateGuess(prompt, game.wotd)
	return g
}

// RateGuess returns the calculated emoji square result of the given guess and solution.
func RateGuess(guess string, solution string) string {
	rating := []string{noSpot, noSpot, noSpot, noSpot, noSpot}
	if guess == solution {
		for i := range rating {
			rating[i] = rightSpot
		}
		return strings.Join(rating, "")
	}
	for i := range guess {
		// Is the character present in the rest of the solution?
		for j := range solution {
			if j != i && (guess[i] == solution[j]) {
				rating[i] = wrongSpot
			}
		}
		// Is the character exactly the same?
		if guess[i] == solution[i] {
			rating[i] = rightSpot
		}
	}
	return strings.Join(rating, "")
}

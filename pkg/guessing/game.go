package guessing

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/erikgeiser/promptkit/textinput"
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
	round   uint8
	guesses []*Guess
	wotd    string
}

// NewGame returns a fresh round of gordle!
func NewGame() *Game {
	return &Game{
		round:   1,
		guesses: []*Guess{},
		wotd:    newWotd(),
	}
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
		fmt.Printf("%v %v (Try %v/%v)\n", input.Prompt, input.Rating, g.round, maxTrys)
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
		fmt.Printf("%v %v (%v/%v)\n", v.Prompt, v.Rating, i+1, maxTrys)
	}
	fmt.Printf("\nThe solution was: %s\n", g.wotd)
}

// RateGuess returns the calculated emoji square result of the given guess.
func (g *Game) RateGuess(p *Guess) string {
	return RateGuess(p.Prompt, g.wotd)
}

// GuessWon returns true if the given guess matches the word of the day.
func (g *Game) GuessWon(p *Guess) bool {
	return guessWon(p.Prompt, g.wotd)
}

func newWotd() string {
	wotd := ""

	allWords := readWords()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	for i := 0; len(wotd) != 5; i++ {
		w := allWords[i]
		if len(w) == 5 && onlyAlpha(w) {
			wotd = w
			return strings.ToLower(wotd)
		}
	}
	return wotd
}

func guessWon(guess string, solution string) bool {
	return guess == solution
}

func onlyAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readWords() (words []string) {
	words = []string{"NODIC"}
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	words = strings.Split(string(bytes), "\n")
	return
}

package validation

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

const (
	rightSpot = "🟩"
	wrongSpot = "🟨"
	noSpot    = "⬛"
)

// Guesser ...
type Guesser interface {
	// Won returns true if the whole word matched 100% to the word of the day.
	Won(solution string) bool

	// Match returns the pattern of matching to the given solution word.
	Match(solution string) string
}

// Guess ...
type Guess struct {
	Value      string
	Evaluation string
}

// Match returns the pattern of matching to the given solution word.
func (g *Guess) Match(solution string) string {
	evaluation := []string{noSpot, noSpot, noSpot, noSpot, noSpot}
	if g.Value == solution {
		for i := range evaluation {
			evaluation[i] = rightSpot
		}
		return strings.Join(evaluation, "")
	}
	for i := range g.Value {
		// Is the character present in the rest of the solution?
		for j := range solution {
			if j != i && (g.Value[i] == solution[j]) {
				evaluation[i] = wrongSpot
			}
		}
		// Is the character exactly the same?
		if g.Value[i] == solution[i] {
			evaluation[i] = rightSpot
		}
	}
	return strings.Join(evaluation, "")
}

// Won returns true if the whole word matched 100% to the word of the day.
func (g *Guess) Won() bool {
	return g.Evaluation == "🟩🟩🟩🟩🟩"
}

// NewGuess creates a new guess and matches it to the solution.
func NewGuess(guess string, solution string) *Guess {
	g := &Guess{
		Value: guess,
	}
	g.Evaluation = g.Match(solution)
	return g
}

// NewWotd will return a new word of the day.
func NewWotd() string {
	word := ""

	allWords := readWords()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	for i := 0; len(word) != 5; i++ {
		w := allWords[i]
		if len(w) == 5 && onlyAlpha(w) {
			word = w
			return strings.ToLower(word)
		}
	}
	return word
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
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	words = strings.Split(string(bytes), "\n")
	return
}

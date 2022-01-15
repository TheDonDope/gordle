package validation

import "strings"

const (
	correct = "🟩"
	wrong   = "🟨"
	nospot  = "⬛"
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
	evaluation := []string{"⬛", "⬛", "⬛", "⬛", "⬛"}
	if g.Value == solution {
		for i := range evaluation {
			evaluation[i] = "🟩"
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

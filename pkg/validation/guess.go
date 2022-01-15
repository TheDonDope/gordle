package validation

import "strings"

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
	}
	for i := range g.Value {
		// Is it exactly the same?
		if g.Value[i] == solution[i] {
			evaluation[i] = rightSpot
		}
		// If not, is the character present in the rest of the solution?
		for j := range solution {
			if j != i && (g.Value[i] == solution[j]) {
				evaluation[i] = wrongSpot
			}
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

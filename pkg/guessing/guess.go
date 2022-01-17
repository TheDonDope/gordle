package guessing

import "strings"

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

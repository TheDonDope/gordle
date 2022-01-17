package guessing

import (
	"fmt"
	"testing"
)

func TestRateGuessAllCorrectGuess(t *testing.T) {
	g := NewGame()
	g.wotd = "jesus"
	prompt := "jesus"

	guess := NewGuess(prompt, g)

	want := winGuess
	got := g.RateGuess(guess)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestRateGuessAllWrongGuess(t *testing.T) {
	g := NewGame()
	g.wotd = "yolol"
	prompt := "duned"

	guess := NewGuess(prompt, g)

	want := failGuess
	got := g.RateGuess(guess)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestRateGuessReturnsCorrectResult(t *testing.T) {
	g := NewGame()
	g.wotd = "affen"
	prompt := "after"
	guess := NewGuess(prompt, g)

	want := "🟩🟩⬛🟩⬛"
	got := g.RateGuess(guess)

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestGuessWonSucceeds(t *testing.T) {
	g := NewGame()
	g.wotd = "affen"
	prompt := "affen"
	guess := NewGuess(prompt, g)

	want := true
	got := g.GuessWon(guess)

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestGuessWonFails(t *testing.T) {
	g := NewGame()
	g.wotd = "affen"
	prompt := "after"
	guess := NewGuess(prompt, g)

	want := false
	got := g.GuessWon(guess)

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

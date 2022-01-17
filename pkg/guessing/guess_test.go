package guessing

import (
	"fmt"
	"testing"
)

func TestMatchReturns5GreenSquaresOnCompleteCorrectGuess(t *testing.T) {
	solution := "JESUS"
	prmpt := "JESUS"

	guess := NewGuess(prmpt, solution)

	want := "🟩🟩🟩🟩🟩"
	got := guess.Match(solution)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestMatchReturns5BlackSquaresIfWordDoesNotMatchAtAll(t *testing.T) {
	solution := "YOLOL"
	prmpt := "DUNED"

	guess := NewGuess(prmpt, solution)

	want := "⬛⬛⬛⬛⬛"
	got := guess.Match(solution)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestMatchReturnsCorrectResult(t *testing.T) {
	solution := "AFFEN"
	prmpt := "AFTER"

	guess := NewGuess(prmpt, solution)

	want := "🟩🟩⬛🟩⬛"
	got := guess.Match(solution)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestWonSucceeds(t *testing.T) {
	solution := "AFFEN"
	prmpt := "AFFEN"

	guess := NewGuess(prmpt, solution)

	want := true
	got := guess.Won()
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestWonFails(t *testing.T) {
	solution := "AFFEN"
	prmpt := "ALTER"

	guess := NewGuess(prmpt, solution)

	want := false
	got := guess.Won()
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestNewWotdSucceeds(t *testing.T) {
	want := 5
	got := len(NewWotd())

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

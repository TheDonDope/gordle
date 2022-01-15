package validation

import (
	"fmt"
	"testing"
)

func MatchReturns5GreenSquaresOnCompleteCorrectGuess(t *testing.T) {
	solution := "JESUS"
	prmpt := "JESUS"

	guess := NewGuess(prmpt, solution)

	want := "🟩🟩🟩🟩🟩"
	got := guess.Match(solution)
	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

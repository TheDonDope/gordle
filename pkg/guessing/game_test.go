package guessing

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

func TestGreetPlayersSucceeds(t *testing.T) {
	want := "Welcome to 🟩🟨⬛ Gordle ⬛🟨🟩\n" +
		"You have 6 trys to guess the word of the day.\n" +
		"NOTE: The current implementation will pick a new word on every run!\n" +
		"🟩 means, the letter is in the word and in the correct spot.\n" +
		"🟨 means, that the letter is in the word but in the wrong spot.\n" +
		"⬛ means, that the letter is in not in the word in any spot.\n"

	got := captureOutput(func() {
		GreetPlayers()
	})

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

func TestPrintResultsSucceeds(t *testing.T) {
	g := NewGame()
	g.wotd = "yamls"
	prompt := "yamls"
	guess := NewGuess(prompt, g)
	g.guesses = append(g.guesses, guess)
	today := time.Now().Format("2006-01-02")

	want := "\nYour Gordle results (" + today + "):\n" +
		"🟩🟩🟩🟩🟩 (1/6): yamls\n" +
		"\nThe solution was: yamls\n"

	got := captureOutput(func() {
		g.PrintResults()
	})

	if got != want {
		t.Error(fmt.Printf("Should have matched, got: %v, want: %v", got, want))
	}
}

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

func TestRateGuessSucceedsOnWordle212(t *testing.T) {
	g := NewGame()
	g.wotd = "shire"

	todaysGuesses := []*Guess{NewGuess("alter", g), NewGuess("resin", g), NewGuess("fries", g), NewGuess("heirs", g), NewGuess("shire", g)}
	todaysRatings := []string{"⬛⬛⬛🟨🟨", "🟨🟨🟨🟨⬛", "⬛🟨🟩🟨🟨", "🟨🟨🟩🟩🟨", "🟩🟩🟩🟩🟩"}

	g.guesses = append(g.guesses, todaysGuesses...)

	for i := range todaysGuesses {
		for j := range todaysRatings {
			if i != j {
				continue
			}
			want := todaysRatings[i]

			got := todaysGuesses[i].Rating

			if got != want {
				t.Error(fmt.Printf("Should have matched, got: %v, want: %v\n", got, want))
				break
			}
		}
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

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

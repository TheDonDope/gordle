package storage

import (
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// Dictionary represents a collection of words.
type Dictionary struct {
	maxLength int
	allWords  []string
}

// NewDictionary returns a new, word filled dictionary.
func NewDictionary(maxLength int, dict io.Reader) *Dictionary {
	return &Dictionary{
		maxLength: maxLength,
		allWords:  readWords(dict),
	}
}

// NewWotd returns a new Word of the day.
func (d *Dictionary) NewWotd() string {
	// First of all, shuffle the whole existing word list...
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.allWords), func(i, j int) { d.allWords[i], d.allWords[j] = d.allWords[j], d.allWords[i] })

	// Pick the first word that matches the maxLength constraints
	for _, w := range d.allWords {
		if len(w) == d.maxLength && onlyAlpha(w) {
			return strings.ToLower(w)
		}
	}
	return "!word"
}

func readWords(dict io.Reader) (words []string) {
	parsed, _ := parseDict(dict)
	words = parsed
	return
}

func parseDict(handle io.Reader) ([]string, error) {
	words := []string{"NODIC"}
	bytes, err := ioutil.ReadAll(handle)
	if err != nil {
		return words, err
	}
	words = strings.Split(string(bytes), "\n")
	return words, nil
}

func onlyAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

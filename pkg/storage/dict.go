package storage

import (
	"io/ioutil"
	"math/rand"
	"os"
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
func NewDictionary(maxLength int) *Dictionary {
	return &Dictionary{
		maxLength: maxLength,
		allWords:  readWords(),
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

func onlyAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

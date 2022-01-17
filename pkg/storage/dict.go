package storage

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// NewWotd returns a new Word of the day.
func NewWotd() string {
	wotd := ""

	allWords := readWords()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allWords), func(i, j int) { allWords[i], allWords[j] = allWords[j], allWords[i] })

	for i := 0; len(wotd) != 5; i++ {
		w := allWords[i]
		if len(w) == 5 && onlyAlpha(w) {
			wotd = w
			return strings.ToLower(wotd)
		}
	}
	return wotd
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

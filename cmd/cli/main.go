package main

import (
	"os"

	game "github.com/TheDonDope/gordle/pkg/guessing"
)

func main() {
	dict, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return
	}
	defer dict.Close()

	g := game.NewGame(dict)
	game.GreetPlayers()
	g.Play()
	g.PrintResults()
}

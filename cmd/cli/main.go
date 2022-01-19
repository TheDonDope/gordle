package main

import (
	"os"

	"github.com/TheDonDope/gordle/pkg/game"
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

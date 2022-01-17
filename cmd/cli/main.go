package main

import (
	game "github.com/TheDonDope/gordle/pkg/guessing"
)

func main() {
	g := game.NewGame()
	game.GreetPlayers()
	g.Play()
	g.PrintResults()
}

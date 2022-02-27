package main

import (
	"blackjack/internal/game"
)

func main() {
	game := game.InitGame()
	game.Play()
}

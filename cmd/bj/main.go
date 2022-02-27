package main

import (
	"blackjagg/internal/game"
)

func main() {
	game := game.InitGame()
	game.Play()
}

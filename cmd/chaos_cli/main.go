package main

import (
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/game"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/ui"
)

func main() {
	game := game.NewChaosGame(&game.SierpinskiTriangle, 100, 100)
	game.RunSteps(100000)
	ui.PrintCanvas(game.Canvas.Canvas)
}

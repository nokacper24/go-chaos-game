package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/nokacper24/go-chaos-game/internal/chaosgame/game"
)

func main() {
	width := 1024
	height := 1024
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	game := game.NewChaosGame(&game.SierpinskiTriangle, width, height)
	game.RunSteps(100000000)

	canvas := game.Canvas.Canvas

	for i, row := range canvas {
		for j := range row {
			if canvas[i][j] {
				img.Set(j, i, color.RGBA{255,0,0,255})
			}
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(f, img)
}

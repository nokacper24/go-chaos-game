package game

import (
	"math/rand"

	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
)

type ChaosGame struct {
	description  *ChaosGameDescription
	currentPoint *model.Vector2D
	Canvas       *chaosCanvas
}

func NewChaosGame(description *ChaosGameDescription, width int, height int) *ChaosGame {
	return &ChaosGame{
		description:  description,
		currentPoint: &model.Vector2D{X: 0, Y: 0},
		Canvas:       NewChaosCanvas(width, height, description.MinCoords, description.MaxCoords),
	}
}

func (cg *ChaosGame) RunSteps(steps int) {
	numOfTransforms := len(cg.description.Transforms)
	for range steps {
		transform := cg.description.Transforms[rand.Intn(numOfTransforms)]
		cg.currentPoint = transform.Transform(cg.currentPoint)
		cg.Canvas.PutPixel(*cg.currentPoint)
	}
}

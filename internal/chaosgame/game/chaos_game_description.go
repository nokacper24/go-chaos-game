package game

import (
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/transform"
)

type ChaosGameDescription struct {
	MinCoords  model.Vector2D
	MaxCoords  model.Vector2D
	Transforms []transform.Transform2D
}

var (
	SierpinskiTriangle = ChaosGameDescription{
		MinCoords: model.Vector2D{X: 0, Y: 0},
		MaxCoords: model.Vector2D{X: 1, Y: 1},
		Transforms: []transform.Transform2D{
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0.5, A01: 0, A10: 0, A11: 0.5},
				Vector: model.Vector2D{X: 0, Y: 0}},
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0.5, A01: 0, A10: 0, A11: 0.5},
				Vector: model.Vector2D{X: 0.5, Y: 0}},
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0.5, A01: 0, A10: 0, A11: 0.5},
				Vector: model.Vector2D{X: 0.25, Y: 0.5}},
		},
	}
)

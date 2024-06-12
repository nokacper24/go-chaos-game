package game

import (
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model/sign"
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

	BarnsleyFern = ChaosGameDescription{
		MinCoords: model.Vector2D{X: 0, Y: 0},
		MaxCoords: model.Vector2D{X: 1, Y: 1},
		Transforms: []transform.Transform2D{
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0, A01: 0, A10: 0, A11: 0.16},
				Vector: model.Vector2D{X: 0, Y: 0}},
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0.85, A01: 0.04, A10: -0.04, A11: 0.85},
				Vector: model.Vector2D{X: 0, Y: 1.6}},
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: 0.2, A01: -0.26, A10: 0.23, A11: 0.22},
				Vector: model.Vector2D{X: 0, Y: 0.16}},
			&transform.AffineTransform2D{
				Matrix: model.Matrix2x2{A00: -0.15, A01: 0.28, A10: 0.26, A11: 0.24},
				Vector: model.Vector2D{X: 0, Y: 0.44}},
		},
	}

	JuliaTransform = ChaosGameDescription{
		MinCoords: model.Vector2D{X: -1.6, Y: -1},
		MaxCoords: model.Vector2D{X: 1.6, Y: 1},
		Transforms: []transform.Transform2D{
			&transform.JuliaTransform{
				Point: model.Vector2D{X: -0.74543, Y: 0.11301},
				Sign:  sign.MINUS,
			},
		},
	}
)

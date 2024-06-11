package transform

import (
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model/sign"
)

type Transform2D interface {
	Transform(point *model.Vector2D) *model.Vector2D
}

type AffineTransform2D struct {
	Matrix model.Matrix2x2
	Vector model.Vector2D
}

func (t *AffineTransform2D) Transform(point *model.Vector2D) *model.Vector2D {
	return t.Matrix.Multiply(point).Add(&t.Vector)
}

type JuliaTransform struct {
	Point model.Vector2D
	Sign  sign.Sign
}

func (t *JuliaTransform) Transform(point *model.Vector2D) *model.Vector2D {
	return point.
		Substract(&t.Point).
		Sqrt().
		Multiply(float64(t.Sign))
}

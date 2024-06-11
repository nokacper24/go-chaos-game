package transform

import (
	"math"
	"testing"

	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model"
	"github.com/nokacper24/go-chaos-game/internal/chaosgame/model/sign"
)

func TestAffineTransform(t *testing.T) {
	affine := AffineTransform2D{
		model.Matrix2x2{0.5, 1, 1, 0.5},
		model.Vector2D{3, 1},
	}
	x := model.Vector2D{1, 2}

	result := affine.Transform(&x)

	expected := model.Vector2D{5.5, 3}
	if result.X != expected.X {
		t.Fatalf("Expected %v to be equal %v", result.X, expected.X)
	}

	if result.Y != expected.Y {
		t.Fatalf("Expected %v to be equal %v", result.Y, expected.Y)
	}
}

func TestJuliaTransform(t *testing.T) {
	z := model.Vector2D{0.4, 0.2}

	julia := JuliaTransform{
		model.Vector2D{0.3, 0.6},
		sign.PLUS,
	}

	result := julia.Transform(&z)
	result.X = math.Round(result.X*1000) / 1000
	result.Y = math.Round(result.Y*1000) / 1000

	expected := model.Vector2D{0.506, -0.395}
	if result.X != expected.X {
		t.Fatalf("Expected %v to be equal %v", result.X, expected.X)
	}

	if result.Y != expected.Y {
		t.Fatalf("Expected %v to be equal %v", result.Y, expected.Y)
	}
}

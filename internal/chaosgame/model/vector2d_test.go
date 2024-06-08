package model

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{3, 4}

	result := v1.Add(&v2)

	if result.X != 4 {
		t.Fatalf("Expected %v to be equal %v", result.X, 4)
	}

	if result.Y != 6 {
		t.Fatalf("Expected %v to be equal %v", result.Y, 6)
	}
}

func TestAddNegative(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{-3, -4}

	result := v1.Add(&v2)

	if result.X != -2 {
		t.Fatalf("Expected %v to be equal %v", result.X, -2)
	}

	if result.Y != -2 {
		t.Fatalf("Expected %v to be equal %v", result.Y, -2)
	}
}

func TestSubstract(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{3, 4}

	result := v1.Substract(&v2)

	if result.X != -2 {
		t.Fatalf("Expected %v to be equal %v", result.X, -2)
	}

	if result.Y != -2 {
		t.Fatalf("Expected %v to be equal %v", result.Y, -2)
	}
}

func TestSustractNegative(t *testing.T) {
	v1 := Vector2D{1, 2}
	v2 := Vector2D{-3, -4}

	result := v1.Substract(&v2)

	if result.X != 4 {
		t.Fatalf("Expected %v to be equal %v", result.X, 4)
	}

	if result.Y != 6 {
		t.Fatalf("Expected %v to be equal %v", result.Y, 6)
	}
}

func TestSqrt(t *testing.T) {
	c := Vector2D{0.1, -0.4}
	result := c.Sqrt()
	rounded0 := math.Round(result.X*1000) / 1000
	rounded1 := math.Round(result.Y*1000) / 1000

	expected := Vector2D{0.506, -0.395}

	if rounded0 != expected.X {
		t.Fatalf("Expected %v to be equal %v", rounded0, expected.X)
	}

	if rounded1 != expected.Y {
		t.Fatalf("Expected %v to be equal %v", rounded1, expected.Y)
	}
}

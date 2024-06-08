package model

import "math"

// Represents a pair of numbers. Can be used to represent a point, a vector, or a complex number.
type Vector2D struct {
	X, Y float64
}

// Adds other Vector2D to this Vector2D. Returns a poitner to a new.
func (v *Vector2D) Add(other *Vector2D) *Vector2D {
	return &Vector2D{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// Substracts other Vector2D from this Vector2D. Returns a poitner to a new vector.
func (v *Vector2D) Substract(other *Vector2D) *Vector2D {
	return &Vector2D{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// Calculates the complex square root of of the given vector2d (pair of numbers, x is the real part, y is the imaginary part)
// Returns a poitner to a new vector.
func (c *Vector2D) Sqrt() *Vector2D {
	realPart := math.Sqrt(0.5 * (math.Sqrt(c.X*c.X+(c.Y*c.Y)) + c.X))

	signY := math.Copysign(1, c.Y)
	complexPart := signY * math.Sqrt(0.5*(math.Sqrt(c.X*c.X+c.Y*c.Y)-c.X))

	return &Vector2D{realPart, complexPart}
}

// Multiplies this vector by a scalar. Returns a pointer to a new vector.
func (v *Vector2D) Multiply(scalar float64) *Vector2D {
	return &Vector2D{
		v.X * scalar,
		v.Y * scalar,
	}
}

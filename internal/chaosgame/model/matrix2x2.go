package model

type Matrix2x2 struct {
	a00, a01,
	a10, a11 float64
}

// Rightmultiples this matrix by the given Vector2D.
// Given Vector2D is consumed and a pointer to a new one is returned.
func (m *Matrix2x2) Multiply(v *Vector2D) *Vector2D {
	result := Vector2D{
		m.a00*v.X + m.a01*v.Y,
		m.a10*v.X + m.a11*v.Y,
	}
	return &result
}

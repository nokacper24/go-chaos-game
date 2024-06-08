package model

type Matrix2x2 struct {
	A00, A01,
	A10, A11 float64
}

// Rightmultiples this matrix by the given Vector2D.
// Given Vector2D is consumed and a pointer to a new one is returned.
func (m *Matrix2x2) Multiply(v *Vector2D) *Vector2D {
	result := Vector2D{
		m.A00*v.X + m.A01*v.Y,
		m.A10*v.X + m.A11*v.Y,
	}
	return &result
}

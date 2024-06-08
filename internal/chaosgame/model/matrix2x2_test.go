package model

import "testing"

func TestMultiply(t *testing.T) {
	m := Matrix2x2{
		0.5, 1,
		1, 0.5,
	}
	v := Vector2D{1, 2}

	result := m.Multiply(&v)
	expected := Vector2D{2.5, 2}

	if *result != expected {
		t.Fatalf("Expected %v to be equal %v", *result, expected)
	}
}

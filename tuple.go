package main

// Tuple represents a 3D point (W=1.0) or a vector (W=0.0)
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewPoint creates a new point represented by a 3D Tuple
func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

// NewVector creates a new vector represented by a 3D Tuple
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

// IsPoint if tuple is a point
func (t *Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector if tuple is a vector
func (t *Tuple) IsVector() bool {
	return t.W == 0.0
}

package core

import "math"

// Normal3 is a 3D vector that is perpendicular to a surface.
type Normal3 struct {
	X, Y, Z float64
}

// Neg returns the negation of a normal.
func (n Normal3) Neg() Normal3 {
	return Normal3{-n.X, -n.Y, -n.Z}
}

// Add returns the sum of two normals.
func (n Normal3) Add(other Normal3) Normal3 {
	return Normal3{n.X + other.X, n.Y + other.Y, n.Z + other.Z}
}

// Sub returns the difference of two normals.
func (n Normal3) Sub(other Normal3) Normal3 {
	return Normal3{n.X - other.X, n.Y - other.Y, n.Z - other.Z}
}

// Mul returns the scalar multiplication of a normal and a scalar.
func (n Normal3) Mul(s float64) Normal3 {
	return Normal3{n.X * s, n.Y * s, n.Z * s}
}

// Div returns the scalar division of a normal and a scalar.
func (n Normal3) Div(s float64) Normal3 {
	return Normal3{n.X / s, n.Y / s, n.Z / s}
}

// Mag returns the magnitude of a normal.
func (n Normal3) Mag() float64 {
	return math.Sqrt(n.MagSquared())
}

// MagSquared returns the magnitude squared of a normal.
func (n Normal3) MagSquared() float64 {
	return n.X*n.X + n.Y*n.Y + n.Z*n.Z
}

// Unit returns a unit normal pointing in the same direction as the original normal.
func (n Normal3) Unit() Normal3 {
	return n.Div(n.Mag())
}

// Equals returns true if the two normals are equal.
func (n Normal3) Equals(other Normal3) bool {
	return n.X == other.X && n.Y == other.Y && n.Z == other.Z
}

// FaceForward returns a normal that lies in the same hemisphere as the given vector.
// The result is equivalent to either the original normal or its negation.
func (n Normal3) FaceForward(v Vector3) Normal3 {
	if Vector3(n).Dot(v) < 0 {
		return n.Neg()
	}
	return n
}

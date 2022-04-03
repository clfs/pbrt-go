package core

import "math"

// Vector2 is a 2D vector.
type Vector2 struct {
	X, Y float64
}

// Add returns the sum of two vectors.
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{v.X + other.X, v.Y + other.Y}
}

// Sub returns the difference of two vectors.
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

// Equals returns true if the two vectors are equal.
func (v Vector2) Equals(other Vector2) bool {
	return v.X == other.X && v.Y == other.Y
}

// Mul returns the scalar multiplication of a vector and a scalar.
func (v Vector2) Mul(s float64) Vector2 {
	return Vector2{v.X * s, v.Y * s}
}

// Div returns the scalar division of a vector and a scalar.
func (v Vector2) Div(s float64) Vector2 {
	return Vector2{v.X / s, v.Y / s}
}

// Neg returns the negation of a vector.
func (v Vector2) Neg() Vector2 {
	return Vector2{-v.X, -v.Y}
}

// Mag returns the magnitude of a vector.
func (v Vector2) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MagSquared returns the magnitude of a vector squared.
func (v Vector2) MagSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Vector3 is a 3D vector.
type Vector3 struct {
	X, Y, Z float64
}

// Add returns the sum of two vectors.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub returns the difference of two vectors.
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Equals returns true if the two vectors are equal.
func (v Vector3) Equals(other Vector3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

// Mul returns the scalar multiplication of a vector and a scalar.
func (v Vector3) Mul(s float64) Vector3 {
	return Vector3{v.X * s, v.Y * s, v.Z * s}
}

// Div returns the scalar division of a vector and a scalar.
func (v Vector3) Div(s float64) Vector3 {
	return Vector3{v.X / s, v.Y / s, v.Z / s}
}

// Neg returns the negation of a vector.
func (v Vector3) Neg() Vector3 {
	return Vector3{-v.X, -v.Y, -v.Z}
}

// Mag returns the magnitude of a vector.
func (v Vector3) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagSquared returns the magnitude of a vector squared.
func (v Vector3) MagSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Abs returns a vector representing the component-wise absolute values of the
// original vector.
func (v Vector3) Abs() Vector3 {
	return Vector3{math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)}
}

// Dot returns the dot product of two vectors.
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product of two vectors.
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{v.Y*other.Z - v.Z*other.Y, v.Z*other.X - v.X*other.Z, v.X*other.Y - v.Y*other.X}
}

// Unit returns a unit vector pointing in the same direction as the original.
// It panics if the original vector is zero.
func (v Vector3) Unit() Vector3 {
	return v.Div(v.Mag())
}

// MinComponent returns the minimum component of a vector.
func (v Vector3) MinComponent() float64 {
	return math.Min(v.X, math.Min(v.Y, v.Z))
}

// MaxComponent returns the maximum component of a vector.
func (v Vector3) MaxComponent() float64 {
	return math.Max(v.X, math.Max(v.Y, v.Z))
}

// Min returns a vector representing the component-wise minimums of two vectors.
func (v Vector3) Min(other Vector3) Vector3 {
	return Vector3{math.Min(v.X, other.X), math.Min(v.Y, other.Y), math.Min(v.Z, other.Z)}
}

// Min returns a vector representing the component-wise maximums of two vectors.
func (v Vector3) Max(other Vector3) Vector3 {
	return Vector3{math.Max(v.X, other.X), math.Max(v.Y, other.Y), math.Max(v.Z, other.Z)}
}

// CoordinateSystem returns two unit vectors that form a right-handed coordinate
// system with the original vector as the x-axis.
func (v Vector3) CoordinateSystem() (j Vector3, k Vector3) {
	unit := v.Unit()
	if math.Abs(v.X) > math.Abs(v.Y) {
		j = Vector3{-unit.Z, 0, unit.X}.Unit()
	} else {
		j = Vector3{0, unit.Z, -unit.Y}.Unit()
	}
	k = unit.Cross(j)
	return
}

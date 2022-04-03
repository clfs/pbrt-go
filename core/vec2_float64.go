package core

import "math"

// Vec2Float64 is a 2-dimensional vector of float64 values.
type Vec2Float64 struct {
	X, Y float64
}

// Add returns the sum of two vectors.
func (v Vec2Float64) Add(other Vec2Float64) Vec2Float64 {
	return Vec2Float64{v.X + other.X, v.Y + other.Y}
}

// Sub returns the difference of two vectors.
func (v Vec2Float64) Sub(other Vec2Float64) Vec2Float64 {
	return Vec2Float64{v.X - other.X, v.Y - other.Y}
}

// Eq returns true if the vectors are equal.
func (v Vec2Float64) Eq(other Vec2Float64) bool {
	return v.X == other.X && v.Y == other.Y
}

// SMul returns the scalar multiplication of a vector and a scalar.
func (v Vec2Float64) SMul(s float64) Vec2Float64 {
	return Vec2Float64{v.X * s, v.Y * s}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec2Float64) SDiv(s float64) Vec2Float64 {
	return Vec2Float64{v.X / s, v.Y / s}
}

// Neg returns the negation of a vector.
func (v Vec2Float64) Neg() Vec2Float64 {
	return Vec2Float64{-v.X, -v.Y}
}

// Mag returns the magnitude of a vector.
func (v Vec2Float64) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MagSq returns the squared magnitude of a vector.
func (v Vec2Float64) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

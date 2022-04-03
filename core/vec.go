package core

import "math"

// Vec2Int64 is a 2-dimensional vector.
type Vec2Int64 struct {
	X, Y int64
}

// Add returns the sum of two vectors.
func (v Vec2Int64) Add(other Vec2Int64) Vec2Int64 {
	return Vec2Int64{v.X + other.X, v.Y + other.Y}
}

// Sub returns the difference of two vectors.
func (v Vec2Int64) Sub(other Vec2Int64) Vec2Int64 {
	return Vec2Int64{v.X - other.X, v.Y - other.Y}
}

// Eq returns true if the vectors are equal.
func (v Vec2Int64) Eq(other Vec2Int64) bool {
	return v.X == other.X && v.Y == other.Y
}

// SMul returns the scalar multiplication of a vector and a scalar.
func (v Vec2Int64) SMul(s int64) Vec2Int64 {
	return Vec2Int64{v.X * s, v.Y * s}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec2Int64) SDiv(s int64) Vec2Int64 {
	return Vec2Int64{v.X / s, v.Y / s}
}

// Neg returns the negation of a vector.
func (v Vec2Int64) Neg() Vec2Int64 {
	return Vec2Int64{-v.X, -v.Y}
}

// Mag returns the magnitude of a vector.
func (v Vec2Int64) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

// MagSq returns the squared magnitude of a vector.
func (v Vec2Int64) MagSq() int64 {
	return v.X*v.X + v.Y*v.Y
}

// Vec2Float64 is a 2-dimensional vector.
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

// Vec3Int64 is a 3-dimensional vector.
type Vec3Int64 struct {
	X, Y, Z int64
}

// Add returns the sum of two vectors.
func (v Vec3Int64) Add(other Vec3Int64) Vec3Int64 {
	return Vec3Int64{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub returns the difference of two vectors.
func (v Vec3Int64) Sub(other Vec3Int64) Vec3Int64 {
	return Vec3Int64{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Eq returns true if the vectors are equal.
func (v Vec3Int64) Eq(other Vec3Int64) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

// SMul returns the scalar multiplication of a vector and a scalar.
func (v Vec3Int64) SMul(s int64) Vec3Int64 {
	return Vec3Int64{v.X * s, v.Y * s, v.Z * s}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec3Int64) SDiv(s int64) Vec3Int64 {
	return Vec3Int64{v.X / s, v.Y / s, v.Z / s}
}

// Neg returns the negation of a vector.
func (v Vec3Int64) Neg() Vec3Int64 {
	return Vec3Int64{-v.X, -v.Y, -v.Z}
}

// Mag returns the magnitude of a vector.
func (v Vec3Int64) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
}

// MagSq returns the squared magnitude of a vector.
func (v Vec3Int64) MagSq() int64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Vec3Float64 is a 3-dimensional vector.
type Vec3Float64 struct {
	X, Y, Z float64
}

// Add returns the sum of two vectors.
func (v Vec3Float64) Add(other Vec3Float64) Vec3Float64 {
	return Vec3Float64{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub returns the difference of two vectors.
func (v Vec3Float64) Sub(other Vec3Float64) Vec3Float64 {
	return Vec3Float64{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Eq returns true if the vectors are equal.
func (v Vec3Float64) Eq(other Vec3Float64) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

// SMul returns the scalar multiplication of a vector and a scalar.
func (v Vec3Float64) SMul(s float64) Vec3Float64 {
	return Vec3Float64{v.X * s, v.Y * s, v.Z * s}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec3Float64) SDiv(s float64) Vec3Float64 {
	return Vec3Float64{v.X / s, v.Y / s, v.Z / s}
}

// Neg returns the negation of a vector.
func (v Vec3Float64) Neg() Vec3Float64 {
	return Vec3Float64{-v.X, -v.Y, -v.Z}
}

// Mag returns the magnitude of a vector.
func (v Vec3Float64) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// MagSq returns the squared magnitude of a vector.
func (v Vec3Float64) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

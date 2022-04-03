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
func (v Vec2Int64) SMul(s float64) Vec2Int64 {
	return Vec2Int64{
		int64(float64(v.X) * s),
		int64(float64(v.Y) * s),
	}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec2Int64) SDiv(s float64) Vec2Int64 {
	return Vec2Int64{
		int64(float64(v.X) / s),
		int64(float64(v.Y) / s),
	}
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
func (v Vec3Int64) SMul(s float64) Vec3Int64 {
	return Vec3Int64{
		int64(float64(v.X) * s),
		int64(float64(v.Y) * s),
		int64(float64(v.Z) * s),
	}
}

// SDiv returns the scalar division of a vector and a scalar. It panics if s is
// zero.
func (v Vec3Int64) SDiv(s float64) Vec3Int64 {
	return Vec3Int64{
		int64(float64(v.X) / s),
		int64(float64(v.Y) / s),
		int64(float64(v.Z) / s),
	}
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
func (v Vec3Int64) MagSq() float64 {
	return float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func absInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Abs returns a vector with the absolute value operation applied to its
// components.
func (v Vec3Int64) Abs() Vec3Int64 {
	return Vec3Int64{absInt64(v.X), absInt64(v.Y), absInt64(v.Z)}
}

// Dot returns the dot product of two vectors.
func (v Vec3Int64) Dot(other Vec3Int64) int64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// AbsDot returns the absolute value of the dot product of two vectors.
func (v Vec3Int64) AbsDot(other Vec3Int64) int64 {
	return absInt64(v.X*other.X + v.Y*other.Y + v.Z*other.Z)
}

// Cross returns the cross product of two vectors. The underlying math is
// performed using float64 types to avoid catastrophic cancellation.
func (v Vec3Int64) Cross(other Vec3Int64) Vec3Int64 {
	vX, vY, vZ := float64(v.X), float64(v.Y), float64(v.Z)
	oX, oY, oZ := float64(other.X), float64(other.Y), float64(other.Z)
	return Vec3Int64{
		int64(vY*oZ - vZ*oY),
		int64(vZ*oX - vX*oZ),
		int64(vX*oY - vY*oX),
	}
}

// Unit returns a unit vector pointing in the same direction as the original.
// It panics if the original vector is zero.
func (v Vec3Int64) Unit() Vec3Int64 {
	return v.SDiv(v.Mag())
}

// MinComp returns the minimum component of a vector.
func (v Vec3Int64) MinComp() int64 {
	if v.X < v.Y && v.X < v.Z {
		return v.X
	}
	if v.Y < v.Z {
		return v.Y
	}
	return v.Z
}

// MaxComp returns the maximum component of a vector.
func (v Vec3Int64) MaxComp() int64 {
	if v.X > v.Y && v.X > v.Z {
		return v.X
	}
	if v.Y > v.Z {
		return v.Y
	}
	return v.Z
}

// MinDim returns the dimension of the component with the smallest value.
// Possible return values are 0, 1, or 2.
func (v Vec3Int64) MinDim() int {
	if v.X < v.Y && v.X < v.Z {
		return 0
	}
	if v.Y < v.Z {
		return 1
	}
	return 2
}

// MaxDim returns the dimension of the component with the largest value.
// Possible return values are 0, 1, or 2.
func (v Vec3Int64) MaxDim() int {
	if v.X > v.Y && v.X > v.Z {
		return 0
	}
	if v.Y > v.Z {
		return 1
	}
	return 2
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// Min returns a vector with the minimum components of the two original vectors.
func (v Vec3Int64) Min(other Vec3Int64) Vec3Int64 {
	return Vec3Int64{
		minInt64(v.X, other.X),
		minInt64(v.Y, other.Y),
		minInt64(v.Z, other.Z),
	}
}

// Max returns a vector with the maximum components of the two original vectors.
func (v Vec3Int64) Max(other Vec3Int64) Vec3Int64 {
	return Vec3Int64{
		maxInt64(v.X, other.X),
		maxInt64(v.Y, other.Y),
		maxInt64(v.Z, other.Z),
	}
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

// Abs returns a vector with the absolute value operation applied to its
// components.
func (v Vec3Float64) Abs() Vec3Float64 {
	return Vec3Float64{math.Abs(v.X), math.Abs(v.Y), math.Abs(v.Z)}
}

// Dot returns the dot product of two vectors.
func (v Vec3Float64) Dot(other Vec3Float64) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// AbsDot returns the absolute value of the dot product of two vectors.
func (v Vec3Float64) AbsDot(other Vec3Float64) float64 {
	return math.Abs(v.X*other.X + v.Y*other.Y + v.Z*other.Z)
}

// Cross returns the cross product of two vectors.
func (v Vec3Float64) Cross(other Vec3Float64) Vec3Float64 {
	return Vec3Float64{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// Unit returns a unit vector pointing in the same direction as the original.
// It panics if the original vector is zero.
func (v Vec3Float64) Unit() Vec3Float64 {
	return v.SDiv(v.Mag())
}

// MinComp returns the minimum component of a vector.
func (v Vec3Float64) MinComp() float64 {
	if v.X < v.Y && v.X < v.Z {
		return v.X
	}
	if v.Y < v.Z {
		return v.Y
	}
	return v.Z
}

// MaxComp returns the maximum component of a vector.
func (v Vec3Float64) MaxComp() float64 {
	if v.X > v.Y && v.X > v.Z {
		return v.X
	}
	if v.Y > v.Z {
		return v.Y
	}
	return v.Z
}

// MinDim returns the dimension of the component with the smallest value.
// Possible return values are 0, 1, or 2.
func (v Vec3Float64) MinDim() int {
	if v.X < v.Y && v.X < v.Z {
		return 0
	}
	if v.Y < v.Z {
		return 1
	}
	return 2
}

// MaxDim returns the dimension of the component with the largest value.
// Possible return values are 0, 1, or 2.
func (v Vec3Float64) MaxDim() int {
	if v.X > v.Y && v.X > v.Z {
		return 0
	}
	if v.Y > v.Z {
		return 1
	}
	return 2
}

// Min returns a vector with the minimum components of the two original vectors.
func (v Vec3Float64) Min(other Vec3Float64) Vec3Float64 {
	return Vec3Float64{
		math.Min(v.X, other.X),
		math.Min(v.Y, other.Y),
		math.Min(v.Z, other.Z),
	}
}

// Max returns a vector with the maximum components of the two original vectors.
func (v Vec3Float64) Max(other Vec3Float64) Vec3Float64 {
	return Vec3Float64{
		math.Max(v.X, other.X),
		math.Max(v.Y, other.Y),
		math.Max(v.Z, other.Z),
	}
}

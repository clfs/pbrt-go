package core

import "math"

// Point2 is a 2D point.
type Point2 struct {
	X, Y float64
}

// Add returns the sum of two points.
func (p Point2) Add(other Point2) Point2 {
	return Point2{p.X + other.X, p.Y + other.Y}
}

// Sub returns the vector between two points.
func (p Point2) Sub(other Point2) Vector2 {
	return Vector2{p.X - other.X, p.Y - other.Y}
}

// Equals returns true if the two points are equal.
func (p Point2) Equals(other Point2) bool {
	return p.X == other.X && p.Y == other.Y
}

// Mul returns the scalar multiplication of a point and a scalar.
func (p Point2) Mul(s float64) Point2 {
	return Point2{p.X * s, p.Y * s}
}

// Div returns the scalar division of a point and a scalar.
func (p Point2) Div(s float64) Point2 {
	return Point2{p.X / s, p.Y / s}
}

// Point3 is a 3D point.
type Point3 struct {
	X, Y, Z float64
}

// Add returns the sum of two points.
func (p Point3) Add(other Point3) Point3 {
	return Point3{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

// Sub returns the vector between two points.
func (p Point3) Sub(other Point3) Vector3 {
	return Vector3{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}

// Equals returns true if the two points are equal.
func (p Point3) Equals(other Point3) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

// Mul returns the scalar multiplication of a point and a scalar.
func (p Point3) Mul(s float64) Point3 {
	return Point3{p.X * s, p.Y * s, p.Z * s}
}

// Div returns the scalar division of a point and a scalar.
func (p Point3) Div(s float64) Point3 {
	return Point3{p.X / s, p.Y / s, p.Z / s}
}

// ToPoint2 converts a 3D point to a 2D point by dropping the z coordinate.
func (p Point3) ToPoint2() Point2 {
	return Point2{p.X, p.Y}
}

// AddVec returns a point offset from the original point by the given vector.
func (p Point3) AddVec(v Vector3) Point3 {
	return Point3{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// Distance returns the distance between two points.
func (p Point3) Distance(other Point3) float64 {
	return p.Sub(other).Mag()
}

// DistanceSquared returns the distance between two points squared.
func (p Point3) DistanceSquared(other Point3) float64 {
	return p.Sub(other).MagSquared()
}

// Lerp returns the linear interpolation between two points.
// For t < 0 or t > 1, Lerp extrapolates.
func (p Point3) Lerp(other Point3, t float64) Point3 {
	return p.Mul(1 - t).Add(other.Mul(t))
}

// Min returns a point representing the component-wise minimums of two points.
func (p Point3) Min(other Point3) Point3 {
	return Point3{
		math.Min(p.X, other.X),
		math.Min(p.Y, other.Y),
		math.Min(p.Z, other.Z),
	}
}

// Max returns a point representing the component-wise maximums of two points.
func (p Point3) Max(other Point3) Point3 {
	return Point3{
		math.Max(p.X, other.X),
		math.Max(p.Y, other.Y),
		math.Max(p.Z, other.Z),
	}
}

// Floor returns a point representing the component-wise floors of the original
// point.
func (p Point3) Floor() Point3 {
	return Point3{
		math.Floor(p.X),
		math.Floor(p.Y),
		math.Floor(p.Z),
	}
}

// Ceil returns a point representing the component-wise ceilings of the original
// point.
func (p Point3) Ceil() Point3 {
	return Point3{
		math.Ceil(p.X),
		math.Ceil(p.Y),
		math.Ceil(p.Z),
	}
}

// Abs returns a point representing the component-wise absolute values of the
// original point.
func (p Point3) Abs() Point3 {
	return Point3{
		math.Abs(p.X),
		math.Abs(p.Y),
		math.Abs(p.Z),
	}
}

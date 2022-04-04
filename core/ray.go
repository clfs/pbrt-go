package core

// Ray is a semi-infinite line.
type Ray struct {
	Origin    Point3
	Direction Vector3
	T, TMax   float64
	Medium    *Medium
}

// At returns the point at t along the ray.
func (r Ray) At(t float64) Point3 {
	return r.Origin.AddVec(r.Direction.Mul(t))
}

// RayDifferential represents a ray with additional information about two
// auxiliary rays. These extra rays represent camera rays offset by one sample
// in the x and y direction from the main ray on the film plane.
type RayDifferential struct {
	Ray
	HasDifferentials         bool
	RxOrigin, RyOrigin       Point3
	RxDirection, RyDirection Vector3
}

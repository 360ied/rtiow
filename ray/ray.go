package ray

import (
	"math"

	"rtiow/vec3"
)

type Ray struct {
	Origin    vec3.Point3
	Direction vec3.Vec3
}

func (r Ray) At(t float64) vec3.Point3 {
	return r.Origin.AddVec3(r.Direction.MultiplyFloat(t))
}

// Linearly blends white and blue depending on the height of the y coordinate after scaling the ray direction to unit length
func (r Ray) Colour() vec3.Colour {
	t := r.HitSphere(vec3.Point3{Z: -1}, .5)
	if t > 0.0 {
		N := r.At(t).SubtractVec3(vec3.Vec3{Z: -1}).UnitVector()
		return vec3.Colour{X: N.X + 1, Y: N.Y + 1, Z: N.Z + 1}.MultiplyFloat(.5)
	}
	unitDirection := r.Direction.UnitVector()
	t = .5 * (unitDirection.Y + 1.0)
	return vec3.Colour{
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}.MultiplyFloat(1.0 - t).AddVec3(
		vec3.Colour{
			X: 0.5,
			Y: 0.7,
			Z: 1.0,
		}.MultiplyFloat(t))
}

// Hardcoded sphere
func (r Ray) HitSphere(center vec3.Point3, radius float64) float64 {
	oc := r.Origin.SubtractVec3(center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}

package ray

import (
	"math"

	"rtiow/constants"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

type Ray struct {
	Origin    vec3.Point3
	Direction vec3.Vec3
}

func (r Ray) At(t float64) vec3.Point3 {
	return r.Origin.AddVec3(r.Direction.MultiplyFloat(t))
}

// Linearly blends white and blue depending on the height of the y coordinate after scaling the ray direction to unit length
func (r Ray) Colour(world Hittable, depth int) vec3.Colour {
	// t := r.HitSphere(vec3.Point3{Z: -1}, .5)
	// if t > 0.0 {
	// 	N := r.At(t).SubtractVec3(vec3.Vec3{Z: -1}).UnitVector()
	// 	return vec3.Colour{X: N.X + 1, Y: N.Y + 1, Z: N.Z + 1}.MultiplyFloat(.5)
	// }
	rec, worldHit := world.Hit(r, 0.001, constants.PositiveInfinity)
	if depth <= 0 {
		return vec3.Colour{} // 0, 0, 0
	}
	if worldHit {
		// return rec.Normal.AddVec3(vec3.Colour{X: 1, Y: 1, Z: 1}).MultiplyFloat(.5)
		// target := rec.P.AddVec3(rec.Normal).AddVec3(vec3util.RandomInUnitSphere())
		target := rec.P.AddVec3(vec3util.RandomInHemisphere(rec.Normal))
		return Ray{rec.P, target.SubtractVec3(rec.P)}.Colour(world, depth-1).MultiplyFloat(.5)
	}
	unitDirection := r.Direction.UnitVector()
	t := .5 * (unitDirection.Y + 1.0)
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
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

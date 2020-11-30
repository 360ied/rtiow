package ray

import "rtiow/vec3"

type Ray struct {
	Origin    vec3.Point3
	Direction vec3.Vec3
}

func (r Ray) At(t float64) vec3.Point3 {
	return r.Origin.AddVec3(r.Direction.MultiplyFloat(t))
}

// Linearly blends white and blue depending on the height of the y coordinate after scaling the ray direction to unit length
func (r Ray) Colour() vec3.Colour {
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

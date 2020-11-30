package camera

import (
	"rtiow/material"
	"rtiow/vec3"
)

type Camera struct {
	AspectRatio    float64
	ViewportHeight float64
	ViewportWidth  float64
	FocalLength    float64

	Origin          vec3.Point3
	Horizontal      vec3.Vec3
	Vertical        vec3.Vec3
	LowerLeftCorner vec3.Vec3
}

func NewCamera(
	aspectRatio, viewportHeight, focalLength float64,
	origin vec3.Point3,
	horizontal, vertical vec3.Vec3,
) Camera {
	return Camera{
		AspectRatio:    aspectRatio,
		ViewportHeight: viewportHeight,
		ViewportWidth:  aspectRatio * viewportHeight,
		FocalLength:    focalLength,
		Origin:         origin,
		Horizontal:     horizontal,
		Vertical:       vertical,
		LowerLeftCorner: origin.SubtractVec3(
			horizontal.DivideFloat(2.0),
		).SubtractVec3(
			vertical.DivideFloat(2.0),
		).SubtractVec3(vec3.Vec3{Z: focalLength}), // 0, 0, focalLength,
	}
}

func (c Camera) Ray(u, v float64) material.Ray {
	return material.Ray{
		Origin: c.Origin,
		Direction: c.LowerLeftCorner.AddVec3(
			c.Horizontal.MultiplyFloat(u),
		).AddVec3(
			c.Vertical.MultiplyFloat(v),
		).SubtractVec3(c.Origin),
	}
}

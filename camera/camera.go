package camera

import (
	"math"

	"rtiow/helpers"
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
	// aspectRatio, viewportHeight, focalLength float64,
	// origin vec3.Point3,
	// horizontal, vertical vec3.Vec3,
	vFov, aspectRatio float64,
) Camera {
	theta := helpers.DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	const focalLength = 1.0

	origin := vec3.Point3{} // 0, 0, 0
	horizontal := vec3.Vec3{X: viewportWidth}
	vertical := vec3.Vec3{Y: viewportHeight}

	return Camera{
		AspectRatio:    aspectRatio,
		ViewportHeight: viewportHeight,
		ViewportWidth:  viewportWidth,
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

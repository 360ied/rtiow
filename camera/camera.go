package camera

import (
	"math"

	"rtiow/helpers"
	"rtiow/material"
	"rtiow/vec3"
)

type Camera struct {
	Origin          vec3.Point3
	Horizontal      vec3.Vec3
	Vertical        vec3.Vec3
	LowerLeftCorner vec3.Vec3
}

func NewCamera(
	lookFrom, lookAt vec3.Point3,
	vUp vec3.Vec3,
	vFov, aspectRatio float64,
) (cam Camera) {
	theta := helpers.DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.SubtractVec3(lookAt).UnitVector()
	u := vUp.Cross(w).UnitVector()
	v := w.Cross(u)

	cam.Origin = lookFrom
	cam.Horizontal = u.MultiplyFloat(viewportWidth)
	cam.Vertical = v.MultiplyFloat(viewportHeight)
	cam.LowerLeftCorner = cam.Origin.
		SubtractVec3(cam.Horizontal.DivideFloat(2.0)).
		SubtractVec3(cam.Vertical.DivideFloat(2.0)).
		SubtractVec3(w)
	return
}

func (c Camera) Ray(s, t float64) material.Ray {
	return material.Ray{
		Origin: c.Origin,
		Direction: c.LowerLeftCorner.
			AddVec3(c.Horizontal.MultiplyFloat(s)).
			AddVec3(c.Vertical.MultiplyFloat(t)).
			SubtractVec3(c.Origin),
	}
}

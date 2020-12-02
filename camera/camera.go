package camera

import (
	"math"

	"rtiow/helpers"
	"rtiow/material"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

type Camera struct {
	origin          vec3.Point3
	horizontal      vec3.Vec3
	vertical        vec3.Vec3
	lowerLeftCorner vec3.Vec3
	u, v, w         vec3.Vec3
	lensRadius      float64
	time0, time1    float64 // shutter open/close times
}

func NewCamera(
	lookFrom, lookAt vec3.Point3,
	vUp vec3.Vec3,
	vFov, aspectRatio, aperture, focusDist float64,
	time0, time1 float64,
) (cam Camera) {
	theta := helpers.DegreesToRadians(vFov)
	h := math.Tan(theta / 2.0)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	cam.w = lookFrom.SubtractVec3(lookAt).UnitVector()
	cam.u = vUp.Cross(cam.w).UnitVector()
	cam.v = cam.w.Cross(cam.u)

	cam.origin = lookFrom
	cam.horizontal = cam.u.MultiplyFloat(viewportWidth).MultiplyFloat(focusDist)
	cam.vertical = cam.v.MultiplyFloat(viewportHeight).MultiplyFloat(focusDist)
	cam.lowerLeftCorner = cam.origin.
		SubtractVec3(cam.horizontal.DivideFloat(2.0)).
		SubtractVec3(cam.vertical.DivideFloat(2.0)).
		SubtractVec3(cam.w.MultiplyFloat(focusDist))
	cam.lensRadius = aperture / 2
	cam.time0 = time0
	cam.time1 = time1
	return
}

func (c Camera) Ray(s, t float64) material.Ray {
	// depth of field
	rd := vec3util.RandomInUnitDisk().MultiplyFloat(c.lensRadius)
	offset := c.u.MultiplyFloat(rd.X).AddVec3(c.v.MultiplyFloat(rd.Y))

	return material.Ray{
		Origin: c.origin.AddVec3(offset),
		Direction: c.lowerLeftCorner.
			AddVec3(c.horizontal.MultiplyFloat(s)).
			AddVec3(c.vertical.MultiplyFloat(t)).
			SubtractVec3(c.origin).
			SubtractVec3(offset),
		Time: helpers.RandFloat64(c.time0, c.time1),
	}
}

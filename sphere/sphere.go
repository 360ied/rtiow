package sphere

import (
	"math"

	"rtiow/material"
	"rtiow/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
	Mat    material.Material
}

func (s Sphere) Hit(r material.Ray, tMin float64, tMax float64) (material.HitRecord, bool) {
	oc := r.Origin.SubtractVec3(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return material.HitRecord{}, false
	}
	sqrtD := math.Sqrt(discriminant)

	// find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtD) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return material.HitRecord{}, false
		}
	}

	rec := material.HitRecord{}
	rec.T = root
	rec.P = r.At(rec.T)
	// rec.Normal = rec.P.SubtractVec3(s.Center).DivideFloat(s.Radius)
	outwardNormal := rec.P.SubtractVec3(s.Center).DivideFloat(s.Radius)
	rec = rec.FaceNormal(r, outwardNormal)
	rec.Mat = s.Mat

	return rec, true
}

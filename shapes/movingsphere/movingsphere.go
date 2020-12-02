package movingsphere

import (
	"math"

	"rtiow/material"
	"rtiow/vec3"
)

type MovingSphere struct {
	Center0, Center1 vec3.Point3
	Time0, Time1     float64
	Radius           float64
	Mat              material.Material
}

func (s MovingSphere) Hit(r material.Ray, tMin float64, tMax float64) (material.HitRecord, bool) {
	oc := r.Origin.SubtractVec3(s.Center(r.Time))
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
	outwardNormal := rec.P.SubtractVec3(s.Center(r.Time)).DivideFloat(s.Radius)
	rec = rec.FaceNormal(r, outwardNormal)
	rec.Mat = s.Mat

	return rec, true
}

func (m MovingSphere) Center(time float64) vec3.Point3 {
	return m.Center0.AddVec3(m.Center1.SubtractVec3(m.Center0).MultiplyFloat((time - m.Time0) / (m.Time1 - m.Time0)))
}

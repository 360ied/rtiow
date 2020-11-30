package sphere

import (
	"math"

	"rtiow/ray"
	"rtiow/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
}

func (s Sphere) Hit(r ray.Ray, tMin float64, tMax float64) (ray.Record, bool) {
	oc := r.Origin.SubtractVec3(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return ray.Record{}, false
	}
	sqrtD := math.Sqrt(discriminant)

	// find the nearest root that lies in the acceptable range
	root := (-halfB - sqrtD) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return ray.Record{}, false
		}
	}

	rec := ray.Record{}
	rec.T = root
	rec.P = r.At(rec.T)
	// rec.Normal = rec.P.SubtractVec3(s.Center).DivideFloat(s.Radius)
	outwardNormal := rec.P.SubtractVec3(s.Center).DivideFloat(s.Radius)
	rec = rec.FaceNormal(r, outwardNormal)

	return rec, true
}

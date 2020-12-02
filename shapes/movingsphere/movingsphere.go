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
	box              material.Box
}

func NewMovingSphere(center0, center1 vec3.Point3, time0, time1, radius float64, mat material.Material) MovingSphere {
	return MovingSphere{
		center0,
		center1,
		time0,
		time1,
		radius,
		mat,
		material.Box{
			center0.SubtractFloat(radius),
			center0.AddFloat(radius),
		}.Surrounding(
			material.Box{
				center1.SubtractFloat(radius),
				center1.AddFloat(radius),
			}),
	}
}

func (s MovingSphere) Box() material.Box {
	return s.box
}

func (s MovingSphere) Hit(r material.Ray, tMin float64, tMax float64) (material.HitRecord, bool) {
	// if !s.box.Hit(r, tMin, tMax) {
	// 	return material.HitRecord{}, false
	// }
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

func (s MovingSphere) Center(time float64) vec3.Point3 {
	return s.Center0.AddVec3(s.Center1.SubtractVec3(s.Center0).MultiplyFloat((time - s.Time0) / (s.Time1 - s.Time0)))
}

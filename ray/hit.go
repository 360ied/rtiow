package ray

import (
	"rtiow/vec3"
)

// HitRecord
type Record struct {
	P         vec3.Point3
	Normal    vec3.Vec3
	T         float64
	FrontFace bool
}

func (r Record) FaceNormal(ray Ray, outwardNormal vec3.Vec3) Record {
	r.FrontFace = ray.Direction.Dot(outwardNormal) < 0
	if r.FrontFace {
		r.Normal = outwardNormal
	} else {
		r.Normal = outwardNormal.Negate()
	}
	return r
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (Record, bool)
}

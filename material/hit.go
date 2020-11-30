package material

import (
	"rtiow/vec3"
)

type HitRecord struct {
	P         vec3.Point3
	Normal    vec3.Vec3
	T         float64
	FrontFace bool
	Mat       Material
}

func (r HitRecord) FaceNormal(ray Ray, outwardNormal vec3.Vec3) HitRecord {
	r.FrontFace = ray.Direction.Dot(outwardNormal) < 0
	if r.FrontFace {
		r.Normal = outwardNormal
	} else {
		r.Normal = outwardNormal.Negate()
	}
	return r
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (HitRecord, bool)
}

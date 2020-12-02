package dielectric

import (
	"math"
	"math/rand"

	"rtiow/material"
	"rtiow/vec3"
)

type Dielectric struct {
	IR float64 // Index of Refraction
}

func (d Dielectric) Scatter(rIn material.Ray, rec material.HitRecord) (vec3.Colour, material.Ray, bool) {
	attenuation := vec3.Colour{X: 1.0, Y: 1.0, Z: 1.0}
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = 1.0 / d.IR
	} else {
		refractionRatio = d.IR
	}
	unitDirection := rIn.Direction.UnitVector()
	cosTheta := math.Min(unitDirection.Negate().Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	var direction vec3.Vec3
	if /* Cannot Refract */ refractionRatio*sinTheta > 1.0 ||
		d.reflectance(cosTheta, refractionRatio) > rand.Float64() {
		direction = unitDirection.Reflect(rec.Normal)
	} else {
		direction = unitDirection.Refract(rec.Normal, refractionRatio)
	}

	scattered := material.Ray{Origin: rec.P, Direction: direction, Time: rIn.Time}
	return attenuation, scattered, true
}

// polynomial approximation by Christophe Schlick
func (d Dielectric) reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

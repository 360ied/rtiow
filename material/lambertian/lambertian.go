package lambertian

import (
	"rtiow/material"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

type Lambertian struct {
	Albedo vec3.Colour
}

func (l Lambertian) Scatter(_ material.Ray, rec material.HitRecord) (vec3.Colour, material.Ray, bool) {
	scatterDirection := rec.Normal.AddVec3(vec3util.RandomUnitVector())
	// Catch degenerate scatter direction
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}
	scattered := material.Ray{Origin: rec.P, Direction: scatterDirection}
	attenuation := l.Albedo
	return attenuation, scattered, true
}

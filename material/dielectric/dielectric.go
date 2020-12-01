package dielectric

import (
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
	refracted := unitDirection.Refract(rec.Normal, refractionRatio)

	scattered := material.Ray{Origin: rec.P, Direction: refracted}
	return attenuation, scattered, true
}

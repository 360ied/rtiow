package metal

import (
	"rtiow/material"
	"rtiow/vec3"
)

type Metal struct {
	Albedo vec3.Colour
}

func (m Metal) Scatter(rIn material.Ray, rec material.HitRecord) (vec3.Colour, material.Ray, bool) {
	reflected := rIn.Direction.UnitVector().Reflect(rec.Normal)
	scattered := material.Ray{Origin: rec.P, Direction: reflected}
	attenuation := m.Albedo
	return attenuation, scattered, scattered.Direction.Dot(rec.Normal) > 0
}

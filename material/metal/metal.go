package metal

import (
	"rtiow/material"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

type Metal struct {
	Albedo vec3.Colour
	Fuzz   float64
}

func (m Metal) Scatter(rIn material.Ray, rec material.HitRecord) (vec3.Colour, material.Ray, bool) {
	reflected := rIn.Direction.UnitVector().Reflect(rec.Normal)
	scattered := material.Ray{
		Origin:    rec.P,
		Direction: reflected.AddVec3(vec3util.RandomInUnitSphere().MultiplyFloat(m.Fuzz)),
	}
	attenuation := m.Albedo
	return attenuation, scattered, scattered.Direction.Dot(rec.Normal) > 0
}

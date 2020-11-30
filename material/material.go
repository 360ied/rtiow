package material

import (
	"rtiow/vec3"
)

type Material interface {
	Scatter(rIn Ray, rec HitRecord) (vec3.Colour, Ray, bool)
}

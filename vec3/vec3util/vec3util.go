package vec3util

import (
	"fmt"
	"math/rand"

	"rtiow/helpers"
	"rtiow/vec3"
)

func Print(v vec3.Vec3) {
	fmt.Printf("%v %v %v", v.X, v.Y, v.Z)
}

func Random() vec3.Vec3 {
	return vec3.Vec3{X: rand.Float64(), Y: rand.Float64(), Z: rand.Float64()}
}

func RandomMinMax(min, max float64) vec3.Vec3 {
	return vec3.Vec3{
		X: helpers.RandFloat64(min, max),
		Y: helpers.RandFloat64(min, max),
		Z: helpers.RandFloat64(min, max),
	}
}

func RandomInUnitSphere() vec3.Vec3 {
	for {
		p := RandomMinMax(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

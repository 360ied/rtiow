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

func RandomInHemisphere(normal vec3.Vec3) vec3.Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	} else {
		return inUnitSphere.Negate()
	}
}

func RandomUnitVector() vec3.Vec3 {
	return RandomInUnitSphere().UnitVector()
}

func RandomInUnitDisk() vec3.Vec3 {
	for {
		p := vec3.Vec3{X: helpers.RandFloat64(-1, 1), Y: helpers.RandFloat64(-1, 1)}
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

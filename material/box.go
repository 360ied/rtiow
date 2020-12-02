package material

import (
	"math"

	"rtiow/vec3"
)

type Box struct {
	Min,
	Max vec3.Point3
}

func (b Box) Hit(r Ray, tMin, tMax float64) bool {
	invDX := 1.0 / r.Direction.X
	t0X := (b.Min.X - r.Origin.X) * invDX
	t1X := (b.Max.X - r.Origin.X) * invDX
	if invDX < 0.0 {
		t0X, t1X = t1X, t0X
	}
	// Note: Consider inlining the Max and Min functions
	tMin = math.Max(t0X, tMin)
	tMax = math.Min(t1X, tMax)
	if tMax <= tMin {
		return false
	}
	invDY := 1.0 / r.Direction.Y
	t0Y := (b.Min.Y - r.Origin.Y) * invDY
	t1Y := (b.Max.Y - r.Origin.Y) * invDY
	if invDY < 0.0 {
		t0Y, t1Y = t1Y, t0Y
	}
	tMin = math.Max(t0Y, tMin)
	tMax = math.Min(t1Y, tMax)
	if tMax <= tMin {
		return false
	}
	invDZ := 1.0 / r.Direction.Z
	t0Z := (b.Min.Z - r.Origin.Z) * invDZ
	t1Z := (b.Max.Z - r.Origin.Z) * invDZ
	if invDZ < 0.0 {
		t0Z, t1Z = t1Z, t0Z
	}
	tMin = math.Max(t0Z, tMin)
	tMax = math.Min(t1Z, tMax)
	if tMax <= tMin {
		return false
	}
	return true
}

func (b Box) Surrounding(c Box) Box {
	return Box{
		vec3.Point3{
			math.Min(b.Min.X, c.Min.X),
			math.Min(b.Min.Y, c.Min.Y),
			math.Min(b.Min.Z, c.Min.Z),
		},
		vec3.Point3{
			math.Max(b.Max.X, c.Max.X),
			math.Max(b.Max.Y, c.Max.Y),
			math.Max(b.Max.Z, c.Max.Z),
		},
	}
}

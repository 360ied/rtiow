package vec3

import "math"

type (
	Point3 = Vec3
	Colour = Vec3
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Negate() Vec3 {
	v.X, v.Y, v.Z = -v.X, -v.Y, -v.Z
	return v
}

func (v Vec3) AddVec3(c Vec3) Vec3 {
	v.X += c.X
	v.Y += c.Y
	v.Z += c.Z
	return v
}

func (v Vec3) SubtractVec3(c Vec3) Vec3 {
	v.X -= c.X
	v.Y -= c.Y
	v.Z -= c.Z
	return v
}

func (v Vec3) MultiplyVec3(c Vec3) Vec3 {
	v.X *= c.X
	v.Y *= c.Y
	v.Z *= c.Z
	return v
}

func (v Vec3) DivideVec3(c Vec3) Vec3 {
	v.X /= c.X
	v.Y /= c.Y
	v.Z /= c.Z
	return v
}

func (v Vec3) AddFloat(t float64) Vec3 {
	v.X += t
	v.Y += t
	v.Z += t
	return v
}

func (v Vec3) SubtractFloat(t float64) Vec3 {
	v.X -= t
	v.Y -= t
	v.Z -= t
	return v
}

func (v Vec3) MultiplyFloat(t float64) Vec3 {
	v.X *= t
	v.Y *= t
	v.Z *= t
	return v
}

func (v Vec3) DivideFloat(t float64) Vec3 {
	v.X /= t
	v.Y /= t
	v.Z /= t
	return v
}

// Dot product
func (v Vec3) Dot(c Vec3) float64 {
	return v.X*c.X + v.Y*c.Y + v.Z*c.Z
}

func (v Vec3) Cross(c Vec3) (e Vec3) {
	e.X = v.Y*c.Z - v.Z*c.Y // 1 2 2 1
	e.Y = v.Z*c.X - v.X*c.Z // 2 0 0 2
	e.Z = v.X*c.Y - v.Y*c.X // 0 1 1 0
	return
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivideFloat(v.Length())
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

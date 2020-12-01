package helpers

import (
	"math"
	"math/rand"
)

// Makes sure x is within min and max
func Clamp(x, min, max float64) float64 {
	return math.Min(max, math.Max(min, x))
}

func RandFloat64(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

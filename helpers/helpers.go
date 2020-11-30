package helpers

import "math"

// Makes sure x is within min and max
func Clamp(x, min, max float64) float64 {
	return math.Min(max, math.Max(min, x))
}

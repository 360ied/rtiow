package colour

import (
	"fmt"
	"math"

	"rtiow/constants"
	"rtiow/helpers"
	"rtiow/vec3"
)

func WriteColour(v vec3.Colour, samplesPerPixel int) {
	r := v.X
	g := v.Y
	b := v.Z

	// Divide colour by the number of samples
	r = divSampleGammaCorrect(r, samplesPerPixel)
	g = divSampleGammaCorrect(g, samplesPerPixel)
	b = divSampleGammaCorrect(b, samplesPerPixel)

	// Write translated value of each colour component
	fmt.Printf("%v %v %v\n", normalize(r), normalize(g), normalize(b))
}

func normalize(n float64) int {
	return int(constants.Almost256 * helpers.Clamp(n, 0.0, constants.Almost1))
}

func normalize65535(n float64) uint32 {
	const almost65535 = 65534.999
	return uint32(almost65535 * helpers.Clamp(n, 0.0, almost65535))
}

func divSampleGammaCorrect(t float64, samplesPerPixel int) float64 {
	return math.Sqrt(t / float64(samplesPerPixel))
}

type VecColour struct {
	vec3.Colour
	SamplesPerPixel int
}

func (c VecColour) RGBA() (uint32, uint32, uint32, uint32) {
	return ToColour(c)
}

func ToColour(v VecColour) (uint32, uint32, uint32, uint32) {
	r := v.X
	g := v.Y
	b := v.Z

	// Divide colour by the number of samples
	r = divSampleGammaCorrect(r, v.SamplesPerPixel)
	g = divSampleGammaCorrect(g, v.SamplesPerPixel)
	b = divSampleGammaCorrect(b, v.SamplesPerPixel)

	return normalize65535(r), normalize65535(g), normalize65535(b), 65535
}

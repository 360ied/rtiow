package colour

import (
	"fmt"

	"rtiow/constants"
	"rtiow/helpers"
	"rtiow/vec3"
)

func WriteColour(v vec3.Colour, samplesPerPixel int) {
	r := v.X
	g := v.Y
	b := v.Z

	// Divide colour by the number of samples
	r /= float64(samplesPerPixel)
	g /= float64(samplesPerPixel)
	b /= float64(samplesPerPixel)

	// Write translated value of each colour component
	fmt.Printf("%v %v %v\n", normalize(r), normalize(g), normalize(b))
}

func normalize(n float64) int {
	return int(constants.Almost256 * helpers.Clamp(n, 0.0, constants.Almost1))
}

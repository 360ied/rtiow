package colour

import (
	"fmt"
	"math"

	"github.com/oleiade/lane"

	"rtiow/constants"
	"rtiow/helpers"
	"rtiow/vec3"
)

func WriteColour(out *lane.PQueue, v vec3.Colour, samplesPerPixel int, position int) {
	r := v.X
	g := v.Y
	b := v.Z

	// Divide colour by the number of samples
	r = divSampleGammaCorrect(r, samplesPerPixel)
	g = divSampleGammaCorrect(g, samplesPerPixel)
	b = divSampleGammaCorrect(b, samplesPerPixel)

	// Write translated value of each colour component
	out.Push(fmt.Sprintf("%v %v %v\n", normalize(r), normalize(g), normalize(b)), position)
}

func normalize(n float64) int {
	return int(constants.Almost256 * helpers.Clamp(n, 0.0, constants.Almost1))
}

func divSampleGammaCorrect(t float64, samplesPerPixel int) float64 {
	return math.Sqrt(t / float64(samplesPerPixel))
}

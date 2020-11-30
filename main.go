package main

import (
	"fmt"
	"os"

	"rtiow/colour"
	"rtiow/vec3"
)

func main() {
	const (
		imageWidth  = 256
		imageHeight = 256
	)

	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		_, _ = fmt.Fprintf(os.Stderr, "Scanlines remaining: %v\n", j)
		for i := 0; i < imageWidth; i++ {
			pixel := vec3.Vec3{X: float64(i) / (imageWidth - 1), Y: float64(j) / (imageHeight - 1), Z: .25}
			colour.WriteColour(pixel)
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

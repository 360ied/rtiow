package main

import (
	"fmt"
	"os"

	"rtiow/colour"
	"rtiow/ray"
	"rtiow/vec3"
)

func main() {
	// Image
	const (
		aspectRatio = 16.0 / 9.0
		imageWidth  = 400
		imageHeight = imageWidth / aspectRatio
	)

	// Camera
	const (
		viewportHeight = 2.0
		viewportWidth  = aspectRatio * viewportHeight
		focalLength    = 1.0
	)
	origin := vec3.Point3{}                   // 0, 0, 0
	horizontal := vec3.Vec3{X: viewportWidth} // viewportWidth, 0, 0
	vertical := vec3.Vec3{Y: viewportHeight}  // 0, viewportHeight, 0
	lowerLeftCorner := origin.SubtractVec3(
		horizontal.DivideFloat(2.0),
	).SubtractVec3(
		vertical.DivideFloat(2.0),
	).SubtractVec3(vec3.Vec3{Z: focalLength}) // 0, 0, focalLength

	// Render

	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)

	for j := int(imageHeight - 1); j >= 0; j-- {
		_, _ = fmt.Fprintf(os.Stderr, "Scanlines remaining: %v\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / (imageWidth - 1)
			v := float64(j) / (imageHeight - 1)
			r := ray.Ray{
				Origin: origin,
				Direction: lowerLeftCorner.AddVec3(
					horizontal.MultiplyFloat(u),
				).AddVec3(
					vertical.MultiplyFloat(v),
				).SubtractVec3(origin),
			}
			pixelColour := r.Colour()
			colour.WriteColour(pixelColour)
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

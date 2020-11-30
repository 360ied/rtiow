package main

import (
	"fmt"
	"math/rand"
	"os"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/ray"
	"rtiow/sphere"
	"rtiow/vec3"
)

func main() {
	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 400
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 100
	)

	// World
	world := ray.HittableList{
		Objects: []ray.Hittable{
			sphere.Sphere{
				Center: vec3.Point3{Z: -1},
				Radius: .5,
			},
			sphere.Sphere{
				Center: vec3.Point3{Y: -100.5, Z: -1},
				Radius: 100,
			},
		},
	}

	// Camera
	const (
		viewportHeight = 2.0
		viewportWidth  = aspectRatio * viewportHeight
		focalLength    = 1.0
	)
	origin := vec3.Point3{}                   // 0, 0, 0
	horizontal := vec3.Vec3{X: viewportWidth} // viewportWidth, 0, 0
	vertical := vec3.Vec3{Y: viewportHeight}  // 0, viewportHeight, 0
	// lowerLeftCorner := origin.SubtractVec3(
	// 	horizontal.DivideFloat(2.0),
	// ).SubtractVec3(
	// 	vertical.DivideFloat(2.0),
	// ).SubtractVec3(vec3.Vec3{Z: focalLength}) // 0, 0, focalLength

	cam := camera.NewCamera(aspectRatio, viewportHeight, focalLength, origin, horizontal, vertical)

	// Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := int(imageHeight - 1); j >= 0; j-- {
		_, _ = fmt.Fprintf(os.Stderr, "Scanlines remaining: %v\n", j)
		for i := 0; i < imageWidth; i++ {
			pixelColour := vec3.Colour{}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / (imageWidth - 1)
				v := (float64(j) + rand.Float64()) / (imageHeight - 1)
				r := cam.Ray(u, v)
				pixelColour = pixelColour.AddVec3(r.Colour(world))
			}
			colour.WriteColour(pixelColour, samplesPerPixel)
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

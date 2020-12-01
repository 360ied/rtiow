package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sync"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/material"
	"rtiow/material/lambertian"
	"rtiow/sphere"
	"rtiow/vec3"
)

func main() {
	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 400
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 25
		maxDepth        = 10
	)

	// World
	r := math.Cos(math.Pi / 4)

	world := material.HittableList{
		Objects: []material.Hittable{
			sphere.Sphere{
				Center: vec3.Point3{X: -r, Z: -1},
				Radius: r,
				Mat:    lambertian.Lambertian{Albedo: vec3.Colour{Z: 1}},
			},
			sphere.Sphere{
				Center: vec3.Point3{X: r, Z: -1},
				Radius: r,
				Mat:    lambertian.Lambertian{Albedo: vec3.Colour{X: 1}},
			},
		},
	}

	// Camera
	cam := camera.NewCamera(90.0, aspectRatio)

	wg := new(sync.WaitGroup)

	img := make([]vec3.Colour, imageWidth*imageHeight)

	// Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := 0; j < imageHeight; j++ {
		j := j
		wg.Add(1)
		go func() {
			for i := 0; i < imageWidth; i++ {
				pixelColour := vec3.Colour{}
				for s := 0; s < samplesPerPixel; s++ {
					u := (float64(i) + rand.Float64()) / (imageWidth - 1)
					v := (float64(j) + rand.Float64()) / (imageHeight - 1)
					r := cam.Ray(u, v)
					pixelColour = pixelColour.AddVec3(r.Colour(world, maxDepth))
				}
				// colour.WriteColour(pqueue, pixelColour, samplesPerPixel, j*imageWidth-i)
				img[j*imageWidth+i] = pixelColour
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// The image comes out upside down, so go through it from the bottom to top to reverse it
	for x := int(imageHeight) - 1; x >= 0; x-- {
		for y := 0; y < imageWidth; y++ {
			colour.WriteColour(img[x*imageWidth+y], samplesPerPixel)
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

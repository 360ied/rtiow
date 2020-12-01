package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/material"
	"rtiow/material/dielectric"
	"rtiow/material/lambertian"
	"rtiow/material/metal"
	"rtiow/sphere"
	"rtiow/vec3"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func main() {
	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 400
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 20
		maxDepth        = 10
	)

	// World
	matGround := lambertian.Lambertian{vec3.Colour{0.8, 0.8, 0.0}}
	matCenter := lambertian.Lambertian{vec3.Colour{0.1, 0.2, 0.5}}
	matLeft := dielectric.Dielectric{1.5}
	matRight := metal.Metal{vec3.Colour{0.8, 0.6, 0.2}, 0.0}

	world := material.HittableList{}
	// ground
	world.Add(sphere.Sphere{vec3.Point3{0.0, -100.5, -1.0}, 100.0, matGround})
	// blue ball in the center
	world.Add(sphere.Sphere{vec3.Point3{0.0, 0.0, -1.0}, 0.5, matCenter})
	// glass ball on the left
	world.Add(sphere.Sphere{vec3.Point3{-1.0, 0.0, -1.0}, 0.5, matLeft})
	// negative radius glass ball within the glass ball on the left
	world.Add(sphere.Sphere{vec3.Point3{-1.0, 0.0, -1.0}, -0.45, matLeft})
	// yellow metal ball on the right
	world.Add(sphere.Sphere{vec3.Point3{1.0, 0.0, -1.0}, 0.5, matRight})

	// Camera
	lookFrom := vec3.Point3{3, 3, 2}
	lookAt := vec3.Point3{0, 0, -1}
	vUp := vec3.Vec3{0, 1, 0}
	distToFocus := (lookFrom.SubtractVec3(lookAt)).Length()
	aperture := 2.0

	cam := camera.NewCamera(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

	// Render
	wg := new(sync.WaitGroup)
	img := make([]vec3.Colour, imageWidth*imageHeight)

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

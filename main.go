package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/helpers"
	"rtiow/material"
	"rtiow/material/dielectric"
	"rtiow/material/lambertian"
	"rtiow/material/metal"
	"rtiow/sphere"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func main() {
	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 3840
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 1000
		maxDepth        = 100
	)

	// World
	_, _ = fmt.Fprintln(os.Stderr, "Generating random scene...")
	world := randomScene()
	_, _ = fmt.Fprintln(os.Stderr, "Done generating random scene.")

	// Camera
	lookFrom := vec3.Point3{13, 2, 3}
	lookAt := vec3.Point3{0, 0, 0}
	vUp := vec3.Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.1

	cam := camera.NewCamera(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

	// Render
	wg := new(sync.WaitGroup)
	img := make([]vec3.Colour, imageWidth*imageHeight)

	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	_, _ = fmt.Fprintln(os.Stderr, "Starting render...")
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
	_, _ = fmt.Fprintln(os.Stderr, "Waiting for rendering to finish...")
	wg.Wait()

	_, _ = fmt.Fprintln(os.Stderr, "Rendering is done.\nWriting image...")
	// The image comes out upside down, so go through it from the bottom to top to reverse it
	for x := int(imageHeight) - 1; x >= 0; x-- {
		for y := 0; y < imageWidth; y++ {
			colour.WriteColour(img[x*imageWidth+y], samplesPerPixel)
		}
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

func randomScene() material.HittableList {
	world := material.HittableList{}

	groundMaterial := lambertian.Lambertian{vec3.Colour{0.5, 0.5, 0.5}}
	world.Add(sphere.Sphere{vec3.Point3{0, -1000, 0}, 1000, groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vec3.Point3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.SubtractVec3(vec3.Point3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := vec3util.Random().MultiplyVec3(vec3util.Random())
					sphereMaterial := lambertian.Lambertian{albedo}
					world.Add(sphere.Sphere{center, 0.2, sphereMaterial})
				} else if chooseMat < 0.95 {
					// metal
					albedo := vec3util.RandomMinMax(0.5, 1)
					fuzz := helpers.RandFloat64(0, 0.5)
					sphereMaterial := metal.Metal{albedo, fuzz}
					world.Add(sphere.Sphere{center, 0.2, sphereMaterial})
				} else {
					// glass
					sphereMaterial := dielectric.Dielectric{1.5}
					world.Add(sphere.Sphere{center, 0.2, sphereMaterial})
				}
			}
		}
	}

	material1 := dielectric.Dielectric{1.5}
	world.Add(sphere.Sphere{vec3.Point3{0, 1, 0}, 1.0, material1})

	material2 := lambertian.Lambertian{vec3.Colour{0.4, 0.2, 0.1}}
	world.Add(sphere.Sphere{vec3.Point3{-4, 1, 0}, 1.0, material2})

	material3 := metal.Metal{vec3.Colour{0.7, 0.6, 0.5}, 0.0}
	world.Add(sphere.Sphere{vec3.Point3{4, 1, 0}, 1.0, material3})

	return world
}

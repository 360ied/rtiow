package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/oleiade/lane"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/material"
	"rtiow/material/lambertian"
	"rtiow/material/metal"
	"rtiow/sphere"
	"rtiow/vec3"
)

func main() {
	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 1920
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 100
		maxDepth        = 50
	)

	materialGround := lambertian.Lambertian{Albedo: vec3.Colour{X: 0.8, Y: 0.8}}
	materialCenter := lambertian.Lambertian{Albedo: vec3.Colour{X: 0.7, Y: 0.3, Z: 0.3}}
	materialLeft := metal.Metal{Albedo: vec3.Colour{X: 0.8, Y: 0.8, Z: 0.8}}
	materalRight := metal.Metal{Albedo: vec3.Colour{X: 0.8, Y: 0.6, Z: 0.2}}

	// World
	world := material.HittableList{
		Objects: []material.Hittable{
			sphere.Sphere{
				Center: vec3.Point3{Y: -100.5, Z: -1.0}, // 0.0, -100.5, -1.0
				Radius: 100.0,
				Mat:    materialGround,
			},
			sphere.Sphere{
				Center: vec3.Point3{Z: -1.0}, // 0.0, 0.0, -1.0
				Radius: 0.5,
				Mat:    materialCenter,
			},
			sphere.Sphere{
				Center: vec3.Point3{X: -1.0, Z: -1.0}, // -1.0, 0.0, -1.0
				Radius: 0.5,
				Mat:    materialLeft,
			},
			sphere.Sphere{
				Center: vec3.Point3{X: 1.0, Z: -1.0}, // 1.0, 0.0, -1.0
				Radius: 0.5,
				Mat:    materalRight,
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

	wg := new(sync.WaitGroup)
	pqueue := lane.NewPQueue(lane.MAXPQ)

	// Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := int(imageHeight - 1); j >= 0; j-- {
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
				colour.WriteColour(pqueue, pixelColour, samplesPerPixel, j*imageWidth-i)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	for {
		pix, _ := pqueue.Pop()
		if pix == nil {
			break
		}
		fmt.Print(pix)
	}

	_, _ = fmt.Fprintf(os.Stderr, "Done.\n")
}

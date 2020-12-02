package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"rtiow/camera"
	"rtiow/colour"
	"rtiow/helpers"
	"rtiow/material"
	"rtiow/material/dielectric"
	"rtiow/material/lambertian"
	"rtiow/material/metal"
	"rtiow/shapes/movingsphere"
	"rtiow/shapes/sphere"
	"rtiow/vec3"
	"rtiow/vec3/vec3util"
)

//goland:noinspection GoStructInitializationWithoutFieldNames
func main() {
	rand.Seed(time.Now().UnixNano())

	// Image
	const (
		aspectRatio     = 16.0 / 9.0
		imageWidth      = 400
		imageHeight     = imageWidth / aspectRatio
		samplesPerPixel = 100
		maxDepth        = 50
	)

	// Create window
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, windowErr := sdl.CreateWindow("RTIOW", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, imageWidth, imageHeight, sdl.WINDOW_SHOWN)
	if windowErr != nil {
		panic(windowErr)
	}
	defer func() {
		windowDestroyErr := window.Destroy()
		if windowDestroyErr != nil {
			panic(windowDestroyErr)
		}
	}()

	surface, surfaceErr := window.GetSurface()
	if surfaceErr != nil {
		panic(surfaceErr)
	}

	surfaceFillRectErr := surface.FillRect(nil, 0)
	if surfaceFillRectErr != nil {
		panic(surfaceFillRectErr)
	}

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

	cam := camera.NewCamera(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus, 0.0, 1.0)

	// Render
	wg := new(sync.WaitGroup)
	img := make([]vec3.Colour, imageWidth*imageHeight)

	counter := uint64(0)
	startTime := time.Now()

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
				img[j*imageWidth+i] = pixelColour
				atomic.AddUint64(&counter, 1)
				surface.Set(i, imageHeight-j-1, colour.VecColour{pixelColour, samplesPerPixel})
				// _ = window.UpdateSurface()
			}
			wg.Done()
		}()
	}
	_, _ = fmt.Fprintln(os.Stderr, "Waiting for rendering to finish...")
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pixelsDone := atomic.LoadUint64(&counter)
			elapsedTime := time.Now().Sub(startTime)
			_ = window.UpdateSurface()
			_, _ = fmt.Fprintf(
				os.Stderr,
				"%v/%v pixels rendered. %v%v done. %v left.\n",
				pixelsDone, len(img), float64(pixelsDone)/float64(len(img))*100, "%",
				time.Duration(float64(uint64(len(img))-pixelsDone)/(float64(pixelsDone)/elapsedTime.Seconds()))*time.Second)
		}
	}()
	wg.Wait()

	_, _ = fmt.Fprintf(os.Stderr, "Took %v\n", time.Now().Sub(startTime))

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
	world.Add(sphere.NewSphere(vec3.Point3{0, -1000, 0}, 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vec3.Point3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.SubtractVec3(vec3.Point3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := vec3util.Random().MultiplyVec3(vec3util.Random())
					sphereMaterial := lambertian.Lambertian{albedo}
					center2 := center.AddVec3(vec3.Vec3{0, helpers.RandFloat64(0, 0.5), 0})
					// world.Add(sphere.Sphere{center, 0.2, sphereMaterial})
					world.Add(movingsphere.NewMovingSphere(center, center2, 0.0, 1.0, 0.2, sphereMaterial))
				} else if chooseMat < 0.95 {
					// metal
					albedo := vec3util.RandomMinMax(0.5, 1)
					fuzz := helpers.RandFloat64(0, 0.5)
					sphereMaterial := metal.Metal{albedo, fuzz}
					world.Add(sphere.NewSphere(center, 0.2, sphereMaterial))
				} else {
					// glass
					sphereMaterial := dielectric.Dielectric{1.5}
					world.Add(sphere.NewSphere(center, 0.2, sphereMaterial))
				}
			}
		}
	}

	material1 := dielectric.Dielectric{1.5}
	world.Add(sphere.NewSphere(vec3.Point3{0, 1, 0}, 1.0, material1))

	material2 := lambertian.Lambertian{vec3.Colour{0.4, 0.2, 0.1}}
	world.Add(sphere.NewSphere(vec3.Point3{-4, 1, 0}, 1.0, material2))

	material3 := metal.Metal{vec3.Colour{0.7, 0.6, 0.5}, 0.0}
	world.Add(sphere.NewSphere(vec3.Point3{4, 1, 0}, 1.0, material3))

	return world
}

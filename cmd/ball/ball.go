package main

import (
	rt "github.com/gumuz/go-raytracer/raytracer"
)

func main() {
	rayOrigin := rt.Point(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := rt.NewCanvas(canvasPixels, canvasPixels)
	white := &rt.Color{1, 1, 1}
	sphere := rt.NewSphere()
	// sphere.SetTransform(rt.Scaling(0.5, 1, 1).Mul(rt.RotationZ(math.Pi / 4)).Mul(rt.Shearing(1, 0, 0, 0, 0, 0)))

	// worldX, worldY := 0.0, 0.0
	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := rt.Point(worldX, worldY, wallZ)
			ray := rt.NewRay(rayOrigin, position.Sub(rayOrigin).Norm())
			xs := sphere.Intersect(ray)

			if xs.Hit() != nil {
				canvas.SetAt(x, y, white)
			}
		}
	}

	canvas.Save("ball.png")

}

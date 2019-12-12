package main

import (
	rt "github.com/gumuz/go-raytracer/raytracer"
)

func main() {
	rayOrigin := rt.NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := rt.NewCanvas(canvasPixels, canvasPixels)
	sphere := rt.NewSphere()
	sphere.Material.Color = &rt.Color{1, 0.7, 1}
	light := rt.NewPointLight(rt.NewPoint(-10, 10, -10), &rt.Color{1, 1, 1})

	// sphere.SetTransform(rt.Scaling(0.5, 1, 1).Mul(rt.RotationZ(math.Pi / 4)).Mul(rt.Shearing(1, 0, 0, 0, 0, 0)))

	// worldX, worldY := 0.0, 0.0
	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := rt.NewPoint(worldX, worldY, wallZ)
			ray := rt.NewRay(rayOrigin, position.Sub(rayOrigin).Norm())
			xs := sphere.Intersect(ray)
			hit := xs.Hit()

			if hit != nil {
				point := ray.Pos(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := ray.Direction.Neg()

				color := hit.Object.GetMaterial().Lighting(light, point, eye, normal)

				canvas.SetAt(x, y, color)
			}
		}
	}

	canvas.Save("ball.png")

}

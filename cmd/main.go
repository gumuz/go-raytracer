package main

import (
	"math"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func main() {
	canvas := rt.NewCanvas(500, 500)
	white := rt.NewColor(1, 1, 1)

	middle := rt.NewPoint(0, 0, 0)
	positions := []*rt.Point{middle}

	point := rt.NewPoint(0, 1, 0)
	for i := 0; i < 12; i++ {
		point = rt.RotationZ(math.Pi / 6).MultP(point)
		positions = append(positions, point)
	}

	for _, point := range positions {
		point = rt.Scaling(20, 20, 20).MultP(point)
		point = rt.Translation(250, 250, 0).MultP(point)
		canvas.WritePixel(int(point.X), int(point.Y), white)
	}

	canvas.Save()
}

package main

import (
	"github.com/gumuz/go-raytracer/raytracer"
)

type projectile struct {
	position *raytracer.Point
	velocity *raytracer.Vector
}

type environment struct {
	gravity *raytracer.Vector
	wind    *raytracer.Vector
}

func tick(p *projectile, e *environment) *projectile {
	position := p.position.AddV(p.velocity)
	velocity := p.velocity.Add(e.gravity).Add(e.wind)
	return &projectile{position, velocity}
}

func main() {
	canvas := raytracer.NewCanvas(900, 550)
	white := raytracer.NewColor(1, 1, 1)
	p := &projectile{raytracer.NewPoint(0, 1, 0), raytracer.NewVector(1, 1.8, 0).Norm().Mult(11.25)}
	e := &environment{raytracer.NewVector(0, -0.1, 0), raytracer.NewVector(-0.01, 0, 0)}

	for p.position.Y > 0 {
		p = tick(p, e)
		canvas.WritePixel(int(p.position.X), int(p.position.Y), white)
	}

	canvas.Save()
}

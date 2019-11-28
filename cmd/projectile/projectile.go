package main

import (
	rt "github.com/gumuz/go-raytracer/raytracer"
)

func main() {
	projectile := struct {
		position *rt.Tuple
		velocity *rt.Tuple
	}{
		rt.Point(0, 1, 0),
		rt.Vector(1, 1.8, 0).Norm().Mul(11.25),
	}

	environment := struct {
		gravity *rt.Tuple
		wind    *rt.Tuple
	}{
		rt.Vector(0, -0.1, 0),
		rt.Vector(-0.01, 0, 0),
	}

	width, height := 900, 550

	canvas := rt.NewCanvas(width, height)
	white := &rt.Color{1, 1, 1}

	for projectile.position.Y > 0 {
		canvas.SetAt(int(projectile.position.X), height-int(projectile.position.Y), white)

		projectile.position = projectile.position.Add(projectile.velocity)
		projectile.velocity = projectile.velocity.Add(environment.gravity).Add(environment.wind)
	}

	canvas.Save("projectile.png")

}

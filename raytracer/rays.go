package raytracer

type Ray struct {
	Origin    *Point
	Direction *Vector
}

func NewRay(origin *Point, direction *Vector) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Pos(distance float64) *Point {
	return r.Origin.AddV(r.Direction.Mult(distance))
}

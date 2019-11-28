package raytracer

type Ray struct {
	Origin    *Tuple
	Direction *Tuple
}

func NewRay(origin *Tuple, direction *Tuple) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Pos(distance float64) *Tuple {
	return r.Origin.Add(r.Direction.Mul(distance))
}

func (r *Ray) Transform(transformation Matrix) *Ray {
	return NewRay(
		transformation.MulT(r.Origin),
		transformation.MulT(r.Direction),
	)
}

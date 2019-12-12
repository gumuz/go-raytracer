package raytracer

import "math"

type Material struct {
	Color     *Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() *Material {
	return &Material{&Color{1, 1, 1}, 0.1, 0.9, 0.9, 200.0}
}

func (m *Material) Equals(b *Material) bool {
	return m.Color.Equals(b.Color) && m.Ambient == b.Ambient && m.Diffuse == b.Diffuse && m.Specular == b.Specular && m.Shininess == b.Shininess
}

func (m *Material) Lighting(l *PointLight, p *Tuple, eyev *Tuple, normalv *Tuple) *Color {
	effectiveColor := m.Color.Prod(l.Intensity)

	lightv := l.Position.Sub(p).Norm()

	ambient := effectiveColor.Mul(m.Ambient)

	lightDotNormal := lightv.Dot(normalv)
	diffuse := &Color{0, 0, 0}
	specular := &Color{0, 0, 0}
	if lightDotNormal >= 0 {
		diffuse = effectiveColor.Mul(m.Diffuse).Mul(lightDotNormal)

		reflectv := lightv.Neg().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye <= 0 {
			specular = &Color{0, 0, 0}
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Mul(m.Specular).Mul(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}

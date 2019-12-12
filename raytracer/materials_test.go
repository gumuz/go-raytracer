package raytracer_test

import (
	"math"
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestDefaultMaterial(t *testing.T) {
	/* Scenario: The default material
	   Given m ← material()
	   Then m.color = color(1, 1, 1)
	     And m.ambient = 0.1
	     And m.diffuse = 0.9
	     And m.specular = 0.9
		 And m.shininess = 200.0 */
	m := rt.NewMaterial()

	if !m.Color.Equals(&rt.Color{1, 1, 1}) {
		t.Errorf("Error: %v", m.Color)
	}

	if m.Ambient != 0.1 {
		t.Errorf("Error: %v", m.Ambient)
	}

	if m.Diffuse != 0.9 {
		t.Errorf("Error: %v", m.Diffuse)
	}

	if m.Specular != 0.9 {
		t.Errorf("Error: %v", m.Specular)
	}

	if m.Shininess != 200.0 {
		t.Errorf("Error: %v", m.Shininess)
	}
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	/* Scenario: Lighting with the eye between the light and the surface
	   Given eyev ← vector(0, 0, -1)
	     And normalv ← vector(0, 0, -1)
	     And light ← point_light(point(0, 0, -10), color(1, 1, 1))
	   When result ← lighting(m, light, position, eyev, normalv)
	   Then result = color(1.9, 1.9, 1.9) */
	m := rt.NewMaterial()
	position := rt.NewPoint(0, 0, 0)

	eyev := rt.NewVector(0, 0, -1)
	normalv := rt.NewVector(0, 0, -1)
	light := rt.NewPointLight(rt.NewPoint(0, 0, -10), &rt.Color{1, 1, 1})

	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equals(&rt.Color{1.9, 1.9, 1.9}) {
		t.Errorf("Error: %v", result)
	}
}
func TestLightingEyeBetweenLightAndSurfaceEyeOffset45(t *testing.T) {
	/* Scenario: Lighting with the eye between light and surface, eye offset 45°
	   Given eyev ← vector(0, √2/2, -√2/2)
	     And normalv ← vector(0, 0, -1)
	     And light ← point_light(point(0, 0, -10), color(1, 1, 1))
	   When result ← lighting(m, light, position, eyev, normalv)
	   Then result = color(1.0, 1.0, 1.0) */
	m := rt.NewMaterial()
	position := rt.NewPoint(0, 0, 0)

	eyev := rt.NewVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := rt.NewVector(0, 0, -1)
	light := rt.NewPointLight(rt.NewPoint(0, 0, -10), &rt.Color{1, 1, 1})

	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equals(&rt.Color{1.0, 1.0, 1.0}) {
		t.Errorf("Error: %v", result)
	}
}
func TestLightingEyeOppositeSurfaceLightOffset45(t *testing.T) {
	/* Scenario: Lighting with eye opposite surface, light offset 45°
	   Given eyev ← vector(0, 0, -1)
	     And normalv ← vector(0, 0, -1)
	     And light ← point_light(point(0, 10, -10), color(1, 1, 1))
	   When result ← lighting(m, light, position, eyev, normalv)
	   Then result = color(0.7364, 0.7364, 0.7364) */
	m := rt.NewMaterial()
	position := rt.NewPoint(0, 0, 0)

	eyev := rt.NewVector(0, 0, -1)
	normalv := rt.NewVector(0, 0, -1)
	light := rt.NewPointLight(rt.NewPoint(0, 10, -10), &rt.Color{1, 1, 1})

	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equals(&rt.Color{0.7364, 0.7364, 0.7364}) {
		t.Errorf("Error: %v", result)
	}
}
func TestLightingEyeInPathOfReflectionVector(t *testing.T) {
	/* Scenario: Lighting with eye in the path of the reflection vector
	   Given eyev ← vector(0, -√2/2, -√2/2)
	     And normalv ← vector(0, 0, -1)
	     And light ← point_light(point(0, 10, -10), color(1, 1, 1))
	   When result ← lighting(m, light, position, eyev, normalv)
	   Then result = color(1.6364, 1.6364, 1.6364) */
	m := rt.NewMaterial()
	position := rt.NewPoint(0, 0, 0)

	eyev := rt.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := rt.NewVector(0, 0, -1)
	light := rt.NewPointLight(rt.NewPoint(0, 10, -10), &rt.Color{1, 1, 1})

	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equals(&rt.Color{1.6364, 1.6364, 1.6364}) {
		t.Errorf("Error: %v", result)
	}
}
func TestLightingLightBehindSurface(t *testing.T) {
	/* Scenario: Lighting with the light behind the surface
	   Given eyev ← vector(0, 0, -1)
	     And normalv ← vector(0, 0, -1)
	     And light ← point_light(point(0, 0, 10), color(1, 1, 1))
	   When result ← lighting(m, light, position, eyev, normalv)
	   Then result = color(0.1, 0.1, 0.1) */
	m := rt.NewMaterial()
	position := rt.NewPoint(0, 0, 0)

	eyev := rt.NewVector(0, 0, -1)
	normalv := rt.NewVector(0, 0, -1)
	light := rt.NewPointLight(rt.NewPoint(0, 0, 10), &rt.Color{1, 1, 1})

	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equals(&rt.Color{0.1, 0.1, 0.1}) {
		t.Errorf("Error: %v", result)
	}
}

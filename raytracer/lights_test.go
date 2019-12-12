package raytracer_test

import (
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestPointLight(t *testing.T) {
	/* Scenario: A point light has a position and intensity
	   Given intensity ← color(1, 1, 1)
	     And position ← point(0, 0, 0)
	   When light ← point_light(position, intensity)
	   Then light.position = position
		 And light.intensity = intensity */
	intensity := &rt.Color{1, 1, 1}
	position := rt.NewPoint(0, 0, 0)

	light := rt.NewPointLight(position, intensity)

	if light.Intensity != intensity {
		t.Errorf("Error: %v", light.Intensity)
	}

	if light.Position != position {
		t.Errorf("Error: %v", light.Position)
	}
}

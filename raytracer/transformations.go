package raytracer

import (
	"math"
)

func Translation(x, y, z float64) Matrix {
	matrix := Identity()

	matrix[0][3] = x
	matrix[1][3] = y
	matrix[2][3] = z

	return matrix
}

func Scaling(x, y, z float64) Matrix {
	matrix := Identity()

	matrix[0][0] = x
	matrix[1][1] = y
	matrix[2][2] = z

	return matrix
}

func RotationX(radians float64) Matrix {
	matrix := Identity()

	matrix[1][1] = math.Cos(radians)
	matrix[1][2] = -math.Sin(radians)
	matrix[2][1] = math.Sin(radians)
	matrix[2][2] = math.Cos(radians)

	return matrix
}

func RotationY(radians float64) Matrix {
	matrix := Identity()

	matrix[0][0] = math.Cos(radians)
	matrix[0][2] = math.Sin(radians)
	matrix[2][0] = -math.Sin(radians)
	matrix[2][2] = math.Cos(radians)

	return matrix
}

func RotationZ(radians float64) Matrix {
	matrix := Identity()

	matrix[0][0] = math.Cos(radians)
	matrix[0][1] = -math.Sin(radians)
	matrix[1][0] = math.Sin(radians)
	matrix[1][1] = math.Cos(radians)

	return matrix
}

func Shearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	matrix := Identity()

	matrix[0][1] = xy
	matrix[0][2] = xz
	matrix[1][0] = yx
	matrix[1][2] = yz
	matrix[2][0] = zx
	matrix[2][1] = zy

	return matrix
}

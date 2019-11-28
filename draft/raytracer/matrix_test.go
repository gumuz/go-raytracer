package raytracer_test

import (
	"testing"

	"github.com/gumuz/go-raytracer/raytracer"
)

func TestNewMatrix2(t *testing.T) {
	matrix := raytracer.NewMatrix2(
		-3, 5,
		1, -2,
	)

	if matrix[0][0] != -3 || matrix[0][1] != 5 || matrix[1][1] != -2 {
		t.Errorf("%v is incorrect", matrix)
	}
}

func TestNewMatrix3(t *testing.T) {
	matrix := raytracer.NewMatrix3(
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	)

	if matrix[0][0] != -3 || matrix[1][1] != -2 || matrix[2][2] != 1 {
		t.Errorf("%v is incorrect", matrix)
	}
}

func TestNewMatrix4(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	)

	if matrix[0][0] != 1 || matrix[0][3] != 4 || matrix[1][0] != 5.5 || matrix[1][2] != 7.5 || matrix[2][2] != 11 || matrix[3][0] != 13.5 || matrix[3][2] != 15.5 {
		t.Errorf("%v is incorrect", matrix)
	}
}

func TestNewTuple4(t *testing.T) {
	tuple := raytracer.NewTuple4(
		1, 2, 3, 4,
	)

	if tuple != [4]float64{1, 2, 3, 4} {
		t.Errorf("%v is incorrect", tuple)
	}
}

func TestMatrix2Equal(t *testing.T) {
	matrix1 := raytracer.NewMatrix2(
		1, 2,
		5, 6,
	)

	matrix2 := raytracer.NewMatrix2(
		2, 3,
		6, 7,
	)

	if !matrix1.Equals(matrix1) {
		t.Errorf("%v != %v", matrix1, matrix1)
	}

	if matrix1.Equals(matrix2) {
		t.Errorf("%v == %v", matrix1, matrix2)
	}
}

func TestMatrix3Equal(t *testing.T) {
	matrix1 := raytracer.NewMatrix3(
		1, 2, 3,
		5, 6, 7,
		9, 8, 7,
	)

	matrix2 := raytracer.NewMatrix3(
		2, 3, 4,
		6, 7, 8,
		8, 7, 6,
	)

	if !matrix1.Equals(matrix1) {
		t.Errorf("%v != %v", matrix1, matrix1)
	}

	if matrix1.Equals(matrix2) {
		t.Errorf("%v == %v", matrix1, matrix2)
	}
}

func TestMatrix3SubMatrix(t *testing.T) {
	matrix := raytracer.NewMatrix3(
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	)

	result := matrix.SubMatrix(0, 2)
	expected := raytracer.NewMatrix2(
		-3, 2,
		0, 6,
	)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

// func TestMatrix3Minor(t *testing.T) {
// 	matrix := raytracer.NewMatrix3(
// 		3, 5, 0,
// 		2, -1, -7,
// 		6, -1, 5,
// 	)

// 	result := matrix.Minor(1, 0)

// 	if result != 25 {
// 		t.Errorf("%v is incorrect", result)
// 	}
// }

func TestMatrix3Cofact(t *testing.T) {
	matrix1 := raytracer.NewMatrix3(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)

	result := matrix1.Cofact(0, 0)

	if result != -12 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix1.Cofact(1, 0)

	if result != -25 {
		t.Errorf("%v is incorrect", result)
	}

	matrix2 := raytracer.NewMatrix3(
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	)

	result = matrix2.Cofact(0, 0)

	if result != 56 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix2.Cofact(0, 1)

	if result != 12 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix2.Cofact(0, 2)

	if result != -46 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix2.Determ()

	if result != -196 {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix4Equal(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	matrix2 := raytracer.NewMatrix4(
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	)

	if !matrix1.Equals(matrix1) {
		t.Errorf("%v != %v", matrix1, matrix1)
	}

	if matrix1.Equals(matrix2) {
		t.Errorf("%v == %v", matrix1, matrix2)
	}
}

func TestMatrix4Mult(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	matrix2 := raytracer.NewMatrix4(
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	)

	result := matrix1.Mult(matrix2)
	expected := raytracer.NewMatrix4(
		20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42,
	)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix4MultP(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	)

	point := raytracer.NewPoint(
		1, 2, 3,
	)

	result := matrix.MultP(point)

	expected := raytracer.NewPoint(
		18, 24, 33,
	)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix4Trans(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	)

	result := matrix.Trans()
	expected := raytracer.NewMatrix4(
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}

	if !raytracer.Identity.Trans().Equals(raytracer.Identity) {
		t.Errorf("%v is not Identity", raytracer.Identity.Trans())
	}

}

func TestIdentityMatrix4(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	result := matrix.Mult(raytracer.Identity)

	if !result.Equals(matrix) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix2Determ(t *testing.T) {
	matrix := raytracer.NewMatrix2(
		1, 5,
		-3, 2,
	)

	result := matrix.Determ()

	if result != 17 {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix4SubMatrix(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	)

	result := matrix.SubMatrix(2, 1)
	expected := raytracer.NewMatrix3(
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	)

	if !result.Equals(expected) {
		t.Errorf("%v is incorrect", result)
	}
}

func TestMatrix4Cofact(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	)

	result := matrix1.Cofact(0, 0)

	if result != 690 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix1.Cofact(0, 1)

	if result != 447 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix1.Cofact(0, 2)

	if result != 210 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix1.Cofact(0, 3)

	if result != 51 {
		t.Errorf("%v is incorrect", result)
	}

	result = matrix1.Determ()

	if result != -4071 {
		t.Errorf("%v is incorrect", result)
	}

}

func TestMatrix4Invertible(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6,
	)

	if matrix1.Determ() == 0 {
		t.Errorf("%v is incorrect", matrix1.Determ())
	}

	matrix2 := raytracer.NewMatrix4(
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	)

	if matrix2.Determ() != 0 {
		t.Errorf("%v is incorrect", matrix2.Determ())
	}

}

func TestMatrix4Inverse(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	)

	if matrix1.Determ() != 532 {
		t.Errorf("%v is incorrect", matrix1.Determ())
	}

	if matrix1.Cofact(2, 3) != -160 {
		t.Errorf("%v is incorrect", matrix1.Cofact(2, 3))
	}

	if matrix1.Cofact(3, 2) != 105 {
		t.Errorf("%v is incorrect", matrix1.Cofact(3, 2))
	}

	inverted := matrix1.Inverse()

	expected := raytracer.NewMatrix4(
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	)

	if !inverted.Equals(expected) {
		t.Errorf("%v is incorrect", inverted)
	}

	matrix2 := raytracer.NewMatrix4(
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	)

	inverted = matrix2.Inverse()

	expected = raytracer.NewMatrix4(
		-0.15385, -0.15385, -0.28205, -0.53846,
		-0.07692, 0.12308, 0.02564, 0.03077,
		0.35897, 0.35897, 0.43590, 0.92308,
		-0.69231, -0.69231, -0.76923, -1.92308,
	)

	if !inverted.Equals(expected) {
		t.Errorf("%v is incorrect", inverted)
	}

	matrix3 := raytracer.NewMatrix4(
		9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2,
	)

	inverted = matrix3.Inverse()

	expected = raytracer.NewMatrix4(
		-0.04074, -0.07778, 0.14444, -0.22222,
		-0.07778, 0.03333, 0.36667, -0.33333,
		-0.02901, -0.14630, -0.10926, 0.12963,
		0.17778, 0.06667, -0.26667, 0.33333,
	)

	if !inverted.Equals(expected) {
		t.Errorf("%v is incorrect", inverted)
	}

}

func TestMatrix4InverseMult(t *testing.T) {
	matrix1 := raytracer.NewMatrix4(
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	)
	matrix2 := raytracer.NewMatrix4(
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	)

	prod := matrix1.Mult(matrix2)

	if !prod.Mult(matrix2.Inverse()).Equals(matrix1) {
		t.Errorf("%v is incorrect", prod.Mult(matrix2.Inverse()))
	}
}

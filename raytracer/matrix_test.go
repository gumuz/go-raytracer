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

func TestMatrix4MultT(t *testing.T) {
	matrix := raytracer.NewMatrix4(
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	)

	tuple := raytracer.NewTuple4(
		1, 2, 3, 1,
	)

	result := matrix.MultT(tuple)
	expected := raytracer.NewTuple4(
		18, 24, 33, 1,
	)

	if result != expected {
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

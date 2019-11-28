package raytracer_test

import (
	"testing"

	rt "github.com/gumuz/go-raytracer/raytracer"
)

func TestMatrix4(t *testing.T) {
	/* Scenario: Constructing and inspecting a 4x4 matrix
	   Given the following 4x4 matrix M:
		 |  1   |  2   |  3   |  4   |
		 |  5.5 |  6.5 |  7.5 |  8.5 |
		 |  9   | 10   | 11   | 12   |
		 | 13.5 | 14.5 | 15.5 | 16.5 |
	   Then M[0,0] = 1
		 And M[0,3] = 4
		 And M[1,0] = 5.5
		 And M[1,2] = 7.5
		 And M[2,2] = 11
		 And M[3,0] = 13.5
		 And M[3,2] = 15.5 */
	M := rt.Matrix4(
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5)

	if M[0][0] != 1 || M[0][3] != 4 || M[1][0] != 5.5 || M[1][2] != 7.5 || M[2][2] != 11 || M[3][0] != 13.5 || M[3][2] != 15.5 {
		t.Errorf("Error: %v", M)
	}

}

func TestMatrix2(t *testing.T) {
	/* Scenario: A 2x2 matrix ought to be representable
	   Given the following 2x2 matrix M:
		 | -3 |  5 |
		 |  1 | -2 |
	   Then M[0,0] = -3
		 And M[0,1] = 5
		 And M[1,0] = 1
		 And M[1,1] = -2 */
	M := rt.Matrix2(
		-3, 5,
		1, -2)

	if M[0][0] != -3 || M[0][1] != 5 || M[1][1] != -2 {
		t.Errorf("Error: %v", M)
	}
}

func TestMatrix3(t *testing.T) {
	/* Scenario: A 3x3 matrix ought to be representable
	   Given the following 3x3 matrix M:
		 | -3 |  5 |  0 |
		 |  1 | -2 | -7 |
		 |  0 |  1 |  1 |
	   Then M[0,0] = -3
		 And M[1,1] = -2
		 And M[2,2] = 1 */
	M := rt.Matrix3(
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	)

	if M[0][0] != -3 || M[1][1] != -2 || M[2][2] != 1 {
		t.Errorf("Error: %v", M)
	}
}

func TestMatrix4Equal(t *testing.T) {
	/* Scenario: Matrix equality with identical matrices
	   Given the following matrix A:
		   | 1 | 2 | 3 | 4 |
		   | 5 | 6 | 7 | 8 |
		   | 9 | 8 | 7 | 6 |
		   | 5 | 4 | 3 | 2 |
		 And the following matrix B:
		   | 1 | 2 | 3 | 4 |
		   | 5 | 6 | 7 | 8 |
		   | 9 | 8 | 7 | 6 |
		   | 5 | 4 | 3 | 2 |
	   Then A = B */
	A := rt.Matrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)
	B := rt.Matrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	if !A.Equals(B) {
		t.Errorf("Error: %v != %v", A, B)
	}

}

func TestMatrix4NotEqual(t *testing.T) {
	/* Scenario: Matrix equality with different matrices
	   Given the following matrix A:
		   | 1 | 2 | 3 | 4 |
		   | 5 | 6 | 7 | 8 |
		   | 9 | 8 | 7 | 6 |
		   | 5 | 4 | 3 | 2 |
	   And the following matrix B:
		   | 2 | 3 | 4 | 5 |
		   | 6 | 7 | 8 | 9 |
		   | 8 | 7 | 6 | 5 |
		   | 4 | 3 | 2 | 1 |
	   Then A != B */
	A := rt.Matrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	B := rt.Matrix4(
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	)
	if A.Equals(B) {
		t.Errorf("Error: %v == %v", A, B)
	}
}

func TestMatrix4Mul(t *testing.T) {
	/* Scenario: Multiplying two matrices
	   Given the following matrix A:
		   | 1 | 2 | 3 | 4 |
		   | 5 | 6 | 7 | 8 |
		   | 9 | 8 | 7 | 6 |
		   | 5 | 4 | 3 | 2 |
		 And the following matrix B:
		   | -2 | 1 | 2 |  3 |
		   |  3 | 2 | 1 | -1 |
		   |  4 | 3 | 6 |  5 |
		   |  1 | 2 | 7 |  8 |
	   Then A * B is the following 4x4 matrix:
		   | 20|  22 |  50 |  48 |
		   | 44|  54 | 114 | 108 |
		   | 40|  58 | 110 | 102 |
		   | 16|  26 |  46 |  42 | */
	A := rt.Matrix4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	)

	B := rt.Matrix4(
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	)

	result := A.Mul(B)

	expected := rt.Matrix4(
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix4MulTuple(t *testing.T) {
	/* Scenario: A matrix multiplied by a tuple
	   Given the following matrix A:
		   | 1 | 2 | 3 | 4 |
		   | 2 | 4 | 4 | 2 |
		   | 8 | 6 | 4 | 1 |
		   | 0 | 0 | 0 | 1 |
		 And b ← tuple(1, 2, 3, 1)
	   Then A * b = tuple(18, 24, 33, 1) */
	A := rt.Matrix4(
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	)

	b := &rt.Tuple{
		1, 2, 3, 1,
	}

	result := A.MulT(b)

	expected := &rt.Tuple{
		18, 24, 33, 1,
	}
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix4MulIdentity(t *testing.T) {
	/* Scenario: Multiplying a matrix by the identity matrix
	   Given the following matrix A:
		 | 0 | 1 |  2 |  4 |
		 | 1 | 2 |  4 |  8 |
		 | 2 | 4 |  8 | 16 |
		 | 4 | 8 | 16 | 32 |
	   Then A * identity_matrix = A */
	A := rt.Matrix4(
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	)

	if !A.Equals(A.Mul(rt.Identity())) {
		t.Errorf("Error: %v", A)
	}
}

func TestMatrix4MulIdentityTuple(t *testing.T) {
	/* Scenario: Multiplying the identity matrix by a tuple
	   Given a ← tuple(1, 2, 3, 4)
	   Then identity_matrix * a = a */
	a := &rt.Tuple{1, 2, 3, 4}

	if !rt.Identity().MulT(a).Equals(a) {
		t.Errorf("Error: %v", rt.Identity().MulT(a))
	}
}

func TestMatrix4Transpose(t *testing.T) {
	/* Scenario: Transposing a matrix
	   Given the following matrix A:
		 | 0 | 9 | 3 | 0 |
		 | 9 | 8 | 0 | 8 |
		 | 1 | 8 | 5 | 3 |
		 | 0 | 0 | 5 | 8 |
	   Then transpose(A) is the following matrix:
		 | 0 | 9 | 1 | 0 |
		 | 9 | 8 | 8 | 0 |
		 | 3 | 0 | 5 | 5 |
		 | 0 | 8 | 3 | 8 | */
	A := rt.Matrix4(
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	)

	result := A.Trans()

	expected := rt.Matrix4(
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	)
	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix4TransposeIdentity(t *testing.T) {
	/* Scenario: Transposing the identity matrix
	   Given A ← transpose(identity_matrix)
	   Then A = identity_matrix */
	A := rt.Identity().Trans()

	if !A.Equals(rt.Identity()) {
		t.Errorf("Error: %v", A)
	}
}

func TestMatrix2Det(t *testing.T) {
	/* Scenario: Calculating the determinant of a 2x2 matrix
	   Given the following 2x2 matrix A:
		 |  1 | 5 |
		 | -3 | 2 |
	   Then determinant(A) = 17 */
	A := rt.Matrix2(
		1, 5,
		-3, 2,
	)

	result := A.Det()

	if result != 17 {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix3Submatrix(t *testing.T) {
	/* Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix
	   Given the following 3x3 matrix A:
		 |  1 | 5 |  0 |
		 | -3 | 2 |  7 |
		 |  0 | 6 | -3 |
	   Then submatrix(A, 0, 2) is the following 2x2 matrix:
		 | -3 | 2 |
		 |  0 | 6 | */
	A := rt.Matrix3(
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	)

	result := A.SubMatrix(0, 2)
	expected := rt.Matrix2(
		-3, 2,
		0, 6,
	)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}

}
func TestMatrix4Submatrix(t *testing.T) {
	/* Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix
	   Given the following 4x4 matrix A:
		 | -6 |  1 |  1 |  6 |
		 | -8 |  5 |  8 |  6 |
		 | -1 |  0 |  8 |  2 |
		 | -7 |  1 | -1 |  1 |
	   Then submatrix(A, 2, 1) is the following 3x3 matrix:
		 | -6 |  1 | 6 |
		 | -8 |  8 | 6 |
		 | -7 | -1 | 1 | */
	A := rt.Matrix4(
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	)

	result := A.SubMatrix(2, 1)
	expected := rt.Matrix3(
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	)

	if !result.Equals(expected) {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix3Minor(t *testing.T) {
	/* Scenario: Calculating a minor of a 3x3 matrix
	   Given the following 3x3 matrix A:
		   |  3 |  5 |  0 |
		   |  2 | -1 | -7 |
		   |  6 | -1 |  5 |
		 And B ← submatrix(A, 1, 0)
	   Then determinant(B) = 25
	 	And minor(A, 1, 0) = 25 */
	A := rt.Matrix3(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)

	result := A.Minor(1, 0)

	if result != 25 {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix3Cofactor(t *testing.T) {
	/* Scenario: Calculating a cofactor of a 3x3 matrix
	   Given the following 3x3 matrix A:
		   |  3 |  5 |  0 |
		   |  2 | -1 | -7 |
		   |  6 | -1 |  5 |
	   Then minor(A, 0, 0) = -12
		 And cofactor(A, 0, 0) = -12
		 And minor(A, 1, 0) = 25
		 And cofactor(A, 1, 0) = -25 */
	A := rt.Matrix3(
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	)

	result := A.Minor(0, 0)

	if result != -12 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 0)

	if result != -12 {
		t.Errorf("Error: %v", result)
	}

	result = A.Minor(1, 0)

	if result != 25 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(1, 0)

	if result != -25 {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix3Det(t *testing.T) {
	/* Scenario: Calculating the determinant of a 3x3 matrix
	   Given the following 3x3 matrix A:
		 |  1 |  2 |  6 |
		 | -5 |  8 | -4 |
		 |  2 |  6 |  4 |
	   Then cofactor(A, 0, 0) = 56
		 And cofactor(A, 0, 1) = 12
		 And cofactor(A, 0, 2) = -46
		 And determinant(A) = -196 */
	A := rt.Matrix3(
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	)

	result := A.Cofact(0, 0)

	if result != 56 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 1)

	if result != 12 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 2)

	if result != -46 {
		t.Errorf("Error: %v", result)
	}

	result = A.Det()

	if result != -196 {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix4Det(t *testing.T) {
	/* Scenario: Calculating the determinant of a 4x4 matrix
	   Given the following 4x4 matrix A:
		 | -2 | -8 |  3 |  5 |
		 | -3 |  1 |  7 |  3 |
		 |  1 |  2 | -9 |  6 |
		 | -6 |  7 |  7 | -9 |
	   Then cofactor(A, 0, 0) = 690
		 And cofactor(A, 0, 1) = 447
		 And cofactor(A, 0, 2) = 210
		 And cofactor(A, 0, 3) = 51
		 And determinant(A) = -4071 */
	A := rt.Matrix4(
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	)

	result := A.Cofact(0, 0)

	if result != 690 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 1)

	if result != 447 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 2)

	if result != 210 {
		t.Errorf("Error: %v", result)
	}

	result = A.Cofact(0, 3)

	if result != 51 {
		t.Errorf("Error: %v", result)
	}

	result = A.Det()

	if result != -4071 {
		t.Errorf("Error: %v", result)
	}
}

func TestMatrix4Invertible(t *testing.T) {
	/* Scenario: Testing an invertible matrix for invertibility
	   Given the following 4x4 matrix A:
		 |  6 |  4 |  4 |  4 |
		 |  5 |  5 |  7 |  6 |
		 |  4 | -9 |  3 | -7 |
		 |  9 |  1 |  7 | -6 |
	   Then determinant(A) = -2120
	 	And A is invertible */
	A := rt.Matrix4(
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6,
	)

	if A.Det() == 0 {
		t.Errorf("Error: %v != 0", A.Det())
	}
}
func TestMatrix4NotInvertible(t *testing.T) {
	/* Scenario: Testing a noninvertible matrix for invertibility
	   Given the following 4x4 matrix A:
		 | -4 |  2 | -2 | -3 |
		 |  9 |  6 |  2 |  6 |
		 |  0 | -5 |  1 | -5 |
		 |  0 |  0 |  0 |  0 |
	   Then determinant(A) = 0
	 	And A is not invertible */
	A := rt.Matrix4(
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	)

	if A.Det() != 0 {
		t.Errorf("Error: %v == 0", A.Det())
	}
}

func TestMatrix4Inverse(t *testing.T) {
	/* Scenario: Calculating the inverse of a matrix
	   Given the following 4x4 matrix A:
		   | -5 |  2 |  6 | -8 |
		   |  1 | -5 |  1 |  8 |
		   |  7 |  7 | -6 | -7 |
		   |  1 | -3 |  7 |  4 |
		 And B ← inverse(A)
	   Then determinant(A) = 532
		 And cofactor(A, 2, 3) = -160
		 And B[3,2] = -160/532
		 And cofactor(A, 3, 2) = 105
		 And B[2,3] = 105/532
		 And B is the following 4x4 matrix:
		   |  0.21805 |  0.45113 |  0.24060 | -0.04511 |
		   | -0.80827 | -1.45677 | -0.44361 |  0.52068 |
		   | -0.07895 | -0.22368 | -0.05263 |  0.19737 |
		   | -0.52256 | -0.81391 | -0.30075 |  0.30639 | */
	A := rt.Matrix4(
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	)
	B := A.Inv()

	if A.Det() != 532 {
		t.Errorf("Error: %v", A.Det())
	}

	if A.Cofact(2, 3) != -160 {
		t.Errorf("Error: %v", A.Cofact(2, 3))
	}

	if A.Cofact(3, 2) != 105 {
		t.Errorf("Error: %v", A.Cofact(3, 2))
	}

	expected := rt.Matrix4(
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	)

	if !B.Equals(expected) {
		t.Errorf("Error: %v", B)
	}
}
func TestMatrix4Inverse2(t *testing.T) {
	/* Scenario: Calculating the inverse of another matrix
	   Given the following 4x4 matrix A:
		 |  8 | -5 |  9 |  2 |
		 |  7 |  5 |  6 |  1 |
		 | -6 |  0 |  9 |  6 |
		 | -3 |  0 | -9 | -4 |
	   Then inverse(A) is the following 4x4 matrix:
		 | -0.15385 | -0.15385 | -0.28205 | -0.53846 |
		 | -0.07692 |  0.12308 |  0.02564 |  0.03077 |
		 |  0.35897 |  0.35897 |  0.43590 |  0.92308 |
		 | -0.69231 | -0.69231 | -0.76923 | -1.92308 | */
	A := rt.Matrix4(
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	)
	B := A.Inv()

	expected := rt.Matrix4(
		-0.15385, -0.15385, -0.28205, -0.53846,
		-0.07692, 0.12308, 0.02564, 0.03077,
		0.35897, 0.35897, 0.43590, 0.92308,
		-0.69231, -0.69231, -0.76923, -1.92308,
	)

	if !B.Equals(expected) {
		t.Errorf("Error: %v", B)
	}
}
func TestMatrix4Inverse3(t *testing.T) {
	/* Scenario: Calculating the inverse of a third matrix
	   Given the following 4x4 matrix A:
		 |  9 |  3 |  0 |  9 |
		 | -5 | -2 | -6 | -3 |
		 | -4 |  9 |  6 |  4 |
		 | -7 |  6 |  6 |  2 |
	   Then inverse(A) is the following 4x4 matrix:
		 | -0.04074 | -0.07778 |  0.14444 | -0.22222 |
		 | -0.07778 |  0.03333 |  0.36667 | -0.33333 |
		 | -0.02901 | -0.14630 | -0.10926 |  0.12963 |
		 |  0.17778 |  0.06667 | -0.26667 |  0.33333 | */
	A := rt.Matrix4(
		9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2,
	)
	B := A.Inv()

	expected := rt.Matrix4(
		-0.04074, -0.07778, 0.14444, -0.22222,
		-0.07778, 0.03333, 0.36667, -0.33333,
		-0.02901, -0.14630, -0.10926, 0.12963,
		0.17778, 0.06667, -0.26667, 0.33333,
	)

	if !B.Equals(expected) {
		t.Errorf("Error: %v", B)
	}

}

func TestMatrix4MulInverse(t *testing.T) {
	/* Scenario: Multiplying a product by its inverse
	   Given the following 4x4 matrix A:
		   |  3 | -9 |  7 |  3 |
		   |  3 | -8 |  2 | -9 |
		   | -4 |  4 |  4 |  1 |
		   | -6 |  5 | -1 |  1 |
		 And the following 4x4 matrix B:
		   |  8 |  2 |  2 |  2 |
		   |  3 | -1 |  7 |  0 |
		   |  7 |  0 |  5 |  4 |
		   |  6 | -2 |  0 |  5 |
		 And C ← A * B
	   Then C * inverse(B) = A */
	A := rt.Matrix4(
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	)
	B := rt.Matrix4(
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	)
	C := A.Mul(B)

	result := C.Mul(B.Inv())

	if !result.Equals(A) {
		t.Errorf("Error: %v", result)
	}
}

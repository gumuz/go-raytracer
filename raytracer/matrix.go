package raytracer

import (
	"math"
)

type Matrix2 [2][2]float64
type Matrix3 [3][3]float64
type Matrix4 [4][4]float64
type Tuple4 [4]float64

var Identity = NewMatrix4(
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
)

func NewMatrix2(values ...float64) Matrix2 {
	var matrix Matrix2
	for idx, val := range values {
		matrix[idx/2][idx%2] = val
	}
	return matrix
}

func NewMatrix3(values ...float64) Matrix3 {
	var matrix Matrix3
	for idx, val := range values {
		matrix[idx/3][idx%3] = val
	}
	return matrix
}

func NewMatrix4(values ...float64) Matrix4 {
	var matrix Matrix4
	for idx, val := range values {
		matrix[idx/4][idx%4] = val
	}
	return matrix
}

func NewTuple4(values ...float64) Tuple4 {
	var tuple Tuple4
	for idx, val := range values {
		tuple[idx] = val
	}
	return tuple
}

func (m Matrix2) Equals(b Matrix2) bool {
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			if math.Abs(m[y][x]-b[y][x]) > epsilon {
				return false
			}
		}
	}
	return true
}

func (m Matrix2) Determ() float64 {
	return (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
}

func (m Matrix3) Equals(b Matrix3) bool {
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if math.Abs(m[y][x]-b[y][x]) > epsilon {
				return false
			}
		}
	}
	return true
}

func (m Matrix3) SubMatrix(row, col int) Matrix2 {
	var matrix Matrix2

	idx := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if x == col || y == row {
				continue
			}
			matrix[idx/2][idx%2] = m[y][x]
			idx++
		}
	}

	return matrix
}

func (m Matrix3) Cofact(row, col int) float64 {
	matrix2 := m.SubMatrix(row, col)

	if (row+col)%2 == 0 {
		return matrix2.Determ()
	} else {
		return -matrix2.Determ()
	}
}

func (m Matrix3) Determ() float64 {
	return m[0][0]*m.Cofact(0, 0) + m[0][1]*m.Cofact(0, 1) + m[0][2]*m.Cofact(0, 2)
}

func (m Matrix4) Equals(b Matrix4) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if math.Abs(m[y][x]-b[y][x]) > epsilon {
				return false
			}
		}
	}
	return true
}

func (m Matrix4) Mult(b Matrix4) Matrix4 {
	var matrix Matrix4

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			for i := 0; i < 4; i++ {
				matrix[y][x] += m[y][i] * b[i][x]
			}
		}
	}

	return matrix
}

func (m Matrix4) MultP(p *Point) *Point {
	tupleIn := [4]float64{p.X, p.Y, p.Z, 1}
	tupleOut := [4]float64{}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tupleOut[y] += m[y][x] * tupleIn[x]
		}
	}

	return NewPoint(tupleOut[0], tupleOut[1], tupleOut[2])
}

func (m Matrix4) MultV(v *Vector) *Vector {
	tupleIn := [4]float64{v.X, v.Y, v.Z, 0}
	tupleOut := [4]float64{}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tupleOut[y] += m[y][x] * tupleIn[x]
		}
	}

	return NewVector(tupleOut[0], tupleOut[1], tupleOut[2])
}

func (m Matrix4) Trans() Matrix4 {
	var matrix Matrix4

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			matrix[x][y] = m[y][x]
		}
	}

	return matrix
}

func (m Matrix4) SubMatrix(row, col int) Matrix3 {
	var matrix Matrix3

	idx := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if x == col || y == row {
				continue
			}
			matrix[idx/3][idx%3] = m[y][x]
			idx++
		}
	}

	return matrix
}

func (m Matrix4) Cofact(row, col int) float64 {
	matrix3 := m.SubMatrix(row, col)

	if (row+col)%2 == 0 {
		return matrix3.Determ()
	} else {
		return -matrix3.Determ()
	}
}

func (m Matrix4) Determ() float64 {
	return m[0][0]*m.Cofact(0, 0) + m[0][1]*m.Cofact(0, 1) + m[0][2]*m.Cofact(0, 2) + m[0][3]*m.Cofact(0, 3)
}

func (m Matrix4) Inverse() Matrix4 {
	var matrix Matrix4

	det := m.Determ()
	if det == 0 {
		panic("Matrix is not invertible")
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			matrix[x][y] = m.Cofact(y, x) / m.Determ()
		}
	}

	return matrix
}

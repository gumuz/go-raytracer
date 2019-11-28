package raytracer

import "math"

type Matrix [][]float64

func Identity() Matrix {
	return Matrix4(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func matrix(size int, values ...float64) Matrix {
	matrix := make(Matrix, size)
	for idx, _ := range matrix {
		matrix[idx] = make([]float64, size)
	}

	for idx, value := range values {
		matrix[idx/size][idx%size] = value
	}

	return matrix
}

func Matrix2(values ...float64) Matrix {
	return matrix(2, values...)
}

func Matrix3(values ...float64) Matrix {
	return matrix(3, values...)
}

func Matrix4(values ...float64) Matrix {
	return matrix(4, values...)
}

func (m Matrix) Equals(b Matrix) bool {
	length := len(m)
	if length != len(b) {
		return false
	}

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			if math.Abs(m[row][col]-b[row][col]) > epsilon {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Mul(b Matrix) Matrix {
	length := len(m)
	matrix := matrix(length)

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			for i := 0; i < length; i++ {
				matrix[row][col] += m[row][i] * b[i][col]
			}
		}
	}

	return matrix
}

func (m Matrix) MulT(t *Tuple) *Tuple {
	in := [4]float64{t.X, t.Y, t.Z, t.W}
	out := [4]float64{}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			out[y] += m[y][x] * in[x]
		}
	}

	return &Tuple{out[0], out[1], out[2], out[3]}
}

func (m Matrix) Trans() Matrix {
	length := len(m)
	matrix := matrix(length)

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			matrix[col][row] = m[row][col]
		}
	}

	return matrix
}

func (m Matrix) Det() float64 {
	length := len(m)
	det := 0.0

	if length == 2 {
		det = (m[0][0] * m[1][1]) - (m[0][1] * m[1][0])
	} else {
		for col := 0; col < length; col++ {
			det += m[0][col] * m.Cofact(0, col)
		}
	}

	return det
}

func (m Matrix) SubMatrix(row, col int) Matrix {
	values := []float64{}

	for rowIdx, cols := range m {
		for colIdx, value := range cols {
			if rowIdx == row || colIdx == col {
				continue
			}
			values = append(values, value)
		}
	}

	return matrix(len(m)-1, values...)
}

func (m Matrix) Minor(row, col int) float64 {
	matrix := m.SubMatrix(row, col)

	return matrix.Det()
}

func (m Matrix) Cofact(row, col int) float64 {
	if (row+col)%2 == 0 {
		return m.Minor(row, col)
	} else {
		return -m.Minor(row, col)
	}
}

func (m Matrix) Inv() Matrix {
	matrix := matrix(len(m))

	det := m.Det()
	if det == 0 {
		panic("Matrix is not invertible")
	}

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			matrix[col][row] = m.Cofact(row, col) / det
		}
	}

	return matrix
}

package linear

import "fmt"

const (
	eps = float32(1E-8)
)

func clamp(value float32) float32 {
	// Close enough to 0
	if value < eps && value > -eps {
		return 0
	}

	// Close enough to 1
	if value < (1+eps) && value > (1-eps) {
		return 1
	}

	return value
}

func swap(m Matrix, rowA, rowB int) {
	_, cols := m.Dim()
	for col := 0; col < cols; col++ {
		tmp := m.Get(rowA, col)
		m.Set(rowA, col, m.Get(rowB, col))
		m.Set(rowB, col, tmp)
	}
}

func divide(m Matrix, row int, divisor float32) {
	_, cols := m.Dim()
	for col := 0; col < cols; col++ {
		value := m.Get(row, col) / divisor
		m.Set(row, col, value)
	}
}

// rowA = rowA - rowB
func subtract(m Matrix, rowA, rowB int, value float32) {
	_, cols := m.Dim()
	for col := 0; col < cols; col++ {
		value := m.Get(rowA, col) - value*m.Get(rowB, col)
		m.Set(rowA, col, value)
	}
}

// Inverse returns the inverse of the matrix if possible
func Inverse(m Matrix) (Matrix, error) {
	rows, cols := m.Dim()
	if rows != cols {
		return m, fmt.Errorf("Non square matrix, cannot invert")
	}

	tmp := m.Clone()

	result := NewMatrix(rows, cols)

	// make identity
	for r := 0; r < rows; r++ {
		result.Set(r, r, 1)
	}

	for col := 0; col < cols; col++ {

		// Find first non-zero
		if tmp.Get(col, col) == 0 {
			for row := col; row < rows; row++ {
				if tmp.Get(row, col) != 0 {
					swap(tmp, col, row)
					swap(result, col, row)
					break
				}
			}
		}

		if tmp.Get(col, col) == 0 {
			return m, fmt.Errorf("Nonsingular matrix")
		}
		value := tmp.Get(col, col)
		divide(tmp, col, value)
		divide(result, col, value)

		for row := 0; row < rows; row++ {
			if (row != col) && tmp.Get(row, col) != 0 {
				value := tmp.Get(row, col) / tmp.Get(col, col)

				subtract(tmp, row, col, value) // row = row - value * col
				subtract(result, row, col, value)
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result.Set(r, c, clamp(result.Get(r, c)))
		}
	}

	return result, nil
}

// Determinant of the matrix
func Determinant(m Matrix) (float32, error) {
	rows, cols := m.Dim()
	if rows != cols {
		return 0, fmt.Errorf("Non-singular matrix")
	}

	tmp := m.Clone()

	result := float32(1)

	for col := 0; col < cols; col++ {

		// Find first non-zero
		if tmp.Get(col, col) == 0 {
			for row := col; row < rows; row++ {
				if tmp.Get(row, col) != 0 {
					swap(tmp, col, row)
					result *= -1
					break
				}
			}
		}

		if tmp.Get(col, col) == 0 {
			return 0, fmt.Errorf("Nonsingular matrix")
		}
		value := tmp.Get(col, col)
		divide(tmp, col, value)
		result *= value

		for row := 0; row < rows; row++ {
			if (row != col) && tmp.Get(row, col) != 0 {
				value := tmp.Get(row, col) / tmp.Get(col, col)

				subtract(tmp, row, col, value) // row = row - value * col
			}
		}
	}

	return result, nil
}

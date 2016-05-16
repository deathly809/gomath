package linear

import (
	"bytes"
	"fmt"
)

type mat struct {
	rows, cols int
	data       []Vector
	transposed bool
}

func (m *mat) apply(f func(int, int)) {
	rows, cols := m.Dim()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			f(r, c)
		}
	}
}

func (m *mat) Get(row, col int) float32 {
	if m.transposed {
		return m.data[col].Get(row)
	}
	return m.data[row].Get(col)
}

func (m *mat) Set(row, col int, value float32) {
	if m.transposed {
		m.data[col].Set(row, value)
	} else {
		m.data[row].Set(col, value)
	}
}

func (m *mat) Dim() (int, int) {
	if m.transposed {
		return m.cols, m.rows
	}
	return m.rows, m.cols
}

func dot(a, b Matrix, r, c, l int) float32 {
	result := float32(0.0)
	for i := 0; i < l; i++ {
		result += a.Get(r, i) * b.Get(i, c)
	}
	return result
}

func (m *mat) Mult(o Matrix) {
	rows := m.rows
	_, cols := o.Dim()
	tmp := NewMatrix(rows, cols)
	m.apply(func(r, c int) {
		val := dot(m, o, r, c, m.cols)
		tmp.Set(r, c, val)
	})
	m.apply(func(r, c int) {
		m.Set(r, c, tmp.Get(r, c))
	})
}

func (m *mat) Add(o Matrix) {
	m.apply(func(r, c int) {
		m.Set(r, c, m.Get(r, c)+o.Get(r, c))
	})
}

func (m *mat) Sub(o Matrix) {
	m.apply(func(r, c int) {
		m.Set(r, c, m.Get(r, c)-o.Get(r, c))
	})
}

func (m *mat) Transpose() {
	m.transposed = !m.transposed
}

func (m *mat) Clone() Matrix {
	result := NewMatrix(m.rows, m.cols)
	m.apply(func(r, c int) {
		result.Set(r, c, m.Get(r, c))
	})
	return result
}

func (m *mat) String() string {
	var buffer bytes.Buffer
	rows, cols := m.Dim()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			buffer.WriteString(fmt.Sprintf("%5f ", m.Get(r, c)))
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// NewMatrix returns a new matrix with the dimensions
// specified and all entries set to 0
func NewMatrix(rows, cols int) Matrix {
	result := &mat{
		rows: rows,
		cols: cols,
		data: make([]Vector, rows),
	}

	for r := 0; r < rows; r++ {
		result.data[r] = NewVector(cols)
	}

	return result
}

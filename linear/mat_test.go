package linear

import (
	"math/rand"
	"testing"
)

func randomMatrix(rows, cols int) Matrix {
	result := NewMatrix(rows, cols)
	num := rand.Intn(rows * cols)
	for num > 0 {
		r := rand.Intn(rows)
		c := rand.Intn(cols)
		for result.Get(r, c) != 0 {
			r = rand.Intn(rows)
			c = rand.Intn(cols)
		}
		result.Set(r, c, rand.Float32())
		num--
	}
	return result
}

func TestTranspose(t *testing.T) {
	maxDim := 500
	iter := 100
	for iter > 0 {
		rows, cols := 1+rand.Intn(maxDim), 1+rand.Intn(maxDim)

		mat := randomMatrix(rows, cols)

		cloned := mat.Clone()
		cloned.Transpose()

		newRows, newCols := cloned.Dim()

		if newRows != cols {
			t.Fatalf("Expected %d but found %d", cols, newRows)
		}

		if newCols != rows {
			t.Fatalf("Expected %d but found %d", cols, newCols)
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				expected := mat.Get(r, c)
				found := cloned.Get(c, r)
				if expected != found {
					t.Logf("(%d,%d)->%f does not equal (%d,%d)->%f\n", r, c, expected, c, r, found)
					t.Fail()
				}
			}
		}
		iter--
	}
}

func TestVectorDot(t *testing.T) {

	dim := 100
	a := NewVector(dim)
	b := NewVector(dim)
	expected := float32(0)

	for i := 0; i < dim; i++ {
		v1 := rand.Float32()
		v2 := rand.Float32()
		a.Set(i, v1)
		b.Set(i, v2)
		expected += v1 * v2
	}

	found := a.Dot(b)

	if expected != found {
		t.Logf("Expected %f but found %f\n", expected, found)
		t.Fail()
	}
}

func TestVectorDotFailure(t *testing.T) {
	dim := 100
	a := NewVector(dim + 100)
	b := NewVector(dim)
	expected := float32(0)
	found := a.Dot(b)

	if expected != found {
		t.Logf("Expected %f but found %f\n", expected, found)
		t.Fail()
	}
}

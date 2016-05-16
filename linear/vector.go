package linear

import "math"

type vector struct {
	len  int
	data []float32
	mag  float32
}

func (v vector) Get(pos int) float32 {
	return v.data[pos]
}

func (v vector) Set(pos int, value float32) {
	v.data[pos] = value
}

func (v vector) Dim() int {
	return v.len
}

func (v vector) Dot(o Vector) (result float32) {
	for i := 0; i < v.len; i++ {
		result += v.Get(i) * o.Get(i)
	}
	return
}

func (v vector) Normalize() {
	mag := v.Magnitude()
	if mag > 0 {
		for i := 0; i < v.len; i++ {
			v.data[i] /= mag
		}
	}
}

func (v vector) Magnitude() (result float32) {
	if v.mag < 0 {
		v.mag = 0
		for i := 0; i < v.len; i++ {
			val := v.Get(i)
			v.mag += val * val
		}
		v.mag = float32(math.Sqrt(float64(result)))
	}
	return v.mag

}

func (v vector) Clone() Vector {
	result := NewVector(v.Dim())
	for i := 0; i < v.len; i++ {
		result.Set(i, v.Get(i))
	}
	return result
}

func (v vector) Add(other Vector) {
	for i := 0; i < v.len; i++ {
		v.Set(i, v.Get(i)+other.Get(i))
	}
}

func (v vector) Sub(other Vector) {
	for i := 0; i < v.len; i++ {
		v.Set(i, v.Get(i)-other.Get(i))
	}
}

// NewVector creates a new Vector with the specified length
func NewVector(dim int) Vector {
	return &vector{
		len:  dim,
		data: make([]float32, dim),
		mag:  0,
	}
}

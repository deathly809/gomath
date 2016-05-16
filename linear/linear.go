package linear

// Matrix is an abstract matrix object
type Matrix interface {
	Get(int, int) float32
	Set(int, int, float32)

	Mult(Matrix)
	Add(Matrix)
	Sub(Matrix)

	Dim() (int, int)
	Clone() Matrix
	Transpose()
}

// Vector is a vector
type Vector interface {
	Get(int) float32
	Set(int, float32)

	Add(Vector)
	Sub(Vector)

	Dim() int
	Dot(Vector) float32
	Clone() Vector
	Normalize()
}

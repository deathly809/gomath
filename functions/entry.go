package functions

// Polynomial interface
type Polynomial interface {
	Eval(float32) float32
	GetCoeff(int) float32
	SetCoeff(int, float32)
	Degree() int
}

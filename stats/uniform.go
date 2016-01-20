package stats

/*
   Created using
   https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)
*/

import (
	"math"

	"github.com/deathly809/gomath"
	"github.com/deathly809/gomath/stats"
)

// Uniform holds all information about the distribution
type Uniform struct {
	a, b             float64
	pdf              float64
	mean, median     float64
	variance, stddev float64
}

// Pdf will assume that the value passed in is an integer
func (u *Uniform) Pdf(x float64) float64 {
	result := 0.0
	if x >= u.a && x <= u.b {
		result = u.pdf
	}
	return result
}

// Cdf will assume that the value passed in is an integer
func (u *Uniform) Cdf(x float64) float64 {
	x = gomath.ClampFloat64(u.a, u.b, x)
	return (x - u.a) / (u.b - u.a)
}

// Mean of the discrete uniform distribution
func (u *Uniform) Mean() float64 {
	return u.mean
}

// Median of the discrete uniform distribution
func (u *Uniform) Median() float64 {
	return u.median
}

// Variance of the discrete uniform distribution
func (u *Uniform) Variance() float64 {
	return u.variance
}

// StdDev of the discrete uniform distribution
func (u *Uniform) StdDev() float64 {
	return u.stddev
}

// NewUniform creates a continuous uniform distribution
func NewUniform(low, high int) stats.Distribution {
	a := gomath.MinInt(low, high)
	b := gomath.MaxInt(low, high)

	return &Uniform{
		a:        float64(a),
		b:        float64(b),
		pdf:      1.0 / float64(b-a),
		mean:     0.5 * float64(a+b),
		median:   0.5 * float64(a+b),
		variance: (math.Pow(float64(b-a), 2) - 1.0) / 12.0,
		stddev:   math.Pow(float64(b-a+1), 2) / math.Sqrt(12.0),
	}
}

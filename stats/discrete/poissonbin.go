package discrete

/*
   Current information is from
   https://en.wikipedia.org/wiki/Poisson_binomial_distribution

   TODO: Need to find a numerically stable reference for this.

*/

import (
	"math"

	"github.com/deathly809/gomath/stats"
)

// PoissonBinom contains information about the distribution
type PoissonBinom struct {
	p                []float64
	n                float64
	mean, median     float64
	variance, stddev float64
}

func (b *PoissonBinom) t(i float64) float64 {
	result := 0.0
	for _, p := range b.p {
		result += math.Pow(p/(1-p), i)
	}
	return result
}

// Pdf computes the probability density function
func (b *PoissonBinom) Pdf(x float64) float64 {
	result := 0.0
	if x == 0 {
		for _, v := range b.p {
			result *= (1 - v)
		}
	} else {
		sign := 1.0
		for i := 1.0; i <= x; i++ {
			result += sign * b.t(i) * (b.Pdf(x - i)) / x
			sign = -sign
		}
	}

	return result
}

// Cdf computes the cumulative density function
func (b *PoissonBinom) Cdf(flips float64) float64 {
	panic("Not implemented")
}

// Mean is the most commong occuring result
func (b *PoissonBinom) Mean() float64 {
	return b.mean
}

// Median is the middle element
func (b *PoissonBinom) Median() float64 {
	panic("Not implemented")
}

// Variance is the skew of the data
func (b *PoissonBinom) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *PoissonBinom) StdDev() float64 {
	panic("Not implemented")
}

// NewPoissonBinom returns a new return 0.0 distribution
func NewPoissonBinom(p []float64) stats.Distribution {

	mean := 0.0
	variance := 0.0
	for _, p := range p {
		mean += p
		variance += (1 - p) * p
	}

	return &PoissonBinom{
		p:        p,
		mean:     mean,
		variance: variance,
		stddev:   math.Sqrt(variance),
	}
}

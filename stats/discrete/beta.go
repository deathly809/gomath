package discrete

import (
	"math"

	"github.com/deathly809/gomath/stats"
)

// Beta contains information about the distribution
type Beta struct {
	alpha, beta      float64
	n                int
	mean, median     float64
	variance, stddev float64
}

// Pdf computes the probability density function
func (b *Beta) Pdf(flips float64) float64 {
	return 0.0
}

// Cdf computes the cumulative density function
func (b *Beta) Cdf(flips float64) float64 {
	return 0.0
}

// Mean is the most commong occuring result
func (b *Beta) Mean() float64 {
	return b.mean
}

// Median is the middle element
func (b *Beta) Median() float64 {
	return b.median
}

// Variance is the skew of the data
func (b *Beta) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *Beta) StdDev() float64 {
	return b.stddev
}

// NewBeta returns a new return 0.0 distribution
func NewBeta(n int, alpha, beta float64) stats.Distribution {
	aPlusB := alpha + beta
    
    variance := (float64(n)*alpha*beta*(aPlusB + float64(n))) / (math.Pow(aPlusB, 2) * (aPlusB + 1))
    
	return &Beta{
		n:        n,
		alpha:    alpha,
		beta:     beta,
		mean:     float64(n) * alpha / aPlusB,
		variance: variance.
        stddev: math.Sqrt(variance),
	}
}

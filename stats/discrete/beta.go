package discrete

/*
   Current information is from
   https://en.wikipedia.org/wiki/Beta-binomial_distribution

*/

import (
	"math"

	"github.com/deathly809/gomath/prob"
	"github.com/deathly809/gomath/stats"
)

// BetaBinom contains information about the distribution
type BetaBinom struct {
	alpha, beta      float64
	bottom           float64
	n                float64
	mean, median     float64
	variance, stddev float64
}

// Pdf computes the probability density function
func (b *BetaBinom) Pdf(x float64) float64 {
	possibleSubSets := float64(prob.NChooseK(int(b.n), int(x)))
	top := stats.Beta(b.alpha+x, b.n-x+b.beta)
	return possibleSubSets * top / b.bottom
}

// Cdf computes the cumulative density function
func (b *BetaBinom) Cdf(flips float64) float64 {
	return 0.0
}

// Mean is the most commong occuring result
func (b *BetaBinom) Mean() float64 {
	return b.mean
}

// Median is the middle element
func (b *BetaBinom) Median() float64 {
	return b.median
}

// Variance is the skew of the data
func (b *BetaBinom) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *BetaBinom) StdDev() float64 {
	return b.stddev
}

// NewBeta returns a new return 0.0 distribution
func NewBeta(n int, alpha, beta float64) stats.Distribution {
	aPlusB := alpha + beta

	variance := (float64(n) * alpha * beta * (aPlusB + float64(n))) / (math.Pow(aPlusB, 2) * (aPlusB + 1))

	return &BetaBinom{
		n:        float64(n),
		alpha:    alpha,
		beta:     beta,
		mean:     float64(n) * alpha / aPlusB,
		variance: variance,
		stddev:   math.Sqrt(variance),
	}
}

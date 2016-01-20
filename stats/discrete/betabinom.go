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
	result := 0.0
	for i := 0.0; i <= flips; i++ {
		result += b.Pdf(i)
	}
	return result
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

// NewBetaBinom returns a new return 0.0 distribution
func NewBetaBinom(n int, alpha, beta float64) stats.Distribution {
	aPlusB := alpha + beta

	variance := (float64(n) * alpha * beta * (aPlusB + float64(n))) / (math.Pow(aPlusB, 2) * (aPlusB + 1))

	result := &BetaBinom{
		bottom:   stats.Beta(alpha, beta),
		n:        float64(n),
		alpha:    alpha,
		beta:     beta,
		mean:     float64(n) * alpha / aPlusB,
		variance: variance,
		stddev:   math.Sqrt(variance),
	}

	left, right := 0.0, float64(n)
	for left < right-1 {
		mid := math.Floor((left + right + 1) * 0.5)

		prob := result.Cdf(mid)
		if prob < 0.5 {
			left = mid
		} else {
			right = mid
		}
	}

	result.median = (left + right) * 0.5

	return result
}

package discrete

/*
   Information found at:
   https://en.wikipedia.org/wiki/Binomial_distribution
*/

import (
	"math"

	"github.com/deathly809/gomath/prob"
	"github.com/deathly809/gomath/stats"
)

// Binomial contains information about the distribution
type Binomial struct {
	p                float64
	n                int
	mean, median     float64
	variance, stddev float64
}

// Pdf computes the probability density function
func (b *Binomial) Pdf(flips float64) float64 {
	x := int(flips)

	numberOfSubsets := prob.NChooseK(b.n, x)
	probHeads := math.Pow(b.p, float64(x))
	probTails := math.Pow(b.p, float64(b.n-x))

	return float64(numberOfSubsets) * probHeads * probTails
}

// Cdf computes the cumulative density function
func (b *Binomial) Cdf(flips float64) float64 {
	result := 0.0
	for i := 0.0; i < flips; i++ {
		result += b.Pdf(i)
	}
	return result
}

// Mean is the most commong occuring result
func (b *Binomial) Mean() float64 {
	return b.mean
}

// Median is the middle element
func (b *Binomial) Median() float64 {
	return b.median
}

// Variance is the skew of the data
func (b *Binomial) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *Binomial) StdDev() float64 {
	return b.stddev
}

// NewBinom returns a new binomial distribution
func NewBinom(n int, p float64) stats.Distribution {
	return &Binomial{
		n:        n,
		p:        p,
		mean:     p * float64(n),
		median:   math.Floor(p * float64(n)),
		variance: p * float64(n) * (1 - p),
		stddev:   math.Sqrt(float64(n) * (1 - p)),
	}
}

package stats

/*
   Created using
   https://en.wikipedia.org/wiki/Normal_distribution
*/

import (
	"math"
)

// Geometric holds all information related to a normal distribution
type Geometric struct {
	p        float64
	mean     float64
	median   float64
	stddev   float64
	variance float64
}

// Pdf is the probability density function.  If the distributions
// is discrete then this is the probability mass function
func (n *Geometric) Pdf(x float64) float64 {
	return n.p * math.Pow(1-n.p, x-1)
}

// Cdf is the cumulative density function
func (n *Geometric) Cdf(x float64) float64 {
	return 1 - math.Pow(1-n.p, x)
}

// Mean is the mean of the distribution
func (n *Geometric) Mean() float64 {
	return n.mean
}

// Median is the median of the distribution
func (n *Geometric) Median() float64 {
	return n.mean
}

// Variance is the variance of the distribution
func (n *Geometric) Variance() float64 {
	return n.variance
}

// StdDev is the variance of the distribution
func (n *Geometric) StdDev() float64 {
	return n.stddev
}

// NewGeometric creates a new geometric distribution
func NewGeometric(p float64) Distribution {
	return &Geometric{
		p:        p,
		mean:     1 / p,
		median:   math.Ceil(-1 / math.Log2(1-p)),
		variance: (1 - p) / (p * p),
		stddev:   math.Sqrt(1-p) / p,
	}
}

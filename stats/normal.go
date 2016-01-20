package stats

/*
   Created using
   https://en.wikipedia.org/wiki/Normal_distribution
*/

import (
	"math"
)

const (
	// SqrtTwoPi is the square root of 2 times Pi
	SqrtTwoPi = math.Sqrt(2.0 * math.Pi)
)

// Normal holds all information related to a normal distribution
type Normal struct {
	mean     float64
	stddev   float64
	variance float64
}

// Pdf is the probability density function.  If the distributions
// is discrete then this is the probability mass function
func (n *Normal) Pdf(x float64) float64 {
	exp := -math.pow(x-n.mean, 2) / (2 * n.variance)
	return math.Exp(exp) / (n.stddev * SqrtTwoPi)
}

// Cdf is the cumulative density function
func (n *Normal) Cdf(x float64) float64 {
	return 0.5 * (1 + math.Erf((x-n.mean)/(n.stddev*math.Sqrt2)))
}

// Mean is the mean of the distribution
func (n *Normal) Mean() float64 {
	return n.mean
}

// Median is the median of the distribution
func (n *Normal) Median() float64 {
	return n.mean
}

// Variance is the variance of the distribution
func (n *Normal) Variance() float64 {
	return n.variance
}

// StdDev is the variance of the distribution
func (n *Normal) StdDev() float64 {
	return n.stddev
}

// NewNormal creates a new normal distribution
func NewNormal(mean, variance float64) Distribution {
	return &Normal{
		mean:     mean,
		variance: variance,
		stddev:   math.Sqrt(variance),
	}
}

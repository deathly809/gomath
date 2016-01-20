package stats

import (
	"math"

	"github.com/deathly809/gomath/prob"
    "github.com/deathly809/gomath/stats"
)

// Bernoulli contains information about the distribution
type Bernoulli struct {
	p                float64
	median     float64
	variance, stddev float64
}

// Pdf computes the probability density function
func (b *Bernoulli) Pdf(flips float64) float64 {
	result := 0.0
    switch int(flips) {
    case 0:
        result = p
    case 1:
        result = 1 - p
    }
    return result
}

// Cdf computes the cumulative density function
func (b *Bernoulli) Cdf(flips float64) float64 {
	result := 0.0
	for i := 0.0; i < flips; i++ {
		result += b.Pdf(i)
	}
	return result
}

// Mean is the most commong occuring result
func (b *Bernoulli) Mean() float64 {
	return b.p
}

// Median is the middle element
func (b *Bernoulli) Median() float64 {
	return b.median
}

// Variance is the skew of the data
func (b *Bernoulli) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *Bernoulli) StdDev() float64 {
	return b.stddev
}

// NewBernoulli returns a new Bernoulli distribution
func NewBernoulli(p float64) stats.Distribution {
    med := 0.0
    switch {
        case p == 0.5:
            med = 0.5
        case p > 0.5: 
            med = 1.0
    }
     
	return &Bernoulli{
		p:        p,
		median: med,   
		variance: p * (1-            p),
		stddev:   1.0,
	}
}

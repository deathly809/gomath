package stats

// Distribution contains the required functions for all distributions
type Distribution interface {

	// Pdf is the probability density function.  If the distributions
	// is discrete then this is the probability mass function
	Pdf(float64) float64

	// Cdf is the cumulative density function
	Cdf(float64) float64

	// Mean is the mean of the distribution
	Mean() float64

	// Median is the median of the distribution
	Median() float64

	// Variance is the variance of the distribution
	Variance() float64

	// StdDev is the variance of the distribution
	StdDev() float64
}

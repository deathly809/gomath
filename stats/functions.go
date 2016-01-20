package stats

import (
	"math"
)

// Beta function
func Beta(x, y float64) float64 {
	return math.Gamma(x) * math.Gamma(y) / math.Gamma(x+y)
}

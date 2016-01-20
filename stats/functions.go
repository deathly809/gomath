package stats

import (
	"math"
)

// Beta function
func Beta(x, y float64) float64 {
	return math.Gamma(x) * math.Gamma(y) / math.Gamma(x+y)
}

// Pochhammer is some junk?
func Pochhammer(j int, u float64) float64 {
	result := 1.0
	for u < j {
		result *= u
		u++
	}
	return result
}

// GeneralHyperGeom is the General Hypergeometric function
func GeneralHyperGeom(p, q int, z float64, params ...float64) float64 {
	result := 1.0
	top := 1.0
	bottom := 1.0
	pow := 1.0
	for j := 1; j < 100; j++ {
		result += (top * pow) / (bottom * fact)

		for i := 0; i < p; i++ {
			top *= params[i]
		}

		for i := 0; i < q; i++ {
			bottom *= params[p+i]
		}
		pow *= z
		fact *= j
	}
	return result
}

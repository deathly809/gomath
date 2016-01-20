package discrete

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestBetaBinom(t *testing.T) {
	n := rand.Intn(100)
	alpha, beta := 34.1350, 31.6085
	bDist := NewBetaBinom(n, alpha, beta)

	totalProb := 0.0

	for i := 0.0; i <= float64(n+n); i++ {
		pdf := bDist.Pdf(i)
		totalProb += pdf
		cdf := bDist.Cdf(i)
		if math.Abs(cdf-totalProb) > 1e-5 {
			t.Log("expected: ", totalProb, " found", cdf)
			t.Fail()
			return
		}
		fmt.Println(cdf)
	}

	if totalProb < (1.0 - 1e-5) {
		t.Log("expected ~1.0, but found ", totalProb)
		t.Fail()
		return
	}

}

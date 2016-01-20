package discrete

import "testing"

func TestBetaSymmetric(t *testing.T) {
	expectedMedian := 0.5
	beta := NewBeta(10, 0.5, 0.5)
	if expectedMedian != beta.Median() {
		t.Log(beta.Median())
		t.Fail()
		return
	}

}

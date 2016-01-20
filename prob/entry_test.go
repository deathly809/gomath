package prob

import (
	"fmt"
	"testing"
)

func slowNChooseK(n, k int) int {
	result := 1
	if k > 0 && k < n {
		result = slowNChooseK(n-1, k-1) + slowNChooseK(n-1, k)
	}
	return result
}

func TestNChooseK(t *testing.T) {
	N := 20
	for K := 1; K <= N; K++ {
		expected := slowNChooseK(N, K)
		found := NChooseK(N, K)
		if found != expected {
			fmt.Printf("Witb N = %d and K = %d we expected %d but found %d\n", N, K, expected, found)
			t.Fail()
			break
		}
	}
}

func BenchmarkNChooseKSmall(b *testing.B) {
	b.StopTimer()
	expected := slowNChooseK(20, 5)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		found := NChooseK(20, 5)
		if found != expected {
			fmt.Printf("Witb N = 20 and K = 5 we expected %d but found %d\n", expected, found)
			break
		}
	}
}

func BenchmarkNChooseKLarge(b *testing.B) {
	b.StopTimer()
	expected := 30045015
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		found := NChooseK(30, 20)
		if found != expected {
			fmt.Printf("Witb N = 20 and K = 5 we expected %d but found %d\n", expected, found)
			break
		}
	}
}

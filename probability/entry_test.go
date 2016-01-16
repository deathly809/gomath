package probability

import "testing"

func slowNChooseK(n, k int) int {
    result := 1
    if k > 1 && k < n {
        result = slowNChooseK(n-1,k-1) + slowNChooseK(n-1,k)
    }
    return result
}

func TestNChooseK(t *testing.T) {
    N , K := 10 , 5
    expected := slowNChooseK(N,K)
    have := NChooseK(N, K)
    if have != expected {
        t.Fail()
    }
}
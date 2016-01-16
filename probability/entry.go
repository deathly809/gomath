package probability

import (
    "github.com/deathly809/gomath"
)

// NChooseK returns n!/(k!*(n-k)!)
func NChooseK(n, k int) int {
    result := 1
    
    k = gomath.MaxInt(k,n-k)
    
    for i := 1 ; i <= k ; i++ {
        result = result / i + result % i * n / i
    }
    
    return result
}
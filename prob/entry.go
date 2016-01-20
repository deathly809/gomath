package prob

// NChooseK returns n!/(k!*(n-k)!)
func NChooseK(n, k int) int {
    c := 1
    
    if k > (n-k) {
        k = n - k
    }
    
    for i := 0 ; i < k ; i++ {
        c *= (n-i)
        c /= (i+1)
    }
    
    return c
}
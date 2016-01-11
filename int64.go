
package gomath

// MaxInt64 takes in two or more integers and returns
// the maximum of them
func MaxInt64(a int64, rem ...int64) int64 {
    res := a
    for _,v := range rem {
	if v > res {
           res = v
        }
    }
    return res
}

// MinInt64 takes in two or more integers and returns
// the maximum of them
func MinInt64(a int64, rem ...int64) int64 {
    res := a
    for _,v := range rem {
	if v < res {
           res = v
        }
    }
    return res
}

// AbsInt64 takes an integer and returns its absolute 
// value
func AbsInt64(a int64) int64 {
    return MaxInt64(a,-a)
}
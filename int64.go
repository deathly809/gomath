
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

// ClampInt64 will clamp a given value between a low and 
// high value inclusively.
func ClampInt64(low, high, val int64) (result int64) {
    if val < low {
        result = low
    }else if val > high {
        result = high
    }else {
        result = val
    }
    return
}

// ScaleInt64 will scale a number from the old range
// to the new range
func ScaleInt64(low, high, oldLow, oldHigh, value int64) int64 {
	return low + high*((value-oldLow)/(oldHigh - oldLow))
}
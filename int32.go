
package gomath

// MaxInt32 takes in two or more integers and returns
// the maximum of them
func MaxInt32(a int32, rem ...int32) int32 {
    res := a
    for _,v := range rem {
	if v > res {
           res = v
        }
    }
    return res
}

// MinInt32 takes in two or more integers and returns
// the maximum of them
func MinInt32(a int32, rem ...int32) int32 {
    res := a
    for _,v := range rem {
	if v < res {
           res = v
        }
    }
    return res
}

// AbsInt32 takes an integer and returns its absolute 
// value
func AbsInt32(a int32) int32 {
    return MaxInt32(a,-a)
}

// ClampInt32 will clamp a given value between a low and 
// high value inclusively.
func ClampInt32(low, high, val int32) (result int32) {
    if val < low {
        result = low
    }else if val > high {
        result = high
    }else {
        result = val
    }
    return
}
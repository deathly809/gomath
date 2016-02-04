
package gomath

// MaxInt takes in two or more integers and returns
// the maximum of them
func MaxInt(a int, rem ...int) int {
    res := a
    for _,v := range rem {
	if v > res {
           res = v
        }
    }
    return res
}

// MinInt takes in two or more integers and returns
// the maximum of them
func MinInt(a int, rem ...int) int {
    res := a
    for _,v := range rem {
	if v < res {
           res = v
        }
    }
    return res
}

// AbsInt takes an integer and returns its absolute 
// value
func AbsInt(a int) int {
    return MaxInt(a,-a)
}

// ClampInt will clamp a given value between a low and 
// high value inclusively.
func ClampInt(low, high, val int) (result int) {
    if val < low {
        result = low
    }else if val > high {
        result = high
    }else {
        result = val
    }
    return
}

// ScaleInt will scale a number from the old range
// to the new range
func ScaleInt(low, high, oldLow, oldHigh, value int) int {
	return low + high*((value-oldLow)/(oldHigh - oldLow))
}
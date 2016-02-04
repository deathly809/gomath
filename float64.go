
package gomath

// MaxFloat64 takes in two or more integers and returns
// the maximum of them
func MaxFloat64(a float64, rem ...float64) float64 {
    res := a
    for _,v := range rem {
	if v > res {
           res = v
        }
    }
    return res
}

// MinFloat64 takes in two or more integers and returns
// the maximum of them
func MinFloat64(a float64, rem ...float64) float64 {
    res := a
    for _,v := range rem {
	if v < res {
           res = v
        }
    }
    return res
}

// AbsFloat64 takes an integer and returns its absolute 
// value
func AbsFloat64(a float64) float64 {
    return MaxFloat64(a,-a)
}

// ClampFloat64 will clamp a given value between a low and 
// high value inclusively.
func ClampFloat64(low, high, val float64) (result float64) {
    if val < low {
        result = low
    }else if val > high {
        result = high
    }else {
        result = val
    }
    return
}

// ScaleFloat64 will scale a number from the old range
// to the new range
func ScaleFloat64(low, high, oldLow, oldHigh, value float64) float64 {
	return low + high*((value-oldLow)/(oldHigh - oldLow))
}
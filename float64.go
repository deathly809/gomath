
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
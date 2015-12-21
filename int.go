
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
package gomath

// MaxFloat32 takes in two or more integers and returns
// the maximum of them
func MaxFloat32(a float32, rem ...float32) float32 {
	res := a
	for _, v := range rem {
		if v > res {
			res = v
		}
	}
	return res
}

// MinFloat32 takes in two or more integers and returns
// the maximum of them
func MinFloat32(a float32, rem ...float32) float32 {
	res := a
	for _, v := range rem {
		if v < res {
			res = v
		}
	}
	return res
}

// AbsFloat32 takes an integer and returns its absolute
// value
func AbsFloat32(a float32) float32 {
	return MaxFloat32(a, -a)
}

// ClampFloat32 will clamp a given value between a low and
// high value inclusively.
func ClampFloat32(low, high, val float32) (result float32) {
	if val < low {
		result = low
	} else if val > high {
		result = high
	} else {
		result = val
	}
	return
}

// ScaleFloat32 will scale a number from the old range
// to the new range
func ScaleFloat32(low, high, oldLow, oldHigh, value float32) float32 {
	return low + (high-low)*((value-oldLow)/(oldHigh-oldLow))
}

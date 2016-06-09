package flops

import (
	"github.com/luxengine/lux/math"
)

const (
	// Epsilon is mathgl's constant
	Epsilon float32 = 1e-5
	// MinNormal is the smallest normal value possible.
	MinNormal = float32(1.1754943508222875e-38) // 1 / 2**(127 - 1)
	// MinValue is the smallest non zero value possible.
	MinValue = float32(math.SmallestNonzeroFloat32)
	// MaxValue is the highest value a float32 can have.
	MaxValue = float32(math.MaxFloat32)
)

var (
	// InfPos is the positive infinity value.
	InfPos = float32(math.Inf(1))
	// InfNeg is the positive infinity value.
	InfNeg = float32(math.Inf(-1))
	// NaN is a shortcut for not a number
	NaN = float32(math.NaN())
)

// Christer's constant
const (
	epsilon = 0.000001
)

// refequal returns true if the floats are approximatelly equal. this function
// is used as reference for the unwrapped equal.
func refequal(a, b float32) bool {
	return math.Abs(a-b) <= epsilon*math.Max(math.Max(1, math.Abs(a)), math.Abs(b))
}

// Eq returns true if the floats are approximatelly equal.
func Eq(a, b float32) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	}

	diff := math.Abs(a - b)
	if a*b == 0 || diff < MinNormal { // If a or b are 0 or both are extremely close to it
		return diff < Epsilon*Epsilon
	}

	// Else compare difference
	return diff/(math.Abs(a)+math.Abs(b)) < Epsilon
}

// EqChrister returns true if the floats are approximatelly equal.
func EqChrister(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			if a > b {
				if a > 1 {
					return a-b <= epsilon*a
				}
				return a-b <= epsilon
			}
			if b > 1 {
				return b-a <= epsilon*b
			}
			return b-a <= epsilon
		}
		return false
	}
	if b > 0 {
		return false
	}
	if a > b {
		if b < -1 {
			return -(b - a) <= epsilon*-b
		}
		return -(b - a) <= epsilon
	}
	if a < -1 {
		return -(a - b) <= epsilon*-a
	}
	return -(a - b) <= epsilon
}

// Ne returns true if the floats are not approximately equal
func Ne(a, b float32) bool {
	if a > 0 {
		if b > 0 {
			if a > b {
				if a > 1 {
					return a-b > epsilon*a
				}
				return a-b > epsilon
			}
			if b > 1 {
				return b-a > epsilon*b
			}
			return b-a > epsilon
		}
		return true
	}
	if b > 0 {
		return true
	}
	if a > b {
		if b < -1 {
			return -(b - a) > epsilon*-b
		}
		return -(b - a) > epsilon
	}
	if a < -1 {
		return -(a - b) > epsilon*-a
	}
	return -(a - b) > epsilon
}

// Lt returns true if a is strictly less than b. Even if a<b would return true
// they could in fact be equal.
func Lt(a, b float32) bool {
	return a < b && Ne(a, b)
}

// Le returns true if a is less than or equal to b. Even if a<b would return
// true they could in fact be equal.
func Le(a, b float32) bool {
	return a < b || Eq(a, b)
}

// Gt returns true if a is strictly greater than b. Even if a>b would return
// true they could in fact be equal.
func Gt(a, b float32) bool {
	return a > b && Ne(a, b)
}

// Ge returns true if a is greater than or equal to b. Even if a>b would return
// true they could in fact be equal.
func Ge(a, b float32) bool {
	return a > b || Eq(a, b)
}

// Ltz returns true if a is strictly less than b.zero.
func Ltz(a float32) bool {
	return a < 0 && !Z(a)
}

// Lez returns true if a is less than or equal to zero.
func Lez(a float32) bool {
	return a < 0 || Z(a)
}

// Gtz returns true if a is strictly greater than zero.
func Gtz(a float32) bool {
	return a > 0 && !Z(a)
}

// Gez returns true if a is greater than or equal to zero.
func Gez(a float32) bool {
	return a > 0 || Z(a)
}

// refz returns true if a is close to zero. This is the reference implementation
// used in testing and benchmarking.
func refz(a float32) bool {
	if math.Abs(a) <= epsilon*math.Max(1, math.Abs(a)) {
		return true
	}
	return false
}

// Z returns true if a is roughly equal to zero.
func Z(a float32) bool {
	if a > 0 {
		return a <= epsilon
	}
	return -a <= epsilon
}

// Nz returns true if a is not roughly equal to zero.
func Nz(a float32) bool {
	if a > 0 {
		return a >= epsilon
	}
	return -a >= epsilon
}

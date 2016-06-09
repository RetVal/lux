package glm

import (
	"github.com/luxengine/lux/math"
)

const (
	// Epsilon is the default epsilon for float equality
	Epsilon = 1e-10
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

// FloatEqualThreshold is a utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
//
// This differs from FloatEqual in that it lets you pass in your comparison threshold, so that you can adjust the comparison value to your specific needs
func FloatEqualThreshold(a, b, epsilon float32) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	}

	diff := math.Abs(a - b)
	if a*b == 0 || diff < MinNormal { // If a or b are 0 or both are extremely close to it
		return diff < epsilon*epsilon
	}

	// Else compare difference
	return diff/(math.Abs(a)+math.Abs(b)) < epsilon
}

// Clamp takes in a value and two thresholds. If the value is smaller than the
// low threshold, it returns the low threshold. If it's bigger than the high
// threshold it returns the high threshold. Otherwise it returns the value.
func Clamp(a, low, high float32) float32 {
	if a < low {
		return low
	} else if a > high {
		return high
	}

	return a
}

// IsClamped checks if a is clamped between low and high as if
// Clamp(a, low, high) had been called.
//
// In most cases it's probably better to just call Clamp
// without checking this since it's relatively cheap.
func IsClamped(a, low, high float32) bool {
	return a >= low && a <= high
}

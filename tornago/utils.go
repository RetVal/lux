package tornago

import (
	"github.com/luxengine/lux/math"
)

// powStandard returns x^y
func powStandard(x, y float32) float32 {
	return math.Pow(x, y)
}

// powSimple returns x^y using e^(ln(x)*y)
func powSimple(x, y float32) float32 {
	return math.Exp(math.Log(x) * y)
}

// powKartik returns x^y using Simpsons rule.
func powKartik(x, y float32) float32 {
	return math.Exp(y * (x - 1) / 6 * (1 + 8/(1+x) + 1/x))
}

// powDamping returns x^y centered at the values used for calculating
// damping^duration. It isn't guaranteed to be accurate outside the range
// damping = [0,1], duration = [0,0.4]
func powDamping(x, y float32) float32 {
	return math.Exp(y * (x - 1) / 6 * (1 + 8/(1+x) + 1/x))
}

package span

import (
	"github.com/luxengine/lux/math"
)

// Span represents an interval.
type Span struct {
	Min, Max float32
}

// NewSpan makes the span of [f, f]
func NewSpan(f float32) Span {
	return Span{f, f}
}

// Add 2 span togheter
//	[a, b] + [c, d] = [a+c, b+d]
func Add(s0, s1 Span) Span {
	return Span{s0.Min + s1.Min, s0.Max + s1.Max}
}

// Sub 2 span togheter
//	[a, b] - [c, d] = [a-c, b-d]
func Sub(s0, s1 Span) Span {
	return Span{s0.Min - s1.Min, s0.Max - s1.Max}
}

// Mul multiply this these 2 span togheter
//	[a, b] * [c, d] = [min(ac, ad, bc, bd), max(ac, ad, bc, bd)]
func Mul(s0, s1 Span) Span {
	return Span{
		math.Min(math.Min(s0.Min*s1.Min, s0.Max*s1.Max), math.Min(s0.Max*s1.Min, s0.Max*s1.Min)),
		math.Max(math.Max(s0.Min*s1.Min, s0.Max*s1.Max), math.Max(s0.Max*s1.Min, s0.Max*s1.Min)),
	}
}

// Div returns s0/s1
func Div(s0, s1 Span) Span {
	s2 := Span{1 / s1.Min, 1 / s1.Max}
	return Mul(s0, s2)
}

// Abs return the absolute of the given span.
func Abs(s Span) Span {
	return Span{math.Abs(s.Min), math.Abs(s.Max)}
}

package glm

import (
	"github.com/luxengine/lux/flops"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Parallel()

	var a float32 = 1.5
	var b float32 = 1.0 + 0.5

	if !flops.Eq(a, a) {
		t.Errorf("Float Equal fails on comparing a number with itself")
	}

	if !flops.Eq(a, b) {
		t.Errorf("Float Equal fails to compare two equivalent numbers with minimal drift")
	} else if !flops.Eq(b, a) {
		t.Errorf("Float Equal is not symmetric for some reason")
	}

	if !flops.Eq(0.0, 0.0) {
		t.Errorf("Float Equal fails to compare zero values correctly")
	}

	if flops.Eq(1.5, 1.51) {
		t.Errorf("Float Equal gives false positive on large difference")
	}

	if !flops.Eq(1.5, 1.5000001) {
		t.Errorf("Float Equal gives false negative on small difference")
	}

	if flops.Eq(1.5, 0.0) {
		t.Errorf("Float Equal gives false positive comparing with zero")
	}
}

func TestEqualThreshold(t *testing.T) {
	t.Parallel()

	// |1.0 - 1.01| < .1
	if !FloatEqualThreshold(1.0, 1.01, 1e-1) {
		t.Errorf("Thresholded equal returns negative on threshold")
	}

	// Comes out to |1.0 - 1.01| < .0001
	if FloatEqualThreshold(1.0, 1.01, 1e-3) {
		t.Errorf("Thresholded equal returns false positive on tolerant threshold")
	}
}

func TestEqualThresholdTable(t *testing.T) {
	t.Parallel()
	// http://floating-point-gui.de/errors/NearlyEqualsTest.java

	tests := []struct {
		A, B, Ep float32
		Expected bool
	}{
		{1.0, 1.01, 1e-1, true},
		{1.0, 1.01, 1e-3, false},

		// Regular large numbers
		{1000000.0, 1000001.0, 0.00001, true},
		{1000001.0, 1000000.0, 0.00001, true},
		{10000.0, 10001.0, 0.00001, false},
		{10001.0, 10000.0, 0.00001, false},

		// Negative large numbers
		{-1000000.0, -1000001.0, 0.00001, true},
		{-1000001.0, -1000000.0, 0.00001, true},
		{-10000.0, -10001.0, 0.00001, false},
		{-10001.0, -10000.0, 0.00001, false},

		// Numbers around 1
		{1.0000001, 1.0000002, 0.00001, true},
		{1.0000002, 1.0000001, 0.00001, true},
		{1.0002, 1.0001, 0.00001, false},
		{1.0001, 1.0002, 0.00001, false},

		// Numbers around -1
		{-1.000001, -1.000002, 0.00001, true},
		{-1.000002, -1.000001, 0.00001, true},
		{-1.0001, -1.0002, 0.00001, false},
		{-1.0002, -1.0001, 0.00001, false},

		// Numbers between 1 and 0
		{0.000000001000001, 0.000000001000002, 0.00001, true},
		{0.000000001000002, 0.000000001000001, 0.00001, true},
		{0.000000000001002, 0.000000000001001, 0.00001, false},
		{0.000000000001001, 0.000000000001002, 0.00001, false},

		// Numbers between -1 and 0
		{-0.000000001000001, -0.000000001000002, 0.00001, true},
		{-0.000000001000002, -0.000000001000001, 0.00001, true},
		{-0.000000000001002, -0.000000000001001, 0.00001, false},
		{-0.000000000001001, -0.000000000001002, 0.00001, false},

		// Comparisons involving zero
		{0.0, 0.0, 0.00001, true},
		{0.0, -0.0, 0.00001, true},
		{-0.0, -0.0, 0.00001, true},
		{0.00000001, 0.0, 0.00001, false},
		{0.0, 0.00000001, 0.00001, false},
		{-0.00000001, 0.0, 0.00001, false},
		{0.0, -0.00000001, 0.00001, false},

		// Comparisons involving infinities
		{InfPos, InfPos, 0.00001, true},
		{InfNeg, InfNeg, 0.00001, true},
		{InfNeg, InfPos, 0.00001, false},
		{InfPos, MaxValue, 0.00001, false},
		{InfNeg, -MaxValue, 0.00001, false},

		// Comparisons involving NaN values
		{NaN, NaN, 0.00001, false},
		{0.0, NaN, 0.00001, false},
		{NaN, 0.0, 0.00001, false},
		{-0.0, NaN, 0.00001, false},
		{NaN, -0.0, 0.00001, false},
		{NaN, InfPos, 0.00001, false},
		{InfPos, NaN, 0.00001, false},
		{NaN, InfNeg, 0.00001, false},
		{InfNeg, NaN, 0.00001, false},
		{NaN, MaxValue, 0.00001, false},
		{MaxValue, NaN, 0.00001, false},
		{NaN, -MaxValue, 0.00001, false},
		{-MaxValue, NaN, 0.00001, false},
		{NaN, MinValue, 0.00001, false},
		{MinValue, NaN, 0.00001, false},
		{NaN, -MinValue, 0.00001, false},
		{-MinValue, NaN, 0.00001, false},

		// Comparisons of numbers on opposite sides of 0
		{1.000000001, -1.0, 0.00001, false},
		{-1.0, 1.000000001, 0.00001, false},
		{-1.000000001, 1.0, 0.00001, false},
		{1.0, -1.000000001, 0.00001, false},
		{10 * MinValue, 10 * -MinValue, 0.00001, true},
		{10000 * MinValue, 10000 * -MinValue, 0.00001, true},

		// Comparisons of numbers very close to zero
		{MinValue, -MinValue, 0.00001, true},
		{-MinValue, MinValue, 0.00001, true},
		{MinValue, 0, 0.00001, true},
		{0, MinValue, 0.00001, true},
		{-MinValue, 0, 0.00001, true},
		{0, -MinValue, 0.00001, true},
		{0.000000001, -MinValue, 0.00001, false},
		{0.000000001, MinValue, 0.00001, false},
		{MinValue, 0.000000001, 0.00001, false},
		{-MinValue, 0.000000001, 0.00001, false},
	}

	for _, c := range tests {
		if r := FloatEqualThreshold(c.A, c.B, c.Ep); r != c.Expected {
			t.Errorf("FloatEqualThreshold(%v, %v, %v) != %v (got %v)", c.A, c.B, c.Ep, c.Expected, r)
		}
	}
}

func TestEqual32(t *testing.T) {
	t.Parallel()

	a := float32(1.5)
	b := float32(1.0 + .5)

	if !flops.Eq(a, a) {
		t.Errorf("Float Equal fails on comparing a number with itself")
	}

	if !flops.Eq(a, b) {
		t.Errorf("Float Equal fails to compare two equivalent numbers with minimal drift")
	} else if !flops.Eq(b, a) {
		t.Errorf("Float Equal is not symmetric for some reason")
	}

	if !flops.Eq(0.0, 0.0) {
		t.Errorf("Float Equal fails to compare zero values correctly")
	}

	if flops.Eq(1.5, 1.51) {
		t.Errorf("Float Equal gives false positive on large difference")
	}

	if flops.Eq(1.5, 0.0) {
		t.Errorf("Float Equal gives false positive comparing with zero")
	}
}

func TestClampf(t *testing.T) {
	t.Parallel()

	if !flops.Eq(Clamp(-1.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp returns incorrect value for below threshold")
	}

	if !flops.Eq(Clamp(0.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp does something weird when value is at threshold")
	}

	if !flops.Eq(Clamp(.14, 0.0, 1.0), .14) {
		t.Errorf("Clamp fails to return correct value when value is within threshold")
	}

	if !flops.Eq(Clamp(1.1, 0.0, 1.0), 1.0) {
		t.Errorf("Clamp fails to return max threshold when appropriate")
	}
}

func TestIsClamped(t *testing.T) {
	t.Parallel()

	if IsClamped(-1.0, 0.0, 1.0) {
		t.Errorf("Test below min is considered clamped")
	}

	if !IsClamped(.15, 0.0, 1.0) {
		t.Errorf("Test in threshold returns false")
	}

	if IsClamped(1.5, 0.0, 1.0) {
		t.Errorf("Test above max threshold returns false positive")
	}

	if IsClamped(1.5, 0.0, 1.0) {
		t.Errorf("Test above max threshold returns false positive")
	}
}

/* These benchmarks probably aren't very interesting, there's not really many ways to optimize the functions they're benchmarking */

func BenchmarkEqual(b *testing.B) {
	var f1 float32 = 2
	var f2 float32 = 1

	for i := 0; i < b.N; i++ {
		flops.Eq(f1, f2)
	}
}

// Here just to get a baseline of how much worse the safer equal is
func BenchmarkBuiltinEqual(b *testing.B) {
	var f1 float32 = 2
	var f2 float32 = 1

	for i := 0; i < b.N; i++ {
		_ = f1 == f2
	}
}

func BenchmarkClampf(b *testing.B) {
	var a float32 = 1.5
	var t1 float32 = 1
	var t2 float32 = 2

	for i := 0; i < b.N; i++ {
		Clamp(a, t1, t2)
	}
}

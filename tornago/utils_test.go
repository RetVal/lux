package tornago

import (
	"github.com/luxengine/lux/math"
	"math/rand"
	"testing"
)

func BenchmarkPowStandard(b *testing.B) {
	for x := 0; x < b.N; x++ {
		powStandard(0.95, 0.16)
	}
}

func BenchmarkPowSimple(b *testing.B) {
	for x := 0; x < b.N; x++ {
		powSimple(0.95, 0.16)
	}
}

func BenchmarkPowKartik(b *testing.B) {
	for x := 0; x < b.N; x++ {
		powKartik(0.95, 0.16)
	}
}

const (
	samples = 1000
)

func TestPow_Accuracy(t *testing.T) {
	var diff float32
	for i := 0; i < samples; i++ {
		x, y := rand.Float32(), rand.Float32()*0.4 //[0,1] [0,0.4]
		actual := math.Pow(x, y)

		f := powStandard(x, y)
		diff += math.Abs(actual - f)
	}
	t.Logf("average difference of Pow = %f", diff/samples)
}

func TestPowSimple_Accuracy(t *testing.T) {
	var diff float32
	for i := 0; i < samples; i++ {
		x, y := rand.Float32(), rand.Float32()*0.4 //[0,1] [0,0.4]
		actual := math.Pow(x, y)

		f := powSimple(x, y)
		diff += math.Abs(actual - f)
	}
	t.Logf("average difference of PowSimple = %f", diff/samples)
}

func TestPowKartik_Accuracy(t *testing.T) {
	var diff float32
	for i := 0; i < samples; i++ {
		x, y := rand.Float32(), rand.Float32()*0.4 //[0,1] [0,0.4]
		actual := math.Pow(x, y)

		f := powKartik(x, y)
		diff += math.Abs(actual - f)
	}
	t.Logf("average difference of PowKartik = %f", diff/samples)
}

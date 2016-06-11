package geo

import (
	"github.com/luxengine/lux/glm"
	"math/rand"
	"testing"
)

func BenchmarkIsConvexQuad(b *testing.B) {
	bench := struct {
		a, b, c, d glm.Vec3
		isconvex   bool
	}{
		a:        glm.Vec3{X: 0, Y: 0, Z: 0},
		b:        glm.Vec3{X: 0, Y: 1, Z: 0},
		c:        glm.Vec3{X: 1, Y: 1, Z: 0},
		d:        glm.Vec3{X: 1, Y: 0, Z: 0},
		isconvex: true,
	}
	for n := 0; n < b.N; n++ {
		IsConvexQuad(&bench.a, &bench.b, &bench.c, &bench.d)
	}
}

var points, dir = func() ([]glm.Vec3, glm.Vec3) {
	r := rand.New(rand.NewSource(999))
	dir := glm.Vec3{X: 1, Y: 0, Z: 0}
	points := make([]glm.Vec3, 1000)
	for n := 0; n < 1000; n++ {
		points[n] = glm.Vec3{X: r.Float32(), Y: r.Float32(), Z: r.Float32()}
	}
	return points, dir
}()

func BenchmarkExtremePointsAlongDirection1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ExtremePointsAlongDirection(&dir, points)
	}
}

func BenchmarkVariance1000(b *testing.B) {
	r := rand.New(rand.NewSource(999))
	data := make([]float32, 1000)
	for n := 0; n < 1000; n++ {
		data[n] = r.Float32()
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Variance(data)
	}
}

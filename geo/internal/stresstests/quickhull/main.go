package main

import (
	"fmt"
	"github.com/hydroflame/isitdoneyet"
	"github.com/luxengine/lux/geo"
	"github.com/luxengine/lux/glm"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("If you're running this good luck because I have no clue what I'm doing")
	const (
		iterations int     = 1000
		cloudsize  int     = 300
		cloudwidth float32 = 100
	)
	doneyet := isitdoneyet.New(os.Stdout, time.Second*3)

	points := make([]glm.Vec3, cloudsize)

	r := rand.New(rand.NewSource(1))

	var crashers int
	done := make(chan struct{})

	doneyet.Start()
	for n := 0; n < iterations; n++ {
		doneyet.Progress(float64(n) / float64(iterations))
		// randomize all the points
		for n := range points {
			points[n] = glm.Vec3{
				r.Float32()*cloudwidth*2 - cloudwidth,
				r.Float32()*cloudwidth*2 - cloudwidth,
				r.Float32()*cloudwidth*2 - cloudwidth,
			}
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					crashers++
				}
			}()
			_ = geo.Quickhull(points)
			done <- struct{}{}
		}()

		select {
		case <-time.After(time.Second * 10):
			crashers++
		case <-done:
			runtime.GC()
			continue
		}
	}

	fmt.Printf("crashers = %d\n", crashers)
}

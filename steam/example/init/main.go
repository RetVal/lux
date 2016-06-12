package main

import (
	"fmt"
	"github.com/luxengine/lux/steam"
)

func main() {
	ok := steam.Init()
	if !ok {
		fmt.Println("not ok")
	}
}

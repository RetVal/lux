package main

import (
	"bytes"
	"fmt"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/lux/ctrl"
	"github.com/luxengine/lux/ctrl/glfw"
	"runtime"
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	r := bytes.NewReader([]byte(`{
	"actionsets":[
		{
			"name": "ingame",
			"buttons":["jump", "run"]
		},
		{
			"name": "inmenu",
			"buttons":["select", "cancel"]
		}
	]
}`))

	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.PollEvents()

	window, err := glfw.CreateWindow(400, 400, "hello input", nil, nil)
	if err != nil {
		panic(err)
	}

	glfwd.Init(window)

	frt, err := ctrl.LoadCtrlFormat(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(frt)
}

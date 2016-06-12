package main

import (
	"github.com/luxengine/lux/geo"
	"github.com/luxengine/lux/glm"
	"log"
	"runtime"

	"github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/math"
	lux "github.com/luxengine/lux/render"
	"github.com/luxengine/lux/render/debug"

	"github.com/go-gl/glfw/v3.1/glfw"
)

// Default window size
var (
	WindowWidth  = 1280
	WindowHeight = 800
)

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	lux.InitGLFW()

	window := lux.CreateWindow(WindowWidth, WindowHeight, "", false)
	// handle retina problem
	vp := gl.Get.Viewport()

	// Little trick to deal with OSX retina displays
	WindowWidth, WindowHeight = int(vp[2]), int(vp[3])

	// redirect errors and warning to stdout, if we can
	debug.EnableGLDebugLogging()
	// === //

	// === asset manager testing === //
	assman := lux.NewAssetManager("../assets/", "models/", "shaders/", "textures/")
	assman.LoadModel("suzanne.obj", "monkey")
	assman.LoadModel("skydome.obj", "skydome")
	assman.LoadModel("ground.obj", "ground")
	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("skydome.png", "skydome")
	assman.LoadTexture("red.png", "red")
	assman.LoadTexture("brown.png", "brown")
	skydome := assman.Models["skydome"]
	ground := assman.Models["ground"]
	monkey := assman.Models["monkey"]
	// === //

	// === transf === //
	skydomeTransf := lux.NewTransform()
	groundTransf := lux.NewTransform()
	monkeyTransf := lux.NewTransform()
	monkeyTransf.Translate(0, 0, -5)

	gbuf, err := lux.NewGBuffer(int32(WindowWidth), int32(WindowHeight))
	if err != nil {
		log.Fatal(err)
	}

	// ===init of app specific stuff=== //
	angle := 0.0
	previousTime := glfw.GetTime()

	// ==camera== //
	var cam lux.Camera
	cameraAngle := glm.DegToRad(45)
	aspect := float32(int32(WindowWidth)) / float32(WindowHeight)
	var znear, zfar float32 = 0.1, 100.0
	cam.SetPerspective(cameraAngle, aspect, znear, zfar)
	cam.View.Ident()
	var fr geo.Frustum
	geo.FrustumFromPerspective(cameraAngle, aspect, znear, zfar, &fr)

	// post process //
	lux.InitPostProcessSystem()
	tonemap, err := lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostProcessFragmentShaderToneMapping)
	if err != nil {
		log.Fatal(err)
	}
	defer tonemap.Delete()

	// === shadow === //
	shadowfbo, err := lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	shadowfbo.SetOrtho(-10, 10, -10, 10, 0, 20)
	shadowfbo.LookAt(0, 5, 5, 0, 0, 0)

	// === lights === //
	var lamp lux.PointLight

	// ===render loop=== //
	quit := false
	for !window.ShouldClose() && !quit {
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			quit = true
		}
		// ==Update== //
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time
		angle += elapsed

		// ==shadow== //
		shadowfbo.BindForDrawing()
		shadowfbo.Render(ground, groundTransf)
		shadowfbo.Render(monkey, monkeyTransf)
		shadowfbo.Unbind()

		lamp.Move(5*math.Cos(float32(angle/2)), 5, 5*math.Sin(float32(angle/2)))
		shadowfbo.LookAt(lamp.X, lamp.Y, lamp.Z, 0, 0, 0)

		// ==Render== //
		gbuf.Bind(&cam)

		// normal rendering
		gbuf.Render(&cam, skydome, assman.Textures["skydome"], skydomeTransf)
		aabb := geo.AABB{
			Center:     glm.Vec3{0, 0, -5},
			HalfExtend: glm.Vec3{0.5, 0.5, 0.5},
		}

		if geo.TestAABBFrustum(&aabb, &fr, &cam.View) {
			gbuf.Render(&cam, monkey, assman.Textures["square"], monkeyTransf)
		}
		gbuf.Render(&cam, ground, assman.Textures["square"], groundTransf)

		// render lights
		gbuf.RenderLight(&cam, &lamp, shadowfbo.ShadowMat(), shadowfbo.ShadowMap(), 0.1, 1.0, 0.9)

		// aggregate
		gbuf.Aggregate()

		tonemap.PreRender()
		tonemap.Render(gbuf.AggregateFramebuffer.Out)
		tonemap.PostRender()

		// ==Maintenance== //
		window.SwapBuffers()
		glfw.PollEvents()
	}
	assman.Clean()
}

package main

import (
	"log"
	"runtime"

	"github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/glm"
	lux "github.com/luxengine/lux/render"
	"github.com/luxengine/lux/tornago"

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

	window := lux.CreateWindow(WindowWidth, WindowHeight, "physics", false)
	// handle retina problem
	vp := gl.Get.Viewport()

	// Little trick to deal with OSX retina displays
	WindowWidth, WindowHeight = int(vp[2]), int(vp[3])

	assman := lux.NewAssetManager("../assets/", "models/", "shaders/", "textures/")
	assman.LoadModel("ground.obj", "ground")
	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("brown.png", "brown")
	ground := assman.Models["ground"]
	// === //

	// === transf === //
	//skydomeTransf := lux.NewTransform()
	groundTransf := lux.NewTransform()
	// ============== //

	gbuf, err := lux.NewGBuffer(int32(WindowWidth), int32(WindowHeight))
	if err != nil {
		log.Fatal(err)
	}

	// ==camera== //
	var cam lux.Camera
	cam.SetPerspective(70.0, float32(WindowWidth)/float32(WindowHeight), 0.1, 100.0)
	cam.LookAtval(-5, 7, 5, 0, 2, 0, 0, 1, 0)

	// post process //
	lux.InitPostProcessSystem()
	tonemap, err := lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostProcessFragmentShaderToneMapping)
	if err != nil {
		log.Fatal(err)
	}
	defer tonemap.Delete()

	// === lights === //
	var lamp lux.PointLight
	lamp.Move(0, 5, 5)

	// === tornago === //
	w := tornago.NewWorld(&tornago.NaiveBroadphase{}, tornago.ContactResolver{})
	sphereBody := tornago.NewRigidBody()
	sphereBody.SetPosition3f(0.5, 2, 0)
	sphereBody.SetVelocity3f(0, 0, 0)
	sphereShape := tornago.NewCollisionSphere(1)
	sphereBody.SetCollisionShape(sphereShape)

	sphereBody.SetRestitution(0)
	sphereBody.SetAngularDamping(1)
	sphereBody.SetLinearDamping(1)
	sphereBody.SetFriction(1)

	boxBody2 := tornago.NewRigidBody()
	boxBody2.SetPosition3f(0.25, 5, 0)
	boxBody2.SetVelocity3f(0, 0, 0)
	boxShape := tornago.NewCollisionBox(glm.Vec3{X: 0.5, Y: 0.5, Z: 0.5})
	boxBody2.SetCollisionShape(boxShape)
	boxBody2.SetRestitution(0)
	boxBody2.SetAngularDamping(1)
	boxBody2.SetLinearDamping(1)
	boxBody2.SetFriction(1)

	b2 := tornago.NewRigidBody()
	boxShapeGround := tornago.NewCollisionBox(glm.Vec3{X: 4, Y: 0.5, Z: 4})
	b2.SetCollisionShape(boxShapeGround)
	b2.SetPosition3f(0, 0, 0)
	b2.SetMass(0)
	w.AddRigidBody(b2)
	b2.SetRestitution(0)
	b2.SetFriction(1)

	gravity := glm.Vec3{X: 0, Y: -5, Z: 0}
	gravFG := tornago.NewGravityForceGenerator(&gravity)
	w.AddForceGenerator(sphereBody, gravFG)
	w.AddForceGenerator(boxBody2, gravFG)

	w.AddRigidBody(sphereBody)
	w.AddRigidBody(boxBody2)

	boxModel := lux.NewVUNModel(boxShape.Mesh())
	sphereModel := lux.NewVUNModel(sphereShape.Mesh())

	boxTransf := lux.NewTransform()
	sphereTransf := lux.NewTransform()

	am := lux.NewAgentManager()
	am.NewAgent(func() bool {
		var m glm.Mat4
		sphereBody.OpenGLMatrix(&m)
		sphereTransf.SetMatrix((*[16]float32)(&m))
		return true
	})
	am.NewAgent(func() bool {
		var m glm.Mat4
		boxBody2.OpenGLMatrix(&m)
		boxTransf.SetMatrix((*[16]float32)(&m))
		return true
	})
	// === //

	// === shadow === //
	shadowfbo, err := lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	shadowfbo.SetOrtho(-10, 10, -10, 10, 0, 20)
	shadowfbo.LookAt(0, 5, 5, 0, 0, 0)

	// ===render loop=== //
	var quit bool
	previousTime := glfw.GetTime()
	for !window.ShouldClose() && !quit {
		// Should we quit ?
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			quit = true
		}

		// ==Update== //
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		// === tornago === //
		w.Step(float32(elapsed))

		// update all the agents (the 2 transforms)
		am.Tick()

		// ==shadow== //
		shadowfbo.BindForDrawing()
		shadowfbo.Render(ground, groundTransf)
		shadowfbo.Render(sphereModel, sphereTransf)
		shadowfbo.Render(boxModel, boxTransf)
		shadowfbo.Unbind()

		// ==Render== //
		gbuf.Bind(&cam)

		// normal rendering
		gbuf.Render(&cam, sphereModel, assman.Textures["brown"], sphereTransf)
		gbuf.Render(&cam, boxModel, assman.Textures["brown"], boxTransf)
		gbuf.Render(&cam, ground, assman.Textures["square"], groundTransf)

		// render lights
		gbuf.RenderLight(&cam, &lamp, shadowfbo.ShadowMat(), shadowfbo.ShadowMap(), 0.1, 1.0, 0.9)

		// Build image
		gbuf.Aggregate()

		// tonemapping needs to happen for things to look even remotely ok.
		tonemap.PreRender()
		tonemap.Render(gbuf.AggregateFramebuffer.Out)
		tonemap.PostRender()

		// ==Maintenance== //
		window.SwapBuffers()
		glfw.PollEvents()
	}
	assman.Clean()
}

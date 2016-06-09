package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
	lux "github.com/luxengine/lux/render"
	"github.com/luxengine/lux/render/debug"
	"github.com/luxengine/lux/tornago"

	"github.com/go-gl/glfw/v3.1/glfw"
	mgl "github.com/go-gl/mathgl/mgl32"
)

var _ = mgl.XYX
var pi = math.Pi

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
	// === ok thats pretty clean === //
	lux.InitGLFW()

	window := lux.CreateWindow(WindowWidth, WindowHeight, "", false)
	vp := gl.Get.Viewport()

	// Little trick to deal with OSX retina displays
	WindowWidth, WindowHeight = int(vp[2]), int(vp[3])

	// redirect errors and warning to stdout
	debug.EnableGLDebugLogging()
	// === //

	// === asset manager testing === //
	assman := lux.NewAssetManager("assets/", "models/", "shaders/", "textures/")
	assman.LoadModel("suzanne.obj", "monkey")
	assman.LoadModel("skydome.obj", "skydome")
	assman.LoadModel("ground.obj", "ground")
	assman.LoadModel("door.obj", "door")
	assman.LoadModel("sphere.obj", "sphere")
	assman.LoadTexture("square.png", "square")
	assman.LoadTexture("skydome.png", "skydome")
	assman.LoadTexture("red.png", "red")
	assman.LoadTexture("brown.png", "brown")
	monkey := assman.Models["monkey"]
	skydome := assman.Models["skydome"]
	ground := assman.Models["ground"]
	door := assman.Models["door"]
	sphere := assman.Models["sphere"]
	_, _, _, _ = monkey, ground, door, sphere
	// === //

	// === transf === //
	skydomeTransf := lux.NewTransform()
	sphereTransf := lux.NewTransform()
	_ = sphereTransf
	groundTransf := lux.NewTransform()
	groundTransf.Translate(0, -2, 0)
	// ============== //

	gbuf, err := lux.NewGBuffer(int32(WindowWidth), int32(WindowHeight))
	if err != nil {
		log.Fatal(err)
	}

	// ===init of app specific stuff=== //
	angle := 0.0
	previousTime := glfw.GetTime()

	// ==fps== //
	nbFrames := 0
	lastTime := glfw.GetTime()

	// ==camera== //
	var cam lux.Camera
	cam.SetPerspective(70.0, float32(int32(WindowWidth))/float32(WindowHeight), 0.1, 100.0)
	cam.LookAtval(-10, 10, 10, 0, 0, 0, 0, 1, 0)

	// post process //
	lux.InitPostProcessSystem()
	fxaa, err := lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostprocessfragmentshaderFxaa)
	if err != nil {
		log.Fatal(err)
	}
	defer fxaa.Delete()

	tonemap, err := lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), lux.PostProcessFragmentShaderToneMapping)
	if err != nil {
		log.Fatal(err)
	}
	defer tonemap.Delete()
	fxaa.SetNext(tonemap)
	woobly, err := lux.NewPostProcessFramebuffer(int32(WindowWidth), int32(WindowHeight), `#version 330
#define width 0.1

uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;

layout (location=0) out vec4 outputColor;

void main() {
	vec2 center = vec2(0.5, 0.5);
	vec2 dir = uv-center;
	float m = mod(length(dir)-time, 1);
	if(m < width){
		vec2 coord = uv;
		float coeff = m*sin(m*(3.1415/width))*0.8;
		coord.x = coord.x+dir.x*coeff;
		coord.y = coord.y+dir.y*coeff;
		outputColor = texture(tex, coord);
	}else {
		outputColor = texture(tex, uv);
	}
}
`+"\x00")
	if err != nil {
		log.Fatal(err)
	}
	defer woobly.Delete()
	//tonemap.SetNext(woobly)

	// === shadow === //
	shadowfbo, err := lux.NewShadowFBO(4096, 4096)
	if err != nil {
		log.Fatal(err)
	}
	shadowfbo.SetOrtho(-10, 10, -10, 10, 0, 20)
	shadowfbo.LookAt(0, 5, 5, 0, 0, 0)

	log.Println("starting rendering")

	// === lights === //
	var lamp lux.PointLight
	lamp.Move(0, 5, 5)

	// === tornago === //
	w := tornago.NewWorld(&tornago.NaiveBroadphase{}, tornago.ContactResolver{})
	sphereBody := tornago.NewRigidBody()
	sphereBody.SetPosition3f(0.5, 0, 0)
	sphereBody.SetVelocity3f(0, 0, 0)
	//boxShape := tornago.NewCollisionBox(glm.Vec3{0.5, 0.5, 0.5})
	boxShape := tornago.NewCollisionSphere(1)
	shape := tornago.NewCollisionSphere(0.5)
	_ = shape
	sphereBody.SetCollisionShape(boxShape)

	//sphereBody.SetAcceleration3f(0, -5, 0)
	sphereBody.SetRestitution(0)
	sphereBody.SetAngularDamping(1)
	sphereBody.SetLinearDamping(1)
	sphereBody.SetFriction(1)

	boxBody2 := tornago.NewRigidBody()
	boxBody2.SetPosition3f(0.25, 3, 0)
	boxBody2.SetVelocity3f(0, 0, 0)
	boxShape2 := tornago.NewCollisionBox(glm.Vec3{0.5, 0.5, 0.5})
	shape2 := tornago.NewCollisionSphere(0.5)
	_, _ = shape2, boxShape2
	boxBody2.SetCollisionShape(boxShape2)
	//boxBody2.SetAcceleration3f(0, -5, 0)
	boxBody2.SetRestitution(0)
	boxBody2.SetAngularDamping(1)
	boxBody2.SetLinearDamping(1)
	boxBody2.SetFriction(1)

	b2 := tornago.NewRigidBody()
	boxShapeGround := tornago.NewCollisionBox(glm.Vec3{999, 1, 999})
	b2.SetCollisionShape(boxShapeGround)
	b2.SetPosition3f(0, -2.5, 0)
	b2.SetMass(0)
	w.AddRigidBody(b2)
	b2.SetRestitution(0)
	b2.SetFriction(1)

	gravity := glm.Vec3{0, -5, 0}
	gravFG := tornago.NewGravityForceGenerator(&gravity)
	w.AddForceGenerator(sphereBody, gravFG)
	w.AddForceGenerator(boxBody2, gravFG)

	w.AddRigidBody(sphereBody)
	w.AddRigidBody(boxBody2)

	boxModel := lux.NewVUNModel(boxShape.Mesh())
	boxModel2 := lux.NewVUNModel(boxShape2.Mesh())

	boxRBTransf := lux.NewTransform()
	boxRBTransf2 := lux.NewTransform()

	am := lux.NewAgentManager()
	am.NewAgent(func() bool {
		var m glm.Mat4
		sphereBody.OpenGLMatrix(&m)
		boxRBTransf.SetMatrix((*[16]float32)(&m))
		return true
	})
	am.NewAgent(func() bool {
		var m glm.Mat4
		boxBody2.OpenGLMatrix(&m)
		boxRBTransf2.SetMatrix((*[16]float32)(&m))
		return true
	})
	// === //

	var i int
	// ===render loop=== //
	quit := false
	for !window.ShouldClose() && !quit {
		if window.GetKey(glfw.KeyEscape) == glfw.Press {
			quit = true
		}
		// ===fps=== //
		currTime := glfw.GetTime()
		nbFrames++
		if currTime-lastTime >= 1.0 {
			mspf := 1000.0 / float64(nbFrames)
			mspf = float64(int(mspf*100.0)) / 100.0
			window.SetTitle("FPS: " + fmt.Sprintf("%d", nbFrames))
			nbFrames = 0
			lastTime += 1.0
		}
		// =end fps= //

		// ==Update== //

		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time
		angle += elapsed

		// === tornago === //
		i++
		w.Step(float32(elapsed))
		am.Tick()

		// ==shadow== //
		shadowfbo.BindForDrawing()
		//shadowfbo.Render(sphere, sphereTransf)
		shadowfbo.Render(ground, groundTransf)
		shadowfbo.Render(boxModel, boxRBTransf)
		shadowfbo.Render(boxModel2, boxRBTransf2)
		shadowfbo.Unbind()

		// cam.LookAtval(float32(5*math.Cos(angle/2)), 5, float32(5*math.Sin(angle/2)), 0, 0, 0, 0, 1, 0)
		lamp.Move(5*math.Cos(float32(angle/2)), 5, 5*math.Sin(float32(angle/2)))
		shadowfbo.LookAt(lamp.X, lamp.Y, lamp.Z, 0, 0, 0)
		// lamp.Move(-5, 0, 0)

		// ==Render== //
		gbuf.Bind(&cam)

		// normal rendering
		gbuf.Render(&cam, skydome, assman.Textures["skydome"], skydomeTransf)
		// gbuf.Render(&cam, terrain, assman.Textures["square"], doorTransf)

		//gbuf.Render(&cam, sphere, assman.Textures["brown"], sphereTransf)
		gbuf.Render(&cam, boxModel, assman.Textures["brown"], boxRBTransf)
		gbuf.Render(&cam, boxModel2, assman.Textures["brown"], boxRBTransf2)
		gbuf.Render(&cam, ground, assman.Textures["square"], groundTransf)

		// render lights
		r, f, k := float32(0.1), float32(1), float32(0.9)
		gbuf.RenderLight(&cam, &lamp, shadowfbo.ShadowMat(), shadowfbo.ShadowMap(), r, f, k)

		// aggregate
		gbuf.Aggregate()

		tonemap.PreRender()
		tonemap.Render(gbuf.AggregateFramebuffer.Out)
		tonemap.PostRender()

		// ==Maintenance== //
		window.SwapBuffers()
		glfw.PollEvents()
	}

	log.Print("cleaning")
	assman.Clean()
}

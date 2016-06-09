#Install
`go get github.com/luxengine/lux`

#Lux  
Lux is a 3D game engine written almost entirely in [Go](http://golang.org/). We aim to provide our users with powerful and flexible tools to make games (and other 3D application too!).
Every lines of code in Lux is coded with the following goal in mind:
* Performance: code should be fast!
* Cross platform across all desktop operating systems.
* Support at least 95% of PC gamers. (currently 98% according to steam hardware survey)
* Flexibility: You, the programmer, should be able to change ANY part of the pipeline if you wanted to.
* Usability: If our library feel like crap to use. It probably is. Trying to make you have as much fun as possible when using the tools.


Features:  
* Basic asset loading. (Who doesn't have that :P)
* OpenGL abstraction layer to make your OpenGL code go-idiomatic.
* tornago, physic engine made from scratch for go.
* native float32 math library.
* Faster and memory friendly matrix library.
* Image postprocessing pipeline. We have some predefined shaders. eg: cel-shading, fxaa, color manipulation, etc
* Defered rendering pipeline.
* Basic shadow mapping.
* Steam wrapper.

WIP:  
* [Particle systems](https://github.com/luxengine/lux/blob/master/particlesystems.go). I've used some but never implemented any, it's actually a lot of fun.
* Stabilisation, documentation and testing of the rendering pipelines.
* Open source solution for UI (preferably html).
* Solution for testing using go framework. Those who tried will quickly realise that every test run in it's own goroutine and that `runtime.LockOsThread` and `TestMain` don't help.

Future work:
* More variety of model loading.
* More common CG techniques preimplemented, ready to use for developpers.
* Framework for game mods. (dynamic library loading/initialising)
* Network game solution.

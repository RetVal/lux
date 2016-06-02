#Install
`go get github.com/luxengine/lux`

#development
Currently on hold waiting to for algorithms in the physics engine. HOWEVER, if you want to use this lib but you're missing something. TELL ME!!, ill bring it up my priority list.

#Lux  
Lux is a 3D game engine written almost entirely in [Go](http://golang.org/). We aim to provide our users with powerfull and flexible tools to make games (and other 3D application too!).
Every lines of code in Lux is coded with the following goal in mind:
* Performance: Our code should be the fastest.!
* Cross platform across all desktop operating systems: Sorry mobile :(
* Support for the vast majority (95%+) of PC gamers: We are currently using OpenGL version 3.3 as default, in the future we would like to be able to switch between version and enable/disable features. Also support [Vulkan](https://www.khronos.org/vulkan) eventually.
* Flexibility: You the programmer should be able to change ANY part of the pipeline if you wanted to.
* Usability: If our library feel like crap to use. It probably is. We're trying to make you have as much fun as possible when using our tools. If we write something and we feel it doesn't meet a certain standard. We won't release it.


Features:  
* Basic asset loading. (Who doesn't have that :P)
* OpenGL abstraction layer [Lux GL](https://github.com/luxengine/gl)! Make your OpenGL code go-idiomatic :D
* [Wrapper](https://github.com/luxengine/gobullet) for Bullet physics engine. Make giant towers of block then throw massive, heavy balls at it and watch it fall.
* native float32 [math library](https://github.com/luxengine/math). Because `vec[0], vec[1], vec[2]` is prettier than `float64(vec[0]), float64(vec[1]), float64(vec[2])`
* Faster and memory friendly [matrix library](https://github.com/luxengine/glm)! go-gl mgl32 is good but sloooooooowwww, also it allocates a lot of memory.
* OpenCL wrapper and abstraction layer. Because sometimes we need to calculate stuff REALLY FREACKING FAST!
* Image postprocessing pipeline. We have some predefined shaders. eg: cel-shading, fxaa, color manipulation, etc
* Forward or Defered shading. Pick whichever you like best.
* Basic shadow mapping.
* Custom tailored [worker pool](https://github.com/luxengine/lux/blob/master/AssetManager.go) for 3d application. <- seriously this is pretty cool.
* [Awesomium wrapper](https://github.com/luxengine/gosomium). FYI we HATE this. The license on Awesomium is HORRIBLE and the latest Awesomium that has a C-api (something that is needed in order to make a wrapper) doesn't support css3! But if you need a quick and dirty html ui. I'ts pretty usefull.

WIP:  
* [Bullet port](https://github.com/luxengine/bullet). Because 1: I'm a bit crazy and 2: pure go stuff has so so many advantages.
* [Particle systems](https://github.com/luxengine/lux/blob/master/particlesystems.go). I've used some but never implemented any, it's actually a lot of fun.
* Stabilisation, documentation and testing of the rendering pipelines. Both defered and Forward. You shouldn't have to care or know how shadows are calculated. You just want them to look good. (But again if you wanted to switch technique or use your own. We want you to be able to).
* Steam wrapper. I'm REALLY hyped for the steam controllers.
* Open source solution for UI (preferably html).
* Solution for testing using go framework. Those who tried will quickly realise that every test run in it's own goroutine and that `runtime.LockOsThread` and `TestMain` don't help.
* Scene manager solution. This will most likely be integrated/based on the same stuff Bullet is.

Future work:
* Support for [Vulkan](https://www.khronos.org/vulkan)
* More variety of model loading.
* More common CG techniques preimplemented, ready to use for developpers.
* Framework for mods. (dynamic library loading/initialising)
* Network game solution.
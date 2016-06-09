// Package tornago is a rigid body physics engine written in pure Go.
//
// Basics
//
// First, you will need to create a world.
//  world := tornago.NewWorld(&tornago.NaiveBroadphase{}, tornago.ContactResolver{})
// The first argument is the Broadphase, which is the algorithm used to detect
// possible collisions. Test different broadphase to see which is more efficient
// for your scene. The second argument is the collision dispatcher. It's the
// algorithm that takes the set of collision for a step and resolves them. For
// now we only have 1 available dispatcher but you're free to implement your
// own.
//
// Next you need to create one or more rigid body.
//	b1 := NewRigidBody()
// rigid bodies are the bread and butter of the physics engine, every physical
// object need one. They can move and rotate in space. RigidBody also need to
// have certain properties set.
//  b1.SetMass(5) // defaults to 1
//  b1.SetLinearDamping(0.995) // defaults to 0.995, needs to be between 0 and 1 but preferably close, but not quite to 1
//  b1.SetAngularDamping(0.995)// same as LinearDamping
//  b1.SetAcceleration3f(0, -10, 0) // simulate gravity (-10 is close enough)
//  b1.SetPosition3f(0, 5, 0) // the world space position
//  qi := glm.QuatIdent() // set some orientation, the identity quaternion is no orientation, good enough
//  b1.SetOrientationQuat(&qi)
// However your rigid bodies still won't collide as they have no shape. So we'll
// need to add one.
//  b1.SetCollisionShape(tornago.NewCollisionBox(glm.Vec3{0.5, 0.5, 0.5}))
// Now you can add this shape to the world
//  world.AddRigidBody(b1)
// and voila, you're ready to step the world.
//  world.Step(1.0/60.0) // 1/60th of a second
//
// Collision groups
//
// You can optionally set a collision group to your rigid bodies. Each rigid
// body has a collision group, a bitfield indicating which group this body is
// part of, and a collision mask, a bitfield indicating which other group this
// body would like to collide with. The rule to decide if 2 object can collide
// is:
//	body0.Group & body1.Mask != 0 && body1.Group & body0.Mask != 0
// tornago defines constants for each of the individual groups possible as well
// as the all group, they are named "Group(None|[1-15]|All)". You can then
// define new constants in your app to have recognizable names. Here's an
// example for super mario bros:
//	const (
//		GroupMario   = tornago.Group1
//		GroupMonster = tornago.Group2
//		GroupPowerup = tornago.Group3
//		GroupScenery = tornago.Group4
//		MaskMario    = GroupMario | GroupMonster | GroupPowerup | GroupScenery // make your hero collide with everything.
//		MaskMonster  = GroupMario | GroupMonster | GroupScenery                // don't grab powerups now.
//		MaskPowerup  = GroupMario | GroupScenery                               // powerups don't collide with a lot.
//		MaskScenery  = GroupMario | GroupMonster | GroupPowerup | GroupScenery // scene colide with everything again.
//	)
// and then set the appropriate group/mask for each body
//	mushroom.Group, mushroom.Mask = GroupPowerup, MaskPowerup
//
// Collision callbacks
//
// every time 2 objects collide their callback funcs will be called and the
// argument is the other rigid body with which it collided, you can use
// RigibBody.Userdata to store a reference to any sort of data you could find
// usefull during collision but closure can also be a great help.
//
// Constraints
//
// constraints are a very important part of every simulation. You might need a
// certain body to never be more then x unit away from another, like a camera
// following a player. Most constraints will be either between 2 rigid bodies or
// a rigid body and the world.
//	str := tornago.NewStringToWorldConstraint(Vec3{0,10,0}, Vec3{0,0,0}, &body, 5)
// note that when we say "localPoint" we mean as point on the rigid body as if
// it was axis aligned, located at {0,0,0}, so if localPoint is {0,0,0} we mean
// the center of mass. Then simply add it to the world.
//  world.AddConstraint(str)
//
// Force generators
//
// Another type of force that you can apply on rigid bodies are force
// generators. Those are used for things that would not generate any hard
// contact, like buoyancy or springs.
//  spring := NewSpringForceGenerator(glm.Vec3{0, 0, 0}, &b2, glm.Vec3{0, 0, 0}, 0.9, 5)
//  world.AddForceGenerator(spring, &b1)
// Force generators are called every frame and apply the force they're supposed
// to.
//
// Ray tests
//
// Sometimes you want to know if your mouse click grabs an object or other
// queries that required casting a ray through the world and receiving who was
// hit. First create a ray and select a result method and call
//  world.RayTest(ray, result)
package tornago

package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

const (
	defaultLinearDamping  = 0.995
	defaultAngularDamping = 0.995
)

// RigidBody is the basic struct that represents any body in space.
type RigidBody struct {
	// Holds the inverse of the mass of the particle. It
	// is more useful to hold the inverse mass because
	// integration is simpler and because in real-time
	// simulation it is more useful to have objects with
	// infinite mass (immovable) than zero mass
	// (completely unstable in numerical simulation).
	inverseMass float32

	// Holds the linear position of the particle in
	// world space.
	position glm.Vec3

	//holds the world space orientation of this rigid body.
	orientation glm.Quat

	// Holds the linear velocity of the particle in
	// world space.
	velocity glm.Vec3

	// Holds the angular velocity of this rigid body.
	rotation glm.Vec3

	// Holds the constant acceleration that this objects
	// receives each frame (like gravity)
	acceleration glm.Vec3

	//Holds the inverse of the inertia tensor of this rigid body.
	inverseInertiaTensor glm.Mat3

	// Holds the damping for the linear movement. It's a Vec3 to help people
	// simulate certain situation, such as a 2D level or a pinball machine.
	linearDamping float32

	// Holds the damping for the angular movement. It's a Vec3 to help people
	// simulate certain situation, such as a 2D level or a pinball machine.
	angularDamping float32

	// the restitution is the amount of energy kept as kinetic energy after a
	// collision.
	restitution float32

	// friction it the roughness of the rigid body, it controls how much force
	// is applied when 2 surface are moving against one another and how much
	// force you need to go from resting to moving when resting on a surface.
	friction float32

	// collisionGroup is the filter this rigidbody has, it will only collide
	// with other rigid bodies
	//    if (this.filter & other.mask != 0) && (other.filter & this.mask != 0)
	collisionGroup, collisionMask uint16

	// userData is a field to store any kind of data you would want retrieved
	// after collision.
	userData interface{}

	// callback is called when this rigid body has a collision with another
	// non-nil rigid body
	callback func(*RigidBody)

	// the shape of the rigid body.
	// TODO(hydroflame): implement fixtures.
	shape CollisionShape

	// Derived data. These data are all derived from other data in the struct.
	// They have no setters or getters.

	//Holds a transform matrix for converting body space into world
	//space and vice versa. This can be achieved by calling the
	//getPointIn*Space functions.
	transformMatrix glm.Mat3x4

	// Holds the inverseInertiaTensor in world coordinates. This is calculated
	// from the transformMatrix.
	inverseInertiaTensorWorld glm.Mat3

	// Holds the current frame linear force acculumation.
	forceAccumulator glm.Vec3

	// Holds the current frame angular force accumulation.
	torqueAccumulator glm.Vec3

	// Holds the acceleration from the last frame.
	lastFrameAcceleration glm.Vec3
}

// NewRigidBody returns a new rigid body with some default values.
func NewRigidBody() *RigidBody {
	var b RigidBody
	b.New()
	return &b
}

// New initializes this rigid body this it's default values. Used for memory
// management.
func (b *RigidBody) New() {
	b.orientation = glm.QuatIdent()
	b.linearDamping = defaultLinearDamping
	b.angularDamping = defaultAngularDamping
	b.inverseMass = 1
	b.collisionGroup = Group(0)
	b.collisionMask = Mask(99)
}

//==============================================================================
//=============================Setters and Getters==============================
//==============================================================================

// UserData returns the current user data attached to this rigid body.
func (b *RigidBody) UserData() interface{} {
	return b.userData
}

// SetUserData sets this rigid body user data for this rigid body. Use this to
// recover the data you need later during collision callback.
func (b *RigidBody) SetUserData(data interface{}) {
	b.userData = data
}

// Callback returns the callback of ths rigid body.
func (b *RigidBody) Callback() func(*RigidBody) {
	return b.callback
}

// SetCallback sets this rigidbody callback.
func (b *RigidBody) SetCallback(f func(*RigidBody)) {
	b.callback = f
}

// SetPosition3f takes 3 float 32 and sets the position of this particle.
func (b *RigidBody) SetPosition3f(x, y, z float32) {
	b.position = glm.Vec3{X: x, Y: y, Z: z}
}

// SetPositionVec3 takes a Vec3 and sets the position of this particle.
func (b *RigidBody) SetPositionVec3(pos *glm.Vec3) {
	b.position = *pos
}

// Position return the position of this particle
func (b *RigidBody) Position() glm.Vec3 {
	return b.position
}

// SetOrientationQuat sets this rigid body orientation to this quaternion.
func (b *RigidBody) SetOrientationQuat(q *glm.Quat) {
	b.orientation = *q
}

// SetOrientation4f sets this rigid body orientation to this quaternion.
func (b *RigidBody) SetOrientation4f(w, x, y, z float32) {
	b.orientation = glm.Quat{W: w, Vec3: glm.Vec3{X: x, Y: y, Z: z}}
}

// Orientation returns the quaternion that represents this rigid body
// orientation.
func (b *RigidBody) Orientation() glm.Quat {
	return b.orientation
}

// SetVelocity3f takes 3 float 32 and sets the velocity of this particle.
func (b *RigidBody) SetVelocity3f(x, y, z float32) {
	b.velocity = glm.Vec3{X: x, Y: y, Z: z}
}

// SetVelocityVec3 takes a Vec3 and sets the velocity of this particle.
func (b *RigidBody) SetVelocityVec3(pos *glm.Vec3) {
	b.velocity = *pos
}

// Velocity returns the velocity of this particle.
func (b *RigidBody) Velocity() glm.Vec3 {
	return b.velocity
}

// Rotation returns the angular velocity of this particle.
func (b *RigidBody) Rotation() glm.Vec3 {
	return b.rotation
}

// SetAcceleration3f takes 3 float 32 and sets the acceleration of this particle.
func (b *RigidBody) SetAcceleration3f(x, y, z float32) {
	b.acceleration = glm.Vec3{X: x, Y: y, Z: z}
}

// SetAccelerationVec3 takes a Vec3 and sets the acceleration of this particle.
func (b *RigidBody) SetAccelerationVec3(pos *glm.Vec3) {
	b.acceleration = *pos
}

// Acceleration returns the velocity of this particle.
func (b *RigidBody) Acceleration() glm.Vec3 {
	return b.acceleration
}

// SetMass sets the mass of this particle. Mass must be positive, mass of 0 is
// interpreted as infinite mass (cannot move).
func (b *RigidBody) SetMass(mass float32) {
	//If mass is 0 we actually want an infinite mass.
	if mass == 0 || math.IsInf(mass, 0) {
		b.inverseMass = 0
		return
	}
	b.inverseMass = 1.0 / mass
}

// Mass returns the velocity of this particle.
func (b *RigidBody) Mass() float32 {
	if b.inverseMass == 0 {
		return 0
	}
	return 1.0 / b.inverseMass
}

// InverseMass returns 1/mass.
func (b *RigidBody) InverseMass() float32 {
	return b.inverseMass
}

// HasFiniteMass returns true if the mass of this body is finite.
func (b *RigidBody) HasFiniteMass() bool {
	return b.inverseMass != 0
}

// OpenGLMatrix is a utility function for rendering. It fills the given mat4.
func (b *RigidBody) OpenGLMatrix(m *glm.Mat4) {
	b.transformMatrix.Mat4In(m)
}

// SetLinearDamping sets the linear damping. Must be in the range [0, 1]. A
// value of 1 being no damping and being unmovable.
func (b *RigidBody) SetLinearDamping(damping float32) {
	b.linearDamping = damping
}

// LinearDamping returns the linear damping of this rigid body.
func (b *RigidBody) LinearDamping() float32 {
	return b.linearDamping
}

// SetAngularDamping sets the angular damping. Must be in the range [0, 1]. A
// value of 1 being no damping and being non-rotatable.
func (b *RigidBody) SetAngularDamping(damping float32) {
	b.angularDamping = damping
}

// AngularDamping returns the angular damping of this rigid body.
func (b *RigidBody) AngularDamping() float32 {
	return b.angularDamping
}

// SetInertiaTensor sets the inertia tensor for this rigid body.
func (b *RigidBody) SetInertiaTensor(inertiaTensor *glm.Mat3) {
	b.inverseInertiaTensor.InverseOf(inertiaTensor)
}

// InertiaTensor returns the inertia tensor of this rigid body.
func (b *RigidBody) InertiaTensor() glm.Mat3 {
	return b.inverseInertiaTensor.Inverse()
}

// InertiaTensorIn sets this matrix as the inertia tensor of this rigid body.
func (b *RigidBody) InertiaTensorIn(m *glm.Mat3) {
	m.InverseOf(&b.inverseInertiaTensor)
}

// SetCollisionShape sets the collision shape for this rigid body. It also sets
// the inertia tensor for this rigid body.
func (b *RigidBody) SetCollisionShape(shape CollisionShape) {
	b.shape = shape
	it := shape.GetInertiaTensor(b)
	b.SetInertiaTensor(&it)
}

// Restitution return the restitution of this rigid body.
func (b *RigidBody) Restitution() float32 {
	return b.restitution
}

// SetRestitution sets the restitution of this rigid body.
func (b *RigidBody) SetRestitution(restitution float32) {
	b.restitution = restitution
}

// SetRotation3f sets the rotation of this rigid body.
func (b *RigidBody) SetRotation3f(x, y, z float32) {
	b.rotation = glm.Vec3{X: x, Y: y, Z: z}
}

// SetFriction sets the friction of this rigid body.
func (b *RigidBody) SetFriction(friction float32) {
	b.friction = friction
}

// Friction returns the friction of this rigid body.
func (b *RigidBody) Friction() float32 {
	return b.friction
}

// SetGroup sets the collision group of this rigid body.
func (b *RigidBody) SetGroup(group uint16) {
	b.collisionGroup = group
}

// Group returns the collsion group of this rigid body.
func (b *RigidBody) Group() uint16 {
	return b.collisionGroup
}

// SetMask sets the collision mask of this rigid body.
func (b *RigidBody) SetMask(mask uint16) {
	b.collisionMask = mask
}

// Mask returns the collsion mask of this rigid body.
func (b *RigidBody) Mask() uint16 {
	return b.collisionMask
}

//==============================================================================
//===============================Applying forces================================
//==============================================================================

// AddTorque adds torque to this object.
func (b *RigidBody) AddTorque(torque *glm.Vec3) {
	b.torqueAccumulator.AddWith(torque)
}

// AddForce adds this force to the force accumulator.
func (b *RigidBody) AddForce(force *glm.Vec3) {
	b.forceAccumulator.AddWith(force)
}

// AddForceAtBodyPoint adds a force to the body at a particular point, this
// will generate some torque. The force is expressed in world coordinates
// but the point is in local coordinates. This method is usefull for contraints.
func (b *RigidBody) AddForceAtBodyPoint(force, point *glm.Vec3) {
	var worldPoint glm.Vec3
	b.PointInWorldCoordinates(point, &worldPoint)
	b.AddForceAtPoint(force, &worldPoint)
}

// AddForceAtPoint takes both vector in world space and adds the appropriate
// torque to the torque accumulator.
func (b *RigidBody) AddForceAtPoint(force, point *glm.Vec3) {
	pt := *point
	pt.SubWith(&b.position)
	var torque glm.Vec3
	torque.CrossOf(&pt, force)

	b.forceAccumulator.AddWith(force)
	b.torqueAccumulator.AddWith(&torque)
}

//==============================================================================
//==================================Transforms==================================
//==============================================================================

// PointInWorldCoordinates takes a point in local coordinates, transforms
// it in local coordiantes and stores the result in dst.
func (b *RigidBody) PointInWorldCoordinates(point, dst *glm.Vec3) {
	b.transformMatrix.TransformIn(point, dst)
}

// Integrate calculates the new position and orientation of this object.
func (b *RigidBody) Integrate(duration float32) {
	if b.inverseMass == 0 {
		b.calculateDerivedData()
		b.clearAccumulators()
		return
	}

	b.lastFrameAcceleration = b.acceleration
	b.lastFrameAcceleration.AddScaledVec(b.inverseMass, &b.forceAccumulator)

	var angularAcceleration glm.Vec3
	b.inverseInertiaTensorWorld.Mul3x1In(&b.torqueAccumulator, &angularAcceleration)

	//update velocity
	b.velocity.AddScaledVec(duration, &b.lastFrameAcceleration)

	//update rotation
	b.rotation.AddScaledVec(duration, &angularAcceleration)

	//drag
	ld, ad := b.linearDamping, b.angularDamping

	ld = powDamping(ld, duration)
	ad = powDamping(ad, duration)

	b.velocity.MulWith(ld)
	b.rotation.MulWith(ad)

	//update position
	b.position.AddScaledVec(duration, &b.velocity)

	//then orientation
	b.orientation.AddScaledVec(duration, &b.rotation)

	b.calculateDerivedData()
	b.clearAccumulators()
}

// clearAccumulators sets both accumulator to zero.
func (b *RigidBody) clearAccumulators() {
	b.forceAccumulator.Zero()
	b.torqueAccumulator.Zero()
}

// calculateDerivedData calculates internal data from state data. This should be
// called after the body’s state is altered directly (it is called automatically
// during integration). If you change the body’s state and then intend to
// integrate before querying any data (such as the transform matrix), then you
// can omit this step.
func (b *RigidBody) calculateDerivedData() {
	b.orientation.Normalize()
	b.transformMatrix.SetOrientationAndPos(&b.orientation, &b.position)

	{ //so this piece of code here rotates the inverse inertia tensor according
		// to the transform matrix and stores the result in the inverse inertia
		// tensor world.
		t4 := b.transformMatrix[0]*b.inverseInertiaTensor[0] + b.transformMatrix[3]*b.inverseInertiaTensor[1] + b.transformMatrix[6]*b.inverseInertiaTensor[2]
		t9 := b.transformMatrix[0]*b.inverseInertiaTensor[3] + b.transformMatrix[3]*b.inverseInertiaTensor[4] + b.transformMatrix[6]*b.inverseInertiaTensor[5]
		t14 := b.transformMatrix[0]*b.inverseInertiaTensor[6] + b.transformMatrix[3]*b.inverseInertiaTensor[7] + b.transformMatrix[6]*b.inverseInertiaTensor[8]
		t28 := b.transformMatrix[1]*b.inverseInertiaTensor[0] + b.transformMatrix[4]*b.inverseInertiaTensor[1] + b.transformMatrix[7]*b.inverseInertiaTensor[2]
		t33 := b.transformMatrix[1]*b.inverseInertiaTensor[3] + b.transformMatrix[4]*b.inverseInertiaTensor[4] + b.transformMatrix[7]*b.inverseInertiaTensor[5]
		t38 := b.transformMatrix[1]*b.inverseInertiaTensor[6] + b.transformMatrix[4]*b.inverseInertiaTensor[7] + b.transformMatrix[7]*b.inverseInertiaTensor[8]
		t52 := b.transformMatrix[2]*b.inverseInertiaTensor[0] + b.transformMatrix[5]*b.inverseInertiaTensor[1] + b.transformMatrix[8]*b.inverseInertiaTensor[2]
		t57 := b.transformMatrix[2]*b.inverseInertiaTensor[3] + b.transformMatrix[5]*b.inverseInertiaTensor[4] + b.transformMatrix[8]*b.inverseInertiaTensor[5]
		t62 := b.transformMatrix[2]*b.inverseInertiaTensor[6] + b.transformMatrix[5]*b.inverseInertiaTensor[7] + b.transformMatrix[8]*b.inverseInertiaTensor[8]

		b.inverseInertiaTensorWorld[0] = t4*b.transformMatrix[0] + t9*b.transformMatrix[3] + t14*b.transformMatrix[6]
		b.inverseInertiaTensorWorld[3] = t4*b.transformMatrix[1] + t9*b.transformMatrix[4] + t14*b.transformMatrix[7]
		b.inverseInertiaTensorWorld[6] = t4*b.transformMatrix[2] + t9*b.transformMatrix[5] + t14*b.transformMatrix[8]

		b.inverseInertiaTensorWorld[1] = t28*b.transformMatrix[0] + t33*b.transformMatrix[3] + t38*b.transformMatrix[6]
		b.inverseInertiaTensorWorld[4] = t28*b.transformMatrix[1] + t33*b.transformMatrix[4] + t38*b.transformMatrix[7]
		b.inverseInertiaTensorWorld[7] = t28*b.transformMatrix[2] + t33*b.transformMatrix[5] + t38*b.transformMatrix[8]

		b.inverseInertiaTensorWorld[2] = t52*b.transformMatrix[0] + t57*b.transformMatrix[3] + t62*b.transformMatrix[6]
		b.inverseInertiaTensorWorld[5] = t52*b.transformMatrix[1] + t57*b.transformMatrix[4] + t62*b.transformMatrix[7]
		b.inverseInertiaTensorWorld[8] = t52*b.transformMatrix[2] + t57*b.transformMatrix[5] + t62*b.transformMatrix[8]
	}
}

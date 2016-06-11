package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

const (
	angularLimit  = 0.2
	velocityLimit = 0.25
)

// contactDerivateData calculate the derivate data used to resolve vecocity and
// penetration.
type contactDerivateData struct {
	// the transform matrix from contact coordinates to world coordinates.
	contactToWorld glm.Mat3

	// the velocity of the contact point itself.
	contactVelocity glm.Vec3

	// the desired delta velocity.
	desiredDeltaVelocity float32

	// the difference between the contact point and the bodies center of mass.
	relativeContactPosition [2]glm.Vec3
}

// Contact holds the data produced by the narrowphase contact generation.
type Contact struct {
	// Holds the rigid bodies involved.
	bodies [2]*RigidBody

	// Holds the contact point in world space.
	point glm.Vec3

	// Holds the contact normal in world space.
	normal glm.Vec3

	// Holds the depth of penetration at the contact point. If both bodies are
	// specified then the contact point should be midway between the
	// inter-penetrating points.
	penetration float32

	// Holds the friction of the contact.
	friction float32

	// Holds the restitution of the contact.
	restitution float32
}

// Penetration returns the penetration depth of the 2 collision shapes.
func (c *Contact) Penetration() float32 {
	return c.penetration
}

// Friction returns the friction between the 2 collision shapes.
func (c *Contact) Friction() float32 {
	return c.friction
}

// Restitution returns the restitution of the collision between the 2 collision
// shapes.
func (c *Contact) Restitution() float32 {
	return c.restitution
}

// swapIfNeed makes sure that if the first rigid body is nil then we change
// bodies[0] and bodies[1]. This will panic if both are nil. (2 nil rigid body
// should NOT have been generated)
func (c *Contact) swapIfNeed() {
	if c.bodies[0] == nil {
		c.normal.Invert()

		c.bodies[0] = c.bodies[1]
		c.bodies[1] = nil
	}

	// remove zero mass bodies as they cant move.
	if c.bodies[1] != nil && c.bodies[1].inverseMass == 0 {
		c.bodies[1] = nil
	} else if c.bodies[0].inverseMass == 0 {
		c.bodies[0] = nil
		c.swapIfNeed()
	}
}

// calculateDerivateData fill in the given structure with all the common data we
// might need to resolve the contact.
func (c *Contact) calculateDerivateData(fill *contactDerivateData, duration float32) {
	c.swapIfNeed()

	// calculate contact to world matrix
	c.calculateContactToWorld(fill)

	// calculate relative contacts positions.
	bp := c.bodies[0].Position()
	fill.relativeContactPosition[0].SubOf(&c.point, &bp)
	if c.bodies[1] != nil {
		bp = c.bodies[1].Position()
		fill.relativeContactPosition[1].SubOf(&c.point, &bp)
	}

	// Calculate contact velocity
	fill.contactVelocity = c.calculateLocalVelocity(fill, 0, duration)
	if c.bodies[1] != nil {
		v := c.calculateLocalVelocity(fill, 1, duration)
		fill.contactVelocity.SubWith(&v)
	}

	// calculate desired delta velocity
	c.calculateDesiredDeltaVelocity(fill, duration)
}

func (c *Contact) calculateLocalVelocity(fill *contactDerivateData, i int, duration float32) glm.Vec3 {
	velocity := c.bodies[i].rotation.Cross(&fill.relativeContactPosition[i])
	velocity.AddWith(&c.bodies[i].velocity)

	// Turn the velocity into contact-coordinates.
	contactVelocity := fill.contactToWorld.Mul3x1Transpose(&velocity)

	// Calculate the amount of velocity that is due to forces without
	// reactions.
	accVelocity := c.bodies[i].lastFrameAcceleration.Mul(duration)

	// Calculate the velocity in contact-coordinates.
	accVelocity = fill.contactToWorld.Mul3x1Transpose(&accVelocity)

	// We ignore any component of acceleration in the contact normal
	// direction, we are only interested in planar acceleration
	accVelocity.X = 0

	// Add the planar velocities - if there's enough friction they will
	// be removed during velocity resolution
	contactVelocity.AddWith(&accVelocity)

	// And return it
	return contactVelocity
}

func (c *Contact) calculateContactToWorld(fill *contactDerivateData) {
	var contactTangent [2]glm.Vec3

	if math.Abs(c.normal.X) > math.Abs(c.normal.Y) {
		s := 1 / math.Sqrt(c.normal.Z*c.normal.Z+c.normal.X*c.normal.X)

		contactTangent[0].X = c.normal.Z * s
		contactTangent[0].Y = 0
		contactTangent[0].Z = -c.normal.X * s

		contactTangent[1].X = c.normal.Y * contactTangent[0].X
		contactTangent[1].Y = c.normal.Z*contactTangent[0].X - c.normal.X*contactTangent[0].Z
		contactTangent[1].Z = -c.normal.Y * contactTangent[0].X
	} else {
		s := 1 / math.Sqrt(c.normal.Z*c.normal.Z+c.normal.Y*c.normal.Y)

		contactTangent[0].X = 0
		contactTangent[0].Y = -c.normal.Z * s
		contactTangent[0].Z = c.normal.Y * s

		contactTangent[1].X = c.normal.Y*contactTangent[0].Z - c.normal.Z*contactTangent[0].Y
		contactTangent[1].Y = -c.normal.X * contactTangent[0].Z
		contactTangent[1].Z = c.normal.X * contactTangent[0].Y
	}

	fill.contactToWorld = glm.Mat3FromCols(&c.normal, &contactTangent[0], &contactTangent[1])
}

func (c *Contact) calculateDesiredDeltaVelocity(fill *contactDerivateData, duration float32) {
	// Calculate the acceleration induced velocity accumulated this frame
	var velocityFromAcc float32

	velocityFromAcc = c.bodies[0].lastFrameAcceleration.Dot(&c.normal) * duration
	if c.bodies[1] != nil {
		velocityFromAcc = c.bodies[1].lastFrameAcceleration.Dot(&c.normal) * duration
	}

	// If the velocity is very slow, limit the restitution
	if math.Abs(fill.contactVelocity.X) < velocityLimit {
		c.restitution = 0
	}

	// Combine the bounce velocity with the removed
	// acceleration velocity.
	fill.desiredDeltaVelocity = -fill.contactVelocity.X - c.restitution*(fill.contactVelocity.X-velocityFromAcc)
}

// resolveVelocity resolves the given contact velocity and penetration.
// velocityChange and rotationChange are actually return values used by the
// contact resolver, it helps keep things on the stack.
func (c *Contact) resolveVelocity(data *contactDerivateData, velocityChange, rotationChange *[2]glm.Vec3) {
	// if the bodies are separating then we want them to continue separating.
	if data.desiredDeltaVelocity <= 0 {
		return
	}

	inverseInertiaTensors := [2]glm.Mat3{c.bodies[0].inverseInertiaTensorWorld, {}}
	if c.bodies[1] != nil {
		inverseInertiaTensors[1] = c.bodies[1].inverseInertiaTensorWorld
	}

	var impulseContact glm.Vec3

	if c.Friction() == 0 {
		impulseContact = c.calculateFrictionlessImpulse(data, &inverseInertiaTensors)
	} else {
		impulseContact = c.calculateFrictionImpulse(data, &inverseInertiaTensors)
	}
	impulse := data.contactToWorld.Mul3x1(&impulseContact)

	impulsiveTorque := data.relativeContactPosition[0].Cross(&impulse)
	rotationChange[0] = inverseInertiaTensors[0].Mul3x1(&impulsiveTorque)
	velocityChange[0] = glm.Vec3{}
	velocityChange[0].AddScaledVec(c.bodies[0].InverseMass(), &impulse)

	c.bodies[0].velocity.AddWith(&velocityChange[0])
	c.bodies[0].rotation.AddWith(&rotationChange[0])

	if c.bodies[1] != nil {
		impulsiveTorque = impulse.Cross(&data.relativeContactPosition[1])
		rotationChange[1] = inverseInertiaTensors[1].Mul3x1(&impulsiveTorque)
		velocityChange[1] = glm.Vec3{}
		velocityChange[1].AddScaledVec(-c.bodies[1].InverseMass(), &impulse)

		c.bodies[1].velocity.AddWith(&velocityChange[1])
		c.bodies[1].rotation.AddWith(&rotationChange[1])

	}
}

// resolvePenetration resolves the penetration of the coliding bodies.
// linearChange and angularChange are actually return values used by the contact
// resolver, it helps keep things on the stack.
func (c *Contact) resolvePenetration(data *contactDerivateData, linearChange, angularChange *[2]glm.Vec3) {
	var angularMove [2]float32
	var linearMove [2]float32

	var totalInertia float32
	var linearInertia [2]float32
	var angularInertia [2]float32

	// We need to work out the inertia of each object in the direction
	// of the contact normal, due to angular inertia only.
	for i := 0; i < 2; i++ {
		if c.bodies[i] != nil {
			// Use the same procedure as for calculating frictionless
			// velocity change to work out the angular inertia.
			angularInertiaWorld := data.relativeContactPosition[i].Cross(&c.normal)
			angularInertiaWorld = c.bodies[i].inverseInertiaTensorWorld.Mul3x1(&angularInertiaWorld)
			angularInertiaWorld = angularInertiaWorld.Cross(&data.relativeContactPosition[i])
			angularInertia[i] = angularInertiaWorld.Dot(&c.normal)

			linearInertia[i] = c.bodies[i].inverseMass

			// Keep track of the total inertia from all components
			totalInertia += linearInertia[i] + angularInertia[i]

			// We break the loop here so that the totalInertia value is
			// completely calculated (by both iterations) before
			// continuing.
		}
	}

	inverseInertia := 1 / totalInertia

	// Loop through again calculating and applying the changes
	for i := 0; i < 2; i++ {
		if c.bodies[i] != nil {
			// The linear and angular movements required are in proportion to
			// the two inverse inertias.
			sign := float32(1)
			if i == 1 {
				sign = -1
			}
			angularMove[i] = sign * c.penetration * angularInertia[i] * inverseInertia
			linearMove[i] = sign * c.penetration * linearInertia[i] * inverseInertia

			// To avoid angular projections that are too great (when mass is large
			// but inertia tensor is small) limit the angular move.
			projection := data.relativeContactPosition[i]
			projection.AddScaledVec(-data.relativeContactPosition[i].Dot(&c.normal), &c.normal)

			// Use the small angle approximation for the sine of the angle (i.e.
			// the magnitude would be sine(angularLimit) * projection.magnitude
			// but we approximate sine(angularLimit) to angularLimit).
			maxMagnitude := angularLimit * projection.Len()

			if angularMove[i] < -maxMagnitude {
				totalMove := angularMove[i] + linearMove[i]
				angularMove[i] = -maxMagnitude
				linearMove[i] = totalMove - angularMove[i]
			} else if angularMove[i] > maxMagnitude {
				totalMove := angularMove[i] + linearMove[i]
				angularMove[i] = maxMagnitude
				linearMove[i] = totalMove - angularMove[i]
			}

			if angularMove[i] == 0 {
				angularChange[i] = glm.Vec3{}
			} else {
				targetAngularDirection := data.relativeContactPosition[i].Cross(&c.normal)
				angularChange[i] = c.bodies[i].inverseInertiaTensorWorld.Mul3x1(&targetAngularDirection)
				angularChange[i].MulWith(angularMove[i] / angularInertia[i] / 10)
			}

			// Velocity change is easier - it is just the linear movement
			// along the contact normal.
			linearChange[i] = c.normal.Mul(linearMove[i])

			// Now we can start to apply the values we've calculated.
			// Apply the linear movement

			c.bodies[i].position.AddScaledVec(linearMove[i], &c.normal)

			// And the change in orientation
			c.bodies[i].orientation.AddScaledVec(1, &angularChange[i])
			c.bodies[i].orientation.Normalize()

			// We need to calculate the derived data for any body that is
			// asleep, so that the changes are reflected in the object's
			// data. Otherwise the resolution will not change the position
			// of the object, and the next collision detection round will
			// have the same penetration.
			c.bodies[i].calculateDerivedData()
		}
	}
}

// calculateFrictionlessImpulse calculates the impulse needed assuming there's
// no friction.
func (c *Contact) calculateFrictionlessImpulse(data *contactDerivateData, inverseInertiaTensors *[2]glm.Mat3) glm.Vec3 {
	// build a vector that shows the change in velocity in World Space for
	// a unit impulse in the direction of the contact normal
	deltaVelWorld := data.relativeContactPosition[0].Cross(&c.normal)
	deltaVelWorld = inverseInertiaTensors[0].Mul3x1(&deltaVelWorld)
	deltaVelWorld = deltaVelWorld.Cross(&data.relativeContactPosition[0])

	deltaVelocity := deltaVelWorld.Dot(&c.normal)

	deltaVelocity += c.bodies[0].InverseMass()

	if c.bodies[1] != nil {
		deltaVelWorld = data.relativeContactPosition[1].Cross(&c.normal)
		deltaVelWorld = inverseInertiaTensors[1].Mul3x1(&deltaVelWorld)
		deltaVelWorld = deltaVelWorld.Cross(&data.relativeContactPosition[1])

		deltaVelocity += deltaVelWorld.Dot(&c.normal)

		deltaVelocity += c.bodies[1].InverseMass()
	}

	return glm.Vec3{X: data.desiredDeltaVelocity / deltaVelocity, Y: 0, Z: 0}
}

// calculateFrictionImpulse calculates the impulse needed to resolve this
// contact with friction.
func (c *Contact) calculateFrictionImpulse(data *contactDerivateData, inverseInertiaTensors *[2]glm.Mat3) glm.Vec3 {
	inverseMass := c.bodies[0].InverseMass()

	var impulseToTorque glm.Mat3
	setSkewSymmetric(&impulseToTorque, &data.relativeContactPosition[0])

	deltaVelWorld := impulseToTorque.Mul3(&inverseInertiaTensors[0])
	deltaVelWorld = deltaVelWorld.Mul3(&impulseToTorque)
	deltaVelWorld.MulWith(-1)

	if c.bodies[1] != nil {
		setSkewSymmetric(&impulseToTorque, &data.relativeContactPosition[1])

		deltaVelWorld2 := impulseToTorque.Mul3(&inverseInertiaTensors[1])
		deltaVelWorld2 = deltaVelWorld2.Mul3(&impulseToTorque)
		deltaVelWorld2.MulWith(-1)

		deltaVelWorld.AddWith(&deltaVelWorld2)

		inverseMass += c.bodies[1].InverseMass()
	}

	deltaVelocity := data.contactToWorld.Transposed()
	deltaVelocity = deltaVelocity.Mul3(&deltaVelWorld)
	deltaVelocity = deltaVelocity.Mul3(&data.contactToWorld)

	deltaVelocity[0] += inverseMass
	deltaVelocity[4] += inverseMass
	deltaVelocity[8] += inverseMass

	impulseMatrix := deltaVelocity.Inverse()

	velKill := glm.Vec3{
		X: data.desiredDeltaVelocity,
		Y: -data.contactVelocity.Y,
		Z: -data.contactVelocity.Z,
	}

	impulseContact := impulseMatrix.Mul3x1(&velKill)

	planarImpulse := math.Sqrt(impulseContact.Y*impulseContact.Y + impulseContact.Z*impulseContact.Z)

	// the abs here is needed in case of -0.0
	if planarImpulse > math.Abs(impulseContact.X*c.Friction()) {
		// Use dynamic friction
		impulseContact.Y /= planarImpulse
		impulseContact.Z /= planarImpulse

		impulseContact.X = deltaVelocity[0] +
			deltaVelocity[3]*c.Friction()*impulseContact.Y +
			deltaVelocity[6]*c.Friction()*impulseContact.Z

		impulseContact.X = data.desiredDeltaVelocity / impulseContact.X
		impulseContact.Y *= c.Friction() * impulseContact.X
		impulseContact.Z *= c.Friction() * impulseContact.X
	}

	return impulseContact
}

func setSkewSymmetric(m *glm.Mat3, v *glm.Vec3) {
	m[0], m[3], m[6] = 0.0, -v.Z, v.Y
	m[1], m[4], m[7] = v.Z, 0.0, -v.X
	m[2], m[5], m[8] = -v.Y, v.X, 0.0
}

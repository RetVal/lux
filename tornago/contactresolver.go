package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/math"
)

const (
	// how many iterations of position resolution we're willing to do max.
	positionIterations = 50
)

// makeOrthonormal takes x as a non-zero vector and sets y and z as
func makeOrthonormal(x, y, z *glm.Vec3) {
	if math.Abs(x.Y) > math.Abs(x.X) {
		y.X, y.Y, y.Z = 1, 0, 0
	} else {
		y.X, y.Y, y.Z = 0, 1, 0
	}
	z.CrossOf(x, y)
	y.CrossOf(z, x)
}

// ContactResolver is the default Dispatcher.
type ContactResolver struct{}

// ResolveContacts takes a contact slice and resolves them using the fancy core
// algorithm.
func (c ContactResolver) ResolveContacts(contacts []Contact, duration float32) {
	// calculate derivate data
	derivateData := make([]contactDerivateData, len(contacts))
	for n := 0; n < len(contacts); n++ {
		contacts[n].calculateDerivateData(&derivateData[n], duration)
		for i, b := range contacts[n].bodies {
			if b != nil && b.callback != nil {
				b.callback(contacts[n].bodies[1-i])
			}
		}
	}

	c.adjustPositions(contacts, derivateData)
	c.adjustVelocities(contacts, derivateData, duration)

}

// adjustVelocities resolves as many velocities as it can.
// TODO(hydroflame): make a list that assumes (with reason) temporal coherence
// to speed up the algorithm. Also this can probably be done in parallel with
// the other adjust.
func (c ContactResolver) adjustVelocities(contacts []Contact, derivateData []contactDerivateData, duration float32) {
	// reserve some memory for keeping track of velocity change.
	var velocityChange, rotationChange [2]glm.Vec3

	// resolve penetration first
	for i := 0; i < positionIterations; i++ {

		// keep track of worse so far
		worst := len(contacts)
		var deltaVelocity float32
		for i := range contacts {
			if desiredDeltaVelocity := derivateData[i].desiredDeltaVelocity; deltaVelocity < desiredDeltaVelocity {
				worst = i
				deltaVelocity = desiredDeltaVelocity
			}
		}

		// if we didn't find any bad contact then we're done.
		if worst == len(contacts) {
			break
		}

		// resolve velocity.
		contacts[worst].resolveVelocity(&derivateData[worst], &velocityChange, &rotationChange)

		var deltaVel glm.Vec3
		for i := range contacts {
			for b := 0; b < 2; b++ {
				if contacts[i].bodies[b] != nil {
					for d := 0; d < 2; d++ {
						if contacts[i].bodies[b] == contacts[worst].bodies[d] {
							deltaVel = velocityChange[d]
							tmp := rotationChange[d].Cross(&derivateData[i].relativeContactPosition[b])
							deltaVel.AddWith(&tmp)

							tmp = derivateData[i].contactToWorld.Mul3x1Transpose(&deltaVel)
							if b == 0 {
								derivateData[i].contactVelocity.AddScaledVec(1, &tmp)
							} else {
								derivateData[i].contactVelocity.AddScaledVec(-1, &tmp)
							}

							contacts[i].calculateDesiredDeltaVelocity(&derivateData[i], duration)
						}
					}
				}
			}
		}
	}
}

// adjustPositions resolves as many positions as it can.
// TODO(hydroflame): make a list that assumes (with reason) temporal coherence
// to speed up the algorithm. Also this can probably be done in parallel with
// the other adjust.
func (c ContactResolver) adjustPositions(contacts []Contact, derivateData []contactDerivateData) {
	// reserve some memory for keeping track of position change.
	var linearChange, angularChange [2]glm.Vec3

	// resolve penetration first
	for i := 0; i < positionIterations; i++ {

		// keep track of worse so far
		worst := len(contacts)
		var penetration float32
		for i := range contacts {
			if contactPenetration := contacts[i].penetration; penetration < contactPenetration {
				worst = i
				penetration = contactPenetration
			}
		}

		// if we didn't find any bad contact then we're done.
		if worst == len(contacts) {
			break
		}

		contacts[worst].resolvePenetration(&derivateData[worst], &linearChange, &angularChange)

		var deltaPosition glm.Vec3
		for i := range contacts {
			for b := 0; b < 2; b++ {
				if contacts[i].bodies[b] != nil {
					for d := 0; d < 2; d++ {
						if contacts[i].bodies[b] == contacts[worst].bodies[d] {
							deltaPosition = angularChange[d].Cross(&derivateData[i].relativeContactPosition[b])
							deltaPosition.AddWith(&linearChange[d])

							contacts[i].penetration += deltaPosition.Dot(&contacts[i].normal) * ((float32(b) * 2) - 1)
						}
					}
				}
			}
		}
	}
}

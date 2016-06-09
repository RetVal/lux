package tornago

import (
	"github.com/luxengine/lux/glm"
	"github.com/tbogdala/cubez"
	"github.com/tbogdala/cubez/math"
	"math/rand"
	"testing"
	"time"
)

/*
&{inverseMass:0 position:[0 -2.5 0] orientation:{W:1 V:[0 0 0]} velocity:[0 0 0] rotation:[0 0 0] acceleration:[0 0 0] transformMatrix:[1 0 0 0 1 0 0 0 1 0 -2.5 0] inverseInertiaTensor:[0.19607842 0 0 0 0.104166664 0 0 0 0.19607842] linearDamping:0.995 angularDamping:0.995 restitution:0.1 userData:<nil> shape:0xc82014e1a0 inverseInertiaTensorWorld:[0.19607842 0 0 0 0.104166664 0 0 0 0.19607842] forceAccumulator:[0 0 0] torqueAccumulator:[0 0 0] lastFrameAcceleration:[0 0 0]}
&{inverseMass:1 position:[0 -1.1275783 0] orientation:{W:1 V:[0 0 0]} velocity:[0 -4.000025 0] rotation:[0 0 0] acceleration:[0 -5 0] transformMatrix:[1 0 0 0 1 0 0 0 1 0 -1.1275783 0] inverseInertiaTensor:[6.666666 0 0 0 6.666666 0 0 0 6.666666] linearDamping:0.7 angularDamping:0.7 restitution:0 userData:<nil> shape:0xc82014e180 inverseInertiaTensorWorld:[6.666666 0 0 0 6.666666 0 0 0 6.666666] forceAccumulator:[0 0 0] torqueAccumulator:[0 0 0] lastFrameAcceleration:[0 -5 0]}
friction 0
point [0.5 -1.6275783 0.5]
normal [-0 -1 -0]
penetration 0.12757826
*/

func mtoglmVec3(v math.Vector3) glm.Vec3 {
	return glm.Vec3{float32(v[0]), float32(v[1]), float32(v[2])}
}

/*
	c 0.5, &[-0.40031807694627436 -0.8273430759774565 -0.3940163345626647], [0 1 0]
	t1 0.413672



	c 1, {-0.400318, -0.827343, -0.394016}, [0 1 0]
	t1 0.827343
*/

func TestCubez(t *testing.T) {
	t.Skip("not caring about cubez rn")
	var tg, cb struct {
		positionChange      [2]glm.Vec3
		angularChange       [2]glm.Vec3
		velocityChange      [2]glm.Vec3
		rotationChange      [2]glm.Vec3
		FrictionImpulse     [2]glm.Vec3
		FrictionlessImpulse [2]glm.Vec3
		contact             struct {
			point       glm.Vec3
			normal      glm.Vec3
			penetration float32
		}
	}

	seed := time.Now().UnixNano()
	seed = 1448499743473438458
	rand.Seed(seed)
	t.Log(seed)

	for x := 0; x < 2; x++ {
		ori1 := glm.Quat{rand.Float32()*2 - 1, glm.Vec3{rand.Float32()*2 - 1, rand.Float32()*2 - 1, rand.Float32()*2 - 1}}
		ori1.Normalize()

		if x == 0 {
			continue
		}

		const (
			timestep = 0.16
		)

		{
			c := struct {
				s1 CollisionBox
				s2 CollisionBox
			}{
				s1: CollisionBox{
					body: &RigidBody{
						inverseMass:          1,
						orientation:          ori1,
						position:             glm.Vec3{0, 1, 0},
						inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
						linearDamping:        1,
						angularDamping:       1,
					},
					halfSize: glm.Vec3{0.5, 0.5, 0.5},
				},
				s2: CollisionBox{
					body: &RigidBody{
						inverseMass:          0,
						orientation:          glm.QuatIdent(),
						inverseInertiaTensor: cuboidInertiaTensor(0, 4, 0.5, 4),
						linearDamping:        1,
						angularDamping:       1,
					},
					halfSize: glm.Vec3{4, 0.5, 4},
				},
			}

			iit1 := cuboidInertiaTensor(1, 0.5, 0.5, 0.5)
			c.s1.body.SetInertiaTensor(&iit1)
			iit2 := cuboidInertiaTensor(0, 4, 0.5, 4)
			c.s2.body.SetInertiaTensor(&iit2)

			c.s1.body.calculateDerivedData()
			c.s2.body.calculateDerivedData()

			t.Errorf("tg iit %v", c.s1.body.inverseInertiaTensor)
			t.Errorf("tg iitw %v", c.s1.body.inverseInertiaTensorWorld)

			var data contactDerivateData

			contacts := make([]Contact, 1)
			boxAndBox(&c.s1, &c.s2, contacts)

			contact := contacts[0]
			contact.friction = 0
			contact.restitution = 0

			tg.contact.point = contact.point
			tg.contact.normal = contact.normal
			tg.contact.penetration = contact.penetration

			contact.calculateDerivateData(&data, timestep)

			contact.resolvePenetration(&data, &tg.positionChange, &tg.angularChange)
			contact.resolveVelocity(&data, &tg.velocityChange, &tg.rotationChange)
		}
		{

			box1 := cubez.NewRigidBody()
			box1.Position = math.Vector3{0, 1, 0}
			box1.SetMass(1)
			box1.Orientation = math.Quat{math.Real(ori1.W), math.Real(ori1.X), math.Real(ori1.Y), math.Real(ori1.Z)}
			box1.LinearDamping = 1
			box1.AngularDamping = 1

			box1Collider := cubez.NewCollisionCube(box1, math.Vector3{0.5, 0.5, 0.5})

			var box1Inertia math.Matrix3
			box1Inertia.SetBlockInertiaTensor(&box1Collider.HalfSize, 1.0)
			box1Collider.Body.SetInertiaTensor(&box1Inertia)

			box1Collider.Body.CalculateDerivedData()
			box1Collider.CalculateDerivedData()

			//
			box2 := cubez.NewRigidBody()
			box2.Position = math.Vector3{0, 0, 0}
			box2.SetInfiniteMass()
			box2.LinearDamping = 1
			box2.AngularDamping = 1
			box2Collider := cubez.NewCollisionCube(box2, math.Vector3{4, 0.5, 4})

			var box2Inertia math.Matrix3
			box2Inertia.SetBlockInertiaTensor(&box2Collider.HalfSize, 0)
			box2Collider.Body.SetInertiaTensor(&box2Inertia)

			box2Collider.Body.CalculateDerivedData()
			box2Collider.CalculateDerivedData()

			t.Errorf("cb iit %v", box1.InverseInertiaTensor)
			t.Errorf("cb iitw %v", box1.GetInverseInertiaTensorWorld())

			var colls []*cubez.Contact
			_, colls = box1Collider.CheckAgainstCube(box2Collider, colls)
			contact := *colls[0]

			cb.contact.point = mtoglmVec3(contact.ContactPoint)
			cb.contact.normal = mtoglmVec3(contact.ContactNormal)
			cb.contact.penetration = float32(contact.Penetration)

			contact.Restitution = 0
			contact.Friction = 0

			contact.CalculateInternals(timestep)

			cpc, cac := contact.ApplyPositionChange(contact.Penetration)
			cvc, crc := contact.ApplyVelocityChange()

			cb.positionChange[0] = mtoglmVec3(cpc[0])
			cb.positionChange[1] = mtoglmVec3(cpc[1])

			cb.angularChange[0] = mtoglmVec3(cac[0])
			cb.angularChange[1] = mtoglmVec3(cac[1])

			cb.velocityChange[0] = mtoglmVec3(cvc[0])
			cb.velocityChange[1] = mtoglmVec3(cvc[1])

			cb.rotationChange[0] = mtoglmVec3(crc[0])
			cb.rotationChange[1] = mtoglmVec3(crc[1])
		}

		const threshold = 1

		if !tg.contact.point.EqualThreshold(&cb.contact.point, threshold) ||
			!tg.contact.normal.EqualThreshold(&cb.contact.normal, threshold) ||
			!glm.FloatEqualThreshold(tg.contact.penetration, cb.contact.penetration, threshold) {
			t.Error("point")
			t.Errorf("\ttg %v", tg.contact.point)
			t.Errorf("\tcb %v", cb.contact.point)

			t.Error("normal")
			t.Errorf("\ttg %v", tg.contact.normal)
			t.Errorf("\tcb %v", cb.contact.normal)

			t.Error("penetration")
			t.Errorf("\ttg %f", tg.contact.penetration)
			t.Errorf("\tcb %f", cb.contact.penetration)
		}

		if !tg.positionChange[0].EqualThreshold(&cb.positionChange[0], threshold) || !tg.positionChange[1].EqualThreshold(&cb.positionChange[1], threshold) ||
			!tg.angularChange[0].EqualThreshold(&cb.angularChange[0], threshold) || !tg.angularChange[1].EqualThreshold(&cb.angularChange[1], threshold) ||
			!tg.velocityChange[0].EqualThreshold(&cb.velocityChange[0], threshold) || !tg.velocityChange[1].EqualThreshold(&cb.velocityChange[1], threshold) ||
			!tg.rotationChange[0].EqualThreshold(&cb.rotationChange[0], threshold) || !tg.rotationChange[1].EqualThreshold(&cb.rotationChange[1], threshold) {
			t.Log("")
			t.Error("position change")
			t.Errorf("\ttg pc %v", tg.positionChange)
			t.Errorf("\tcb pc %v", cb.positionChange)

			t.Error("angular change")
			t.Errorf("\ttg pc %v", tg.angularChange)
			t.Errorf("\tcb pc %v", cb.angularChange)

			t.Error("velocity change")
			t.Errorf("\ttg pc %v", tg.velocityChange)
			t.Errorf("\tcb pc %v", cb.velocityChange)

			t.Error("rotation change")
			t.Errorf("\ttg pc %v", tg.rotationChange)
			t.Errorf("\tcb pc %v", cb.rotationChange)
		}

	}
}

func TestCubez_BoxBox(t *testing.T) {
	t.Skip("not caring about cubez rn")
	seed := time.Now().UnixNano()
	seed = 1448483239209393901
	rand.Seed(seed)

	for x := 0; x < 100; x++ {
		ori := glm.Quat{rand.Float32()*2 - 1, glm.Vec3{rand.Float32()*2 - 1, rand.Float32()*2 - 1, rand.Float32()*2 - 1}}
		ori.Normalize()
		ori = glm.Quat{0.8398189, glm.Vec3{-0.08513479, -0.14371091, 0.5165302}}

		var cbNormal, tgNormal glm.Vec3
		var cbPoint, tgPoint glm.Vec3
		var cbPen, tgPen float32
		{ // cubez
			box1 := cubez.NewRigidBody()
			box1.Position = math.Vector3{0, 1, 0}
			box1.Orientation = math.Quat{math.Real(ori.W), math.Real(ori.X), math.Real(ori.Y), math.Real(ori.Z)}
			box1.SetMass(1)
			box1Collider := cubez.NewCollisionCube(box1, math.Vector3{0.5, 0.5, 0.5})

			var box1Inertia math.Matrix3
			box1Inertia.SetBlockInertiaTensor(&box1Collider.HalfSize, 1.0)
			box1Collider.Body.SetInertiaTensor(&box1Inertia)

			box1Collider.Body.CalculateDerivedData()
			box1Collider.CalculateDerivedData()

			//
			box2 := cubez.NewRigidBody()
			box2.Position = math.Vector3{0, 0, 0}
			box2.SetMass(1)
			box2Collider := cubez.NewCollisionCube(box2, math.Vector3{0.5, 0.5, 0.5})

			var box2Inertia math.Matrix3
			box2Inertia.SetBlockInertiaTensor(&box2Collider.HalfSize, 0)
			box2Collider.Body.SetInertiaTensor(&box2Inertia)

			box2Collider.Body.CalculateDerivedData()
			box2Collider.CalculateDerivedData()

			var colls []*cubez.Contact
			_, colls = box1Collider.CheckAgainstCube(box2Collider, colls)
			contact := *colls[0]
			cbNormal = glm.Vec3{float32(contact.ContactNormal[0]), float32(contact.ContactNormal[1]), float32(contact.ContactNormal[2])}
			cbPoint = glm.Vec3{float32(contact.ContactPoint[0]), float32(contact.ContactPoint[1]), float32(contact.ContactPoint[2])}
			cbPen = float32(contact.Penetration)
		}
		{ // tornago
			c := struct {
				s1 CollisionBox
				s2 CollisionBox
			}{
				s1: CollisionBox{
					body: &RigidBody{
						inverseMass: 1,
						orientation: ori,
						position:    glm.Vec3{0, 1, 0},
						//velocity:              glm.Vec3{0, -4.000025, 0},
						//lastFrameAcceleration: glm.Vec3{0, -5, 0},
						inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
						linearDamping:        1,
						angularDamping:       1,
						restitution:          0,
					},
					halfSize: glm.Vec3{0.5, 0.5, 0.5},
				},
				s2: CollisionBox{
					body: &RigidBody{
						inverseMass:          0,
						orientation:          glm.QuatIdent(),
						position:             glm.Vec3{0, 0, 0},
						velocity:             glm.Vec3{0, 0, 0},
						inverseInertiaTensor: cuboidInertiaTensor(1, 0.5, 0.5, 0.5),
						linearDamping:        1,
						angularDamping:       1,
						restitution:          0,
					},
					halfSize: glm.Vec3{0.5, 0.5, 0.5},
				},
			}

			iit := cuboidInertiaTensor(1, 0.5, 0.5, 0.5)
			c.s1.body.SetInertiaTensor(&iit)
			c.s2.body.SetInertiaTensor(&iit)

			c.s1.body.calculateDerivedData()
			c.s2.body.calculateDerivedData()

			contacts := make([]Contact, 1)
			boxAndBox(&c.s1, &c.s2, contacts)

			tgNormal = contacts[0].normal
			tgPoint = contacts[0].point
			tgPen = contacts[0].penetration
		}

		if !cbPoint.EqualThreshold(&tgPoint, 1e-2) || !cbNormal.EqualThreshold(&tgNormal, 1e-2) || !glm.FloatEqualThreshold(cbPen, tgPen, 1e-2) {
			t.Errorf("iteration: %d", x)
			t.Errorf("seed: %d", seed)
			t.Errorf("ori:  %v", ori)
			t.Error("\tpoint")
			t.Errorf("\t\tcubez:   %v", cbPoint)
			t.Errorf("\t\ttornago: %v", tgPoint)

			t.Error("\tnormal")
			t.Errorf("\t\tcubez:   %v", cbNormal)
			t.Errorf("\t\ttornago: %v", tgNormal)

			t.Error("\tpenetration")
			t.Errorf("\t\tcubez:   %v", cbPen)
			t.Errorf("\t\ttornago: %v", tgPen)
			return
		}
	}

}

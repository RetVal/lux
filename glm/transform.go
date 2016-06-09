package glm

// Transform is a utility type used to aggregate transformations. Transform
// concatenation, like matrix multiplication, is not commutative.
type Transform struct {
	// The position of the transform.
	Position Vec3
	// The orientation of the transform.
	Orientation Quat
	// LocalToWorld transforms local points to world points.
	LocalToWorld Mat4
	// WorldToLocal transforms world points to local points.
	WorldToLocal Mat4
}

// String return a string representation of this transform.
func (t *Transform) String() string { return t.LocalToWorld.String() }

// Iden sets this transform to the identity transform.
func (t *Transform) Iden() {
	t.Position.Zero()
	t.Orientation.Iden()
}

// MoveBy moves this object by the given position.
func (t *Transform) MoveBy(move *Vec3) {
	t.Position.AddWith(move)
}

// Rotate rotates in-place this transform.
func (t *Transform) Rotate() {}

// LookAt sets the rotation of this transform to look at a specific point.
func (t *Transform) LookAt(target, up *Vec3) {
	t.Orientation = QuatLookAtV(&t.Position, target, up)
}

// CalculateInternals calculates the internal LocalToWorld/WorldToLocal
// matrices. Call this once between the time you change position/orientation and
// you use localtoworld/worldtolocal per frame.
func (t *Transform) CalculateInternals() {
	w, x, y, z := t.Orientation.W, t.Orientation.X, t.Orientation.Y, t.Orientation.Z
	t.LocalToWorld = Mat4{
		1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0,
		2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*w*x, 0,
		2*x*z + 2*w*y, 2*y*z - 2*w*x, 1 - 2*x*x - 2*y*y, 0,
		t.Position.X, t.Position.Y, t.Position.Z, 1,
	}
	t.WorldToLocal = Mat4{
		t.LocalToWorld[0], t.LocalToWorld[4], t.LocalToWorld[8], 0,
		t.LocalToWorld[1], t.LocalToWorld[5], t.LocalToWorld[9], 0,
		t.LocalToWorld[2], t.LocalToWorld[6], t.LocalToWorld[10], 0,
		-(t.LocalToWorld[0]*t.Position.X + t.LocalToWorld[1]*t.Position.Y + t.LocalToWorld[2]*t.Position.Z),
		-(t.LocalToWorld[4]*t.Position.X + t.LocalToWorld[5]*t.Position.Y + t.LocalToWorld[6]*t.Position.Z),
		-(t.LocalToWorld[8]*t.Position.X + t.LocalToWorld[9]*t.Position.Y + t.LocalToWorld[10]*t.Position.Z),
		1,
	}
}

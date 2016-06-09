package glm

import (
	"testing"
)

func TestTransform(t *testing.T) {
	var tr Transform
	tr.Iden()

	tr.MoveBy(&Vec3{1, 0, 0})

	q := Quat{1, Vec3{2, 3, 4}}
	q.Normalize()
	tr.Orientation = q

	tr.CalculateInternals()
	t.Log(tr.Position)
	t.Log(tr.Orientation)
	t.Logf("\n%s", tr.LocalToWorld.String())
	t.Logf("\n%s", tr.WorldToLocal.String())

}

package tornago

import (
	"testing"
)

func TestInsert(t *testing.T) {
	br := NaiveBroadphase{}
	var b RigidBody
	var b2 RigidBody
	br.Insert(&b, nil)
	br.Insert(&b2, nil)
	br.Insert(&b2, nil)
	if len(br.objects) != 2 {
		t.Errorf("len(br.objects) = %d, want 2", len(br.objects))
	}
}

func TestRemove(t *testing.T) {
	br := NaiveBroadphase{}
	var b RigidBody
	var b2 RigidBody
	br.Insert(&b, nil)
	br.Insert(&b2, nil)
	br.Insert(&b2, nil)
	br.Remove(&b)
	br.Remove(&b2)
	if len(br.objects) != 0 {
		t.Errorf("len(br.objects) = %d, want 0", len(br.objects))
	}
}

package tornago

import (
	"testing"
)

func TestSap3_Remove(t *testing.T) {
	sap := SAP3{}
	var b RigidBody
	var v BoundingSphere
	v.radius = 3
	sap.Insert(&b, &v)
	sap.Remove(&b, &v)
	if len(sap.axisList[0]) != 0 || len(sap.axisList[1]) != 0 || len(sap.axisList[2]) != 0 {
		t.Errorf("not zero length, %d,%d,%d", len(sap.axisList[0]), len(sap.axisList[1]), len(sap.axisList[2]))
		t.Errorf("%v, %v, %v", sap.axisList[0], sap.axisList[1], sap.axisList[2])
	}
}

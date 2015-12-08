package lux

import (
	"github.com/luxengine/glm"
)

// SceneGraph is a graph for a scene.
type SceneGraph struct {
	root *Node
}

// Update updates the matrix of the entire tree.
func Update(float32) { panic("scene tree not implemented") }

// Traverse traverses the scene and culls object that don't intersect with the
// camera.
func Traverse(*Camera, func(*Node)) { panic("scene tree not implemented") }

// Node is a single item in a scene tree. It may just contain more children or
// it may be a model.
type Node struct {
	Transform
	parent   *Node
	children []Node

	needUpdate bool
}

func (n *Node) propagateNeedUpdate() {

	// the idea here is that we only need to propagate the changes once
	if n.needUpdate {
		return
	}

	n.needUpdate = true
	if n.parent != nil {
		n.parent.propagateNeedUpdate()
	}
}

// AttachChildren attaches a children to this node
func (n *Node) AttachChildren() { panic("scene tree not implemented") }

// Translate add the translation (x,y,z) to the current transform.
func (n *Node) Translate(x, y, z float32) {
	n.propagateNeedUpdate()
	n.Transform.Translate(x, y, z)
}

// SetTranslate reset this transform to represent only the translation transform given by (x,y,z).
func (n *Node) SetTranslate(x, y, z float32) {
	n.propagateNeedUpdate()
	n.Transform.SetTranslate(x, y, z)
}

// QuatRotate add the rotation represented by this (angle,quat) to the current transform.
func (n *Node) QuatRotate(angle float32, axis *glm.Vec3) {
	n.propagateNeedUpdate()
	n.Transform.QuatRotate(angle, axis)
}

// SetQuatRotate will reset this transform to represent the rotation represented by this (angle,quat).
func (n *Node) SetQuatRotate(angle float32, axis *glm.Vec3) {
	n.propagateNeedUpdate()
	n.Transform.SetQuatRotate(angle, axis)
}

// Scale add a scaling operation to the currently stored transform.
// I do not allow non-uniform scaling to prevent ending up with matrices without an inverse.
func (n *Node) Scale(amount float32) {
	n.propagateNeedUpdate()
	n.Transform.Scale(amount)
}

// SetScale reset this transform to represent only the scaling transform of `amount`
// I do not allow non-uniform scaling to prevent ending up with matrices without an inverse.
func (n *Node) SetScale(amount float32) {
	n.propagateNeedUpdate()
	n.Transform.SetScale(amount)
}

// Iden set this transform to the identity 4x4 matrix
func (n *Node) Iden() {
	n.propagateNeedUpdate()
	n.Transform.Iden()
}

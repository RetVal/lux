package tornago

import (
	"github.com/luxengine/lux/glm"
)

// quadTreeNode is a node in a quad tree
type quadTreeNode struct {
	position glm.Vec3
	children [4]*quadTreeNode
}

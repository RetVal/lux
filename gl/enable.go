package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Enable is an alis for glEnable(param).
func Enable(param uint32) {
	gl.Enable(param)
}

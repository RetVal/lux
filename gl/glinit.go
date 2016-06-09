package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// Init calls gl.Init()
func Init() error {
	return gl.Init()
}

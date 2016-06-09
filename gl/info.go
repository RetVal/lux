package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

// GetVendor returns the current OpenGL vendor.
func GetVendor() string {
	return gl.GoStr(gl.GetString(gl.VENDOR))
}

// GetVersion returns the current OpenGL version.
func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

// GetRenderer returns the OpenGL renderer.
func GetRenderer() string {
	return gl.GoStr(gl.GetString(gl.RENDERER))
}

// GetExtensions returns all loaded extension.
func GetExtensions() []string {
	var numExtensions int32
	gl.GetIntegerv(gl.NUM_EXTENSIONS, &numExtensions)
	extensions := make([]string, 0, numExtensions)
	for i := int32(0); i < numExtensions; i++ {
		extensions = append(extensions, gl.GoStr(gl.GetStringi(gl.EXTENSIONS, uint32(i))))
	}
	return extensions
}

// IsExtensionAvailable returns true if the given extension is available. False
// if it's not.
func IsExtensionAvailable(extension string) bool {
	exts := GetExtensions()
	for _, ext := range exts {
		if ext == extension {
			return true
		}
	}
	return false
}

package debug

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/luxengine/lux"
	"log"
	"unsafe"
)

func glfwErrorCallback(err glfw.ErrorCode, desc string) {
	log.Printf("GLFW error %v: %v\n", err, desc)
}

func glDebugCallback(source uint32, gltype uint32, id uint32, severity uint32, length int32, message string, userParam unsafe.Pointer) {
	if shouldPrint(severity) {
		log.Printf("GL:%s %s %s: %s\n", debugseverity(severity), debugsource(source), debugtype(gltype), message)
	}
}

// EnableGLDebugLogging enables debug print message (to stdout). Will not crash
// if your machine does not support GL_ARB_debug_output. Only call this once.
func EnableGLDebugLogging() {
	if _, ok := lux.Extensions["GL_ARB_debug_output"]; ok {
		log.Print("debugging enabled")
		gl.Enable(gl.DEBUG_OUTPUT_SYNCHRONOUS_ARB)
		gl.DebugMessageCallbackARB(gl.DebugProc(glDebugCallback), nil)
	}
}

//should really be more modifiable
func shouldPrint(severity uint32) bool {
	return severity != gl.DEBUG_SEVERITY_NOTIFICATION
}

func debugtype(gltype uint32) string {
	switch gltype {
	default:
		return "???"
	case gl.DEBUG_TYPE_DEPRECATED_BEHAVIOR:
		return "DEPRECATED_BEHAVIOR"
	case gl.DEBUG_TYPE_ERROR:
		return "ERROR"
	case gl.DEBUG_TYPE_MARKER:
		return "MARKER"
	case gl.DEBUG_TYPE_OTHER:
		return "OTHER"
	case gl.DEBUG_TYPE_PERFORMANCE:
		return "PERFORMANCE"
	case gl.DEBUG_TYPE_POP_GROUP:
		return "POP_GROUP"
	case gl.DEBUG_TYPE_PORTABILITY:
		return "PORTABILITY"
	case gl.DEBUG_TYPE_PUSH_GROUP:
		return "PUSH_GROUP"
	case gl.DEBUG_TYPE_UNDEFINED_BEHAVIOR:
		return "UNDEFINED_BEHAVIOR"
	}
}

func debugsource(source uint32) string {
	switch source {
	default:
		return "???"
	case gl.DEBUG_SOURCE_API:
		return "API"
	case gl.DEBUG_SOURCE_APPLICATION:
		return "APPLICATION"
	case gl.DEBUG_SOURCE_OTHER:
		return "OTHER"
	case gl.DEBUG_SOURCE_SHADER_COMPILER:
		return "SHADER_COMPILER"
	case gl.DEBUG_SOURCE_THIRD_PARTY:
		return "THIRD_PARTY"
	case gl.DEBUG_SOURCE_WINDOW_SYSTEM:
		return "WINDOW_SYSTEM"
	}
}

func debugseverity(severity uint32) string {
	switch severity {
	default:
		return "???"
	case gl.DEBUG_SEVERITY_HIGH:
		return "HIGH"
	case gl.DEBUG_SEVERITY_LOW:
		return "LOW"
	case gl.DEBUG_SEVERITY_MEDIUM:
		return "MEDIUM"
	case gl.DEBUG_SEVERITY_NOTIFICATION:
		return "NOTIFICATION"
	}
}

/*
	DEBUG_SEVERITY_NOTIFICATION                                = 0x826B
	DEBUG_SEVERITY_HIGH                                        = 0x9146
	DEBUG_SEVERITY_MEDIUM                                      = 0x9147
	DEBUG_SEVERITY_LOW                                         = 0x9148
*/

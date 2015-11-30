package control

// Driver is a bridge between some input method (Xinput, steam, GLFW) and the
// control package.
type Driver interface {
	Poll()
}

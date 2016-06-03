package lux

import (
	"github.com/luxengine/gl"
)

//Light is the interface that group all the common light operations.
type Light interface {
	SetColor(float32, float32, float32)
	Upload(gl.UniformLocation)
	CastsShadow(bool)
}

//PointLight is a simple point light with colors.
type PointLight struct {
	X, Y, Z, R, G, B float32
}

//Move sets the point light position to (x,y,z).
func (pl *PointLight) Move(x, y, z float32) {
	pl.X = x
	pl.Y = y
	pl.Z = z
}

//SetColor sets this point light color.
func (pl *PointLight) SetColor(r, g, b float32) {
	pl.R = r
	pl.G = g
	pl.B = b
}

//Upload is a placeholder, dont call it will panic.
func (pl *PointLight) Upload(u gl.UniformLocation) {
	panic("unimplemented")
}

//CastsShadow is a placeholder, dont call it will panic.
func (pl *PointLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

//SpotLight are similar to PointLight except they have an angle
type SpotLight struct {
	X, Y, Z, Tx, Ty, Tz, A, R, G, B float32
}

//Move sets the spotlight position to (x,y,z).
func (sl *SpotLight) Move(x, y, z float32) {
	sl.X = x
	sl.Y = y
	sl.Z = z
}

//SetColor sets this spotlight color.
func (sl *SpotLight) SetColor(r, g, b float32) {
	sl.R = r
	sl.G = g
	sl.B = b
}

//Upload is a placeholder, dont call it will panic.
func (sl *SpotLight) Upload(u gl.UniformLocation) {
	panic("unimplemented")
}

//CastsShadow is a placeholder, dont call it will panic.
func (sl *SpotLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

//SetAngle sets the max angle that this spotlight will illuminate.
func (sl *SpotLight) SetAngle(a float32) {
	sl.A = a
}

//DirectionalLight represent sources like the sun. Actually pretty much only the sun.
type DirectionalLight struct {
	Dx, Dy, Dz, R, G, B float32
}

//SetColor sets this spotlight color.
func (dl *DirectionalLight) SetColor(r, g, b float32) {
	dl.R = r
	dl.G = g
	dl.B = b
}

//Upload is a placeholder, dont call it will panic.
func (dl *DirectionalLight) Upload(u gl.UniformLocation) {
	panic("unimplemented")
}

//CastsShadow is a placeholder, dont call it will panic.
func (dl *DirectionalLight) CastsShadow(cast bool) {
	panic("unimplemented")
}

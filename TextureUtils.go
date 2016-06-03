package lux

import (
	"fmt" //error
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
	"image"
	"image/draw"
	"os"
)

//LoadPng tries to load a png file from hard drive and upload it to the GPU.
func LoadPng(file string) (gl2.Texture2D, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, nil
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, nil
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	texture := gl2.GenTexture2D()
	gl.ActiveTexture(gl2.TEXTURE0)
	texture.Bind()
	defer texture.Unbind()
	texture.MinFilter(gl2.LINEAR)
	texture.MagFilter(gl2.LINEAR)
	texture.WrapS(gl2.CLAMP_TO_EDGE)
	texture.WrapT(gl2.CLAMP_TO_EDGE)
	texture.TexImage2D(0, gl2.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl2.RGBA, gl2.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	return texture, nil
}

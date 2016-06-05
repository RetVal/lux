package lux

import (
	"errors"
	"github.com/luxengine/gl"
	"image"
	"image/color"
	"image/png"
	"os"
	"unsafe"
)

// errors that can be returned by SaveTexture2D
var (
	ErrUnsupportedTextureFormat = errors.New("unsupported texture format")
)

// SaveTexture2D take a Texture2D and a filename and saves it as a png image.
func SaveTexture2D(t gl.Texture2D, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	lasttex := gl.Texture2D(gl.Get.TextureBinding2D())
	defer lasttex.Bind()

	t.Bind()
	width, height := int(t.Width(0)), int(t.Height(0))
	nrgba := image.NewRGBA(image.Rect(0, 0, width, height))

	D(width, height)
	var pixels []byte

	internalformat := t.InternalFormat(0)
	switch internalformat {
	case gl.RGBA8:
		pixels = make([]byte, width*height*4)
		t.ReadPixels(0, 0, int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&pixels[0]))
		for x := 0; x < len(pixels); x += 4 {
			nrgba.SetRGBA((x/4)%width, height-(x/4)/width, color.RGBA{pixels[x+0], pixels[x+1], pixels[x+2], 255})
		}
		png.Encode(file, nrgba)
	default:
		return ErrUnsupportedTextureFormat
	}
	return nil
}

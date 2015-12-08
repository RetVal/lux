package lux

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
	"image"
	"image/color"
	"image/png"
	"os"
	"unsafe"
)

//SaveTexture2D take a Texture2D and a filename and saves it as a png image.
func SaveTexture2D(t gl2.Texture2D, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		return err
	}
	lasttex := GetCurrentTexture2D()
	defer lasttex.Bind()

	t.Bind()
	width, height := int(t.Width(0)), int(t.Height(0))
	nrgba := image.NewRGBA(image.Rect(0, 0, width, height))

	D(width, height)
	var pixels []byte

	internalformat := t.InternalFormat(0)
	if internalformat == gl.RGBA8 {
		pixels = make([]byte, width*height*4)
		t.ReadPixels(0, 0, int32(width), int32(height), gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(&pixels[0]))
		for x := 0; x < len(pixels); x += 4 {
			nrgba.SetRGBA((x/4)%width, height-(x/4)/width, color.RGBA{pixels[x], pixels[x+1], pixels[x+2], 255})
		}
	} else {
		return fmt.Errorf("unsupported texture type")
	}
	png.Encode(file, nrgba)
	return nil
}

//GetCurrentTexture2D returns the currently bound texture2D, or 0 is none is bound
func GetCurrentTexture2D() gl2.Texture2D {
	whichID := int32(0)
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &whichID)
	return gl2.Texture2D(whichID)
}

/*
GLint width,height,internalFormat;
glBindTexture(GL_TEXTURE_2D, your_texture_id);
glGetTexLevelParameteriv(GL_TEXTURE_2D, 0, GL_TEXTURE_COMPONENTS, &internalFormat); // get internal format type of GL texture
glGetTexLevelParameteriv(GL_TEXTURE_2D, 0, GL_TEXTURE_WIDTH, &width); // get width of GL texture
glGetTexLevelParameteriv(GL_TEXTURE_2D, 0, GL_TEXTURE_HEIGHT, &height); // get height of GL texture

// GL_TEXTURE_COMPONENTS and GL_INTERNAL_FORMAT are the same.
// just work with RGB8 and RGBA8

GLint numBytes = 0;
switch(internalFormat) // determine what type GL texture has...
{
case GL_RGB:
numBytes = width * height * 3;
break;
case GL_RGBA:
numBytes = width * height * 4;
break;
default: // unsupported type (or you can put some code to support more formats if you need)
break;
}

if(numBytes)
{
unsigned char *pixels = (unsigned char*)malloc(numBytes); // allocate image data into RAM
glGetTexImage(GL_TEXTURE_2D, 0, internalFormat, GL_UNSIGNED_BYTE, pixels);
{
  // TODO with pixels
}
free(pixels); // when you don't need 'pixels' anymore clean a memory page to avoid memory leak.
}
*/

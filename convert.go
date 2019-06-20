package polygonize

import (
	"image"
	// We need jpeg and png processing to read files from the drive
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// ToRGBA converts an arbitrary image.Image object to image.RGBA.
func ToRGBA(img image.Image) *image.RGBA {
	// TODO: Check wether there's a quicker way to do this, at the
	// moment it seems to be a bit dirty.
	size := img.Bounds().Size()
	newImg := image.NewRGBA(img.Bounds())
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			newImg.Set(x, y, img.At(x, y))
		}
	}
	return newImg
}

// FromFile is a function for the users convenience. It reads either
// a jpeg or a png file and returns an image.RGBA object which can
// then be used with this libarary.
func FromFile(path string) (*image.RGBA, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return ToRGBA(src), nil
}

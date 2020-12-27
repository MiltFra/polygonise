package polygonise

import (
	"errors"
	"fmt"
	"image"
	"path/filepath"

	"image/jpeg"
	"image/png"
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
	defer f.Close()
	if err != nil {
		return nil, errors.New("Failed to open the given file: " + err.Error())
	}
	src, _, err := image.Decode(f)
	if err != nil {
		return nil, errors.New("Failed to decode the given image: " + err.Error())
	}
	return ToRGBA(src), nil
}

// ToFile inverts the FromFile function. That is, it saves a given image.RGBA
// object to the disk into a given path.
func ToFile(path string, img *image.RGBA) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return errors.New("Failed to write to the given file: " + err.Error())
	}
	switch filepath.Ext(path) {
	case ".png":
		err = png.Encode(f, img)
		if err != nil {
			return errors.New("Failed to encode png to file: " + err.Error())
		}
		break
	case ".jpg", ".jpeg":
		err = jpeg.Encode(f, img, nil)
		if err != nil {
			return errors.New("Failed to encode jpeg to file: " + err.Error())
		}
		break
	default:
		fmt.Println(filepath.Ext(path))
		return errors.New("This filetype is not supported")
	}
	return nil
}

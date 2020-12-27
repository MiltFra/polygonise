package polygonise

import (
	"image"
	"image/color"
	"testing"
)

func TestToRGBA(t *testing.T) {
	bounds := image.Rect(0, 0, 1000, 1000)
	img := image.NewRGBA(bounds)
	for i := 0; i < 1000000; i++ {
		img.SetRGBA(i/1000, i%1000, color.RGBA{uint8((i / 1000) / 4), uint8((i % 1000) / 4), 0, 0})
	}
	img2 := ToRGBA(img)
	for i := 0; i < 1000000; i++ {
		col := color.RGBA{uint8((i / 1000) / 4), uint8((i % 1000) / 4), 0, 0}
		if img2.At(i/1000, i%1000) != col {
			t.Fail()
		}
	}
}

func TestFromFile(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	_ = ToRGBA(img)
}

func TestToFile(t *testing.T) {
	bounds := image.Rect(0, 0, 1000, 1000)
	img := image.NewRGBA(bounds)
	for i := 0; i < 1000000; i++ {
		img.SetRGBA(i/1000, i%1000, color.RGBA{uint8((i / 1000) / 4), uint8((i % 1000) / 4), 0, 255})
	}
	err := ToFile("teeest/tofile_test.png", img)
	if err == nil {
		t.Fail()
	}
	err = ToFile("test/tofile_test.png", img)
	if err != nil {
		t.Fail()
	}
	err = ToFile("test/tofile_test.jpg", img)
	if err != nil {
		t.Fail()
	}
}

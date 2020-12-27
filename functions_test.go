package polygonise

import (
	"testing"
)

func TestGet(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	f, err := NewGreyFilter(100, true)
	p := Get(img, f, 100)
	if len(p) != 1 || len(p[0]) < 4 {
		t.Fail()
	}
	p[0] = Flatten(p[0])
	if len(p[0]) < 4 || len(p[0])/2 > 8 {
		t.Fail()
	}
}

func TestFlatten(t *testing.T) {
	p := []int{0, 0, 0, 5, 3, 5, 5, 5, 1, 1}
	p = Flatten(p)
	if len(p) != 6 {
		t.Fail()
	}
}

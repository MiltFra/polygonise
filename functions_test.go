package polygonize

import "testing"

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

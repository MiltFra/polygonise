package polygonize

import "testing"

func TestGreyFilter(t *testing.T) {
	f, err := NewGreyFilter(100, false)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
	f, err = NewGreyFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
}

func TestGreyFilterByImage(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	f, err := NewGreyFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if !f.Filter(img.RGBAAt(1850, 63)) ||
		!f.Filter(img.RGBAAt(1850, 64)) {
		t.Fail()
	}
	if f.Filter(img.RGBAAt(0, 0)) {
		t.Fail()
	}
}

func TestGreenFilter(t *testing.T) {
	f, err := NewGreenFilter(100, false)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
	f, err = NewGreenFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
}
func TestGreenFilterByImage(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	f, err := NewGreenFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if !f.Filter(img.RGBAAt(1850, 63)) ||
		!f.Filter(img.RGBAAt(1850, 64)) {
		t.Fail()
	}
	if f.Filter(img.RGBAAt(0, 0)) {
		t.Fail()
	}
}

func TestBlueFilter(t *testing.T) {
	f, err := NewBlueFilter(100, false)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
	f, err = NewBlueFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
}

func TestBlueFilterByImage(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	f, err := NewBlueFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if !f.Filter(img.RGBAAt(1850, 63)) ||
		!f.Filter(img.RGBAAt(1850, 64)) {
		t.Fail()
	}
	if f.Filter(img.RGBAAt(0, 0)) {
		t.Fail()
	}
}

func TestRedFilter(t *testing.T) {
	f, err := NewRedFilter(100, false)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
	f, err = NewRedFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if f.Filter(f.FalseValue()) || !f.Filter(f.TrueValue()) {
		t.Fail()
	}
}
func TestRedFilterByImage(t *testing.T) {
	img, err := FromFile("test/rectangle_test.png")
	if err != nil {
		t.Fail()
	}
	f, err := NewRedFilter(100, true)
	if err != nil {
		t.Fail()
	}
	if !f.Filter(img.RGBAAt(1850, 63)) ||
		!f.Filter(img.RGBAAt(1850, 64)) {
		t.Fail()
	}
	if f.Filter(img.RGBAAt(0, 0)) {
		t.Fail()
	}
}

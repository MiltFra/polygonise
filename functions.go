package polygonize

import (
	"image"

	// We need jpeg & png procesing
	_ "image/jpeg"
	_ "image/png"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Get returns all the polygons that are available based on
// the given filter and the given step size.
//
// Note: A higher
// step size reduces the count of corners per polygon but increases
// the chance that the resulting polygons overlap.
func Get(img *image.RGBA, f Filter, step int) [][]int {
	img = copyImg(img)
	res := make([][]int, 0)
	p := nextPolygon(img, f, step)
	for p != nil {
		res = append(res, p)
		removePolygon(img, p[0], p[1], f, step)
		p = nextPolygon(img, f, step)
	}
	return res
}

// GetNext returns the next polygon in this image and another image
// with the pixels of this polygon removed. The return value is nil
// iff there are no true pixels for this filter in this image.
func GetNext(img *image.RGBA, f Filter, step int) ([]int, *image.RGBA) {
	img = copyImg(img)
	res := nextPolygon(img, f, step)
	removePolygon(img, res[0], res[1], f, step)
	return res, img
}

// Flatten removes every corner from a given polygon that is exactly on
// the line of the two adjacent corners.
func Flatten(p []int) []int {
	newP := make([]int, 0, len(p))
	var r float64
	var x, y, pX, pY, nX, nY int
	l := len(p)
	for i := 2; i < l+2; i += 2 {
		x, y = p[i%l], p[(i+1)%l]
		pX, pY = p[(i-2)%l], p[(i-1)%l]
		nX, nY = p[(i+2)%l], p[(i+3)%l]
		if x == pX && x == nX && ((y < pY) != (y < nY)) {
			continue
		}
		r = float64(nX-pX) / float64(x-pX)
		if int(r*float64(y-pY)) == nY-pY {
			continue
		}
		newP = append(newP, x, y)
	}
	return newP
}

// copyImg creates a new RGBA image identical to the argument and returns
// a pointer to the location.
func copyImg(img *image.RGBA) *image.RGBA {
	new := image.NewRGBA(img.Bounds())
	s := new.Bounds().Size()
	for x := 0; x < s.X; x++ {
		for y := 0; y < s.Y; y++ {
			new.SetRGBA(x, y, img.RGBAAt(x, y))
		}
	}
	return new
}

func nextPolygon(img *image.RGBA, f Filter, step int) []int {
	// Getting the first pixel by traversing the entire image.
	// That means if the polygon is not removed after calling this function
	// it will deterministically find the same polygon again.
	size := img.Bounds().Size()
	startX, startY := -1, -1
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			if f.Filter(img.RGBAAt(x, y)) {
				startX, startY = x, y
				break
			}
		}
	}
	// if there's no true pixel, there can't be any polygon either
	if startX == -1 {
		return nil
	}
	p := make([]int, 0)
	x, y := startX, startY
	dir := 0
	var duplicate bool
	for {
		// in every step we progress a certain amount of pixels
		// until there's a duplicate
		for i := 0; i < step; i++ {
			x, y, dir = nextPixel(img, f, x, y, dir)
			for i := 0; i < len(p); i += 2 {
				if x == p[i] && y == p[i+1] {
					duplicate = true
					break
				}
			}
		}
		// if we found a duplicate, we're done
		if duplicate {
			break
		}
		// add the pixels to the polygon
		p = append(p, x, y)
	}
	return p
}

var neighs = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

// nextPixel returns a pixel and the direction the last pixel is in. It traverses the edge
// of a filtered polygon in an image in clockwise order.
func nextPixel(img *image.RGBA, f Filter, x, y, dir int) (newX, newY, newDir int) {
	seenFalse := true
	for i := 0; i < 8; i++ {
		newX, newY = x+neighs[(i+dir)%8][0], y+neighs[(i+dir)%8][1]
		if !seenFalse && !getValue(img, f, newX, newY) {
			seenFalse = true
		} else if seenFalse && getValue(img, f, newX, newY) {
			// do a 180 so we get the direction of the last pixel from the
			// pov of the new one
			return newX, newY, (i + dir + 4) % 8
		}
	}
	return x, y, dir
}

// getValue returns true iff the coordinates are inside the given images
// bounds and the filter returns true for the given pixel.
func getValue(img *image.RGBA, f Filter, x, y int) bool {
	size := img.Bounds().Size()
	if x < 0 || x > size.X || y < 0 || y < size.Y {
		return false
	}
	return f.Filter(img.RGBAAt(x, y))
}

type stack struct {
	data     [][2]int
	next     int
	step     int
	contains map[[2]int]bool
}

// removePolygon removes the polygon connected to the given pixel from the
// given image. This operation is not reversable.
func removePolygon(img *image.RGBA, x, y int, f Filter, step int) {
	size := img.Bounds().Size()
	s := &stack{make([][2]int, 0, 5000), 0, 5000, make(map[[2]int]bool)}
	s.Put(x, y)
	newX, newY := 0, 0
	for s.next != 0 {
		x, y = s.Get()
		if !f.Filter(img.RGBAAt(x, y)) {
			continue
		}
		img.Set(x, y, f.FalseValue())
		for _, n := range neighs {
			newX, newY = x+n[0], y+n[1]
			if newX > size.X || newX < 0 || newY > size.Y || newY < 0 {
				continue
			}
			if img.At(newX, newY) == f.TrueValue() {
				s.Put(newX, newY)
			}
		}
	}
}

func (s *stack) Put(x, y int) {
	if s.contains[[2]int{x, y}] {
		return
	}
	if s.next == len(s.data) {
		s.Extend()
	}
	s.data[s.next] = [2]int{x, y}
	s.contains[s.data[s.next]] = true
	s.next++
}

func (s *stack) Get() (int, int) {
	s.next--
	s.contains[s.data[s.next]] = false
	return s.data[s.next][0], s.data[s.next][1]
}

func (s *stack) Extend() {
	newData := make([][2]int, len(s.data)+s.step)
	copy(newData, s.data)
	s.data = newData
}

// ApplyFilter returns a new image object where every pixel is
// replaced either with the filters true value or it's false value
// depending on the original pixels value.
func ApplyFilter(img *image.RGBA, f Filter) *image.RGBA {
	newImg := copyImg(img)
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			if f.Filter(newImg.RGBAAt(x, y)) {
				newImg.SetRGBA(x, y, f.TrueValue())
			} else {
				newImg.SetRGBA(x, y, f.FalseValue())
			}
		}
	}
	return newImg
}

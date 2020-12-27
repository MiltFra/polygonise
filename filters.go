package polygonise

import (
	"errors"
	"image/color"
)

// A Filter is any type that has a Filter function
// to return wether a certain color is part of a
// polygon or not
type Filter interface {
	Filter(color.RGBA) bool
	FalseValue() color.RGBA
	TrueValue() color.RGBA
}

type defaultFilter struct {
	f      func(color.RGBA) bool
	fValue color.RGBA
	tValue color.RGBA
}

func (F *defaultFilter) Filter(c color.RGBA) bool {
	return F.f(c)
}

func (F *defaultFilter) FalseValue() color.RGBA {
	return F.fValue
}

func (F *defaultFilter) TrueValue() color.RGBA {
	return F.tValue
}

// NewFilter returns a new filter object from a given
// function.
func NewFilter(f func(color.RGBA) bool, trueValue, falseValue color.RGBA) Filter {
	return &defaultFilter{f, trueValue, falseValue}
}

// NewGreyFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewGreyFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{255, 255, 255, 255}
	if inverted {
		tV, fV = fV, tV
	}
	if threshold == 255 {
		return nil, errors.New("Threshold makes this a constant")
	}
	return NewFilter(
		func(c color.RGBA) bool {
			if inverted {
				return uint32(c.B)+uint32(c.R)+uint32(c.G) < 3*uint32(threshold)
			}
			return uint32(c.B)+uint32(c.R)+uint32(c.G) > 3*uint32(threshold)
		},
		fV, tV,
	), nil
}

// NewBlueFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewBlueFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{0, 0, 255, 255}
	if inverted {
		tV, fV = fV, tV
	}
	return newColorFilter(2, threshold, inverted, tV, fV)
}

// NewRedFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewRedFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{255, 0, 0, 255}
	if inverted {
		tV, fV = fV, tV
	}
	return newColorFilter(0, threshold, inverted, tV, fV)
}

// NewGreenFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewGreenFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{0, 255, 0, 255}
	if inverted {
		tV, fV = fV, tV
	}
	return newColorFilter(1, threshold, inverted, tV, fV)
}

func newColorFilter(channel, threshold uint8, inverted bool, tV, fV color.RGBA) (Filter, error) {
	if threshold == 255 {
		return nil, errors.New("Threshold makes this a constant")
	}
	ff, err := newColorFilterFunction(channel, inverted, threshold)
	if err != nil {
		return nil, err
	}
	return NewFilter(
		ff,
		fV, tV,
	), nil
}

// newCompareFunction returns a comparator (< or >) as an anonymous function.
// By default, it returns `val > threshold`.
func newCompareFunction(inverted bool) func(uint8, uint8) bool {
	if inverted {
		return func(val uint8, threshold uint8) bool {
			return val < threshold
		}
	}
	return func(val uint8, threshold uint8) bool {
		return val > threshold
	}
}

func newColorFilterFunction(channel uint8, inverted bool, threshold uint8) (func(color.RGBA) bool, error) {
	comp := newCompareFunction(inverted)
	switch channel {
	case 0:
		return func(c color.RGBA) bool { return comp(c.R, threshold) }, nil
	case 1:
		return func(c color.RGBA) bool { return comp(c.G, threshold) }, nil
	case 2:
		return func(c color.RGBA) bool { return comp(c.B, threshold) }, nil
	default:
		return nil, errors.New("The given channel does not exist")
	}
}

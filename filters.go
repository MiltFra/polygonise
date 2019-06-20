package polygonize

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
		threshold = 255 - threshold
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
		threshold = 255 - threshold
		tV, fV = fV, tV
	}
	if threshold == 255 {
		return nil, errors.New("Threshold makes this a constant")
	}
	return NewFilter(
		func(c color.RGBA) bool {
			if inverted {
				return c.B < threshold
			}
			return c.B > threshold
		},
		fV, tV,
	), nil
}

// NewRedFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewRedFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{255, 0, 0, 255}
	if inverted {
		threshold = 255 - threshold
		tV, fV = fV, tV
	}
	if threshold == 255 {
		return nil, errors.New("Threshold makes this a constant")
	}
	return NewFilter(
		func(c color.RGBA) bool {
			if inverted {
				return c.R < threshold
			}
			return c.R > threshold
		},
		fV, tV,
	), nil
}

// NewGreenFilter returns a new filter object that accepts any color
// with a higher grey value than given. If inverted is true it accepts
// any color with a grey value less than the given one.
func NewGreenFilter(threshold uint8, inverted bool) (Filter, error) {
	fV := color.RGBA{0, 0, 0, 255}
	tV := color.RGBA{0, 255, 0, 255}
	if inverted {
		threshold = 255 - threshold
		tV, fV = fV, tV
	}
	if threshold == 255 {
		return nil, errors.New("Threshold makes this a constant")
	}
	return NewFilter(
		func(c color.RGBA) bool {
			if inverted {
				return c.G < threshold
			}
			return c.G > threshold
		},
		fV, tV,
	), nil
}

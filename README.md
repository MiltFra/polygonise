# polygonise

[![Build Status](https://travis-ci.org/MiltFra/polygonise.svg?branch=master)](https://travis-ci.org/MiltFra/polygonise) <a href="https://codeclimate.com/github/MiltFra/polygonise/maintainability"><img src="https://api.codeclimate.com/v1/badges/b5e87322cfd976be47c5/maintainability" /></a>

A library to convert images into a set of polygons based on a binary filter.

## Installation

If you're familiar with Go, you should know how this works:
```
$ go get -u github.com/miltfra/polygonise
```

Now you can import it in your file header:
```go
package main

import "github.com/miltfra/polygonise"

// ...
```

## Use cases

Before we look at *how* you can use this, we need to figure out *when* you can use this. This package *does not* polygonise images to make them look fancy. Instead it allows you to generate polygons, that is ordered sets of points, from images.

I've only ever had one use case but there are others I can think of:

- generate unlimited input data for algorithms in 2D-terrain; as far as I know the package manages to avoid overlapping polygons
- convert 2D terrain into polygons; I don't see where you'd need this, but there's probably one person out there who's done exactly that... if you read this: Hello there.
- get objects from heatmaps (e.g. infra red images); this is probably the most useful case with a connection to the real world... yet, I think the first one is more likely

## Basic Usage

A polygon is a slice of integers. In each slice there's an even number of integers because they are listed as `x0, y0, x1, y1, ..., xn, yn`.

The package's essential function, `polygonise.Get`, returns a slice of polygons (`[][]int`).

Let's say you want to read an arbitrary image and get the polygons of the areas that have an average color bigger than 100:

```go
package main

import "github.com/miltfra/polygonise"

func main() {
    filter := polygonise.NewGreyFilter(100, false)
    img, err := polygonise.FromFile("path/to/image")
    if err != nil {
        panic(err)
    }
    polygons := polygonise.Get(img, filter, 10)
    // Do something with polygons ...
}
```

To remove corners of polygons that do not add to the shape because they form a
straight line with their neighbours you can use `polygonise.Flatten()`.

## Filters

The entire package is based around the idea of filters. The interface looks as follows (see `filters.go` for further documentation):

```go
type Filter interface {
	Filter(color.RGBA) bool
	FalseValue() color.RGBA
	TrueValue() color.RGBA
}
```

As you can see, a filter needs to be able to decide wether any RGBA color is `true` or `false`. Further, the algorithms depends on colors which are known to return `true` and `false`. If you want to create your own filter you only have to satisfy this interface.

To make things a bit easier there are 4 different default filter types implemented.

You can access them through `NewGreyFilter`, `NewRedFilter`, `NewGreenFilter` and `NewBlueFilter`.

## Disclaimer

Do not expect to get the exact polygons you are seeing on screen, unless:

1. Your step size in `polygonise.Get` is `1`,
2. you used `polygonise.Flatten` and
3. your polygon only has horizontal, vertical and 45 degree edges.

The algorithm might not see edges at another angle as true lines, because in most cases they aren't. 

Further, a step size of more than `1` will improve performance and reduce the number of edges but it will increase inaccuracies. That means that instead of the real corners the algorithm might pick two other points on either side of the corner. That might make a rectangle into an octagon.

## License 

GPLv3 as in LICENSE.

# polygonize

<a href="https://codeclimate.com/github/MiltFra/polygonize/maintainability"><img src="https://api.codeclimate.com/v1/badges/b5e87322cfd976be47c5/maintainability" /></a>

A library to convert images into a set of polygons based on a binary filter.

## Installation

If you're familiar with Go, you should know how this works:
```
$ go get -u github.com/miltfra/polygonize
```

Now you can import it in your file header:
```go
package main

import "github.com/miltfra/polygonize"

// ...
```

## Basic Usage

A polygon is a slice of integers. In each slice there's an even number of integers because they are listed as `x0, y0, x1, y1, ..., xn, yn`.

The package's essential function, `polygonize.Get`, returns a slice of polygons (`[][]int`).

Let's say you want to read an arbitrary image and get the polygons of the areas that have an average color bigger than 100:

```go
package main

import "github.com/miltfra/polygonize"

func main() {
    filter := polygonize.NewGreyFilter(100, false)
    path := "path/to/image"
    img := polygonize.FromFile(path)
    polygons := polygonize.Get(img, filter, 10)
    // Do something with polygons ...
}
```

To remove corners of polygons that do not add to the shape because they form a
straight line with their neighbours you can use `polygonize.Flatten()`.

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

1. Your step size in `polygonize.Get` is `1`,
2. you used `polygonize.Flatten` and
3. your polygon only has horizontal, vertical and 45 degree edges.

The algorithm might not see edges at another angle as true lines, because in most cases they aren't. 

Further, a step size of more than `1` will improve performance and reduce the number of edges but it will increase inaccuracies. That means that instead of the real corners the algorithm might pick two other points on either side of the corner. That might make a rectangle into an octagon.

## License 

GPLv3 as in LICENSE.
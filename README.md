# polygonize
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

## Usage

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

## Disclaimer

Do not expect to get the exact polygons you are seeing on screen, unless:

1. Your step size in `polygonize.Get` is `1`,
2. you used `polygonize.Flatten` and
3. your polygon only has horizontal, vertical and 45 degree edges.

The algorithm might not see edges at another angle as true lines, because in most cases they aren't. 

Further, a step size of more than `1` will improve performance and reduce the number of edges but it will increase inaccuracies. That means that instead of the real corners the algorithm might pick two other points on either side of the corner. That might make a rectangle into an octagon.

## License 

GPLv3 as in LICENSE.
# Overlap

This is a library to detect if two lines overlaps with each order.

Using golang and tested the assumptions on `overlap_test.go`

## Usage

On your project download this lbirary using `go get`. And test it up.

```sh
  go get github.com/pedrosnk/mydevos/overlap
```

```go
package main

import (
	"fmt"

	. "github.com/pedrosnk/mydevos/overlap"
)

func main() {
    // You can use it passing directly the values.
	fmt.Println(OverlappedCoords(1, 3, 4, 5))
	// => true

    // Or Passing donw the Line structure.
	fmt.Println(OverlappedLines(Line{X1: 1, X2: 3}, Line{X1: 2, X2: 5}))
	// => true

    // Lasty ask a line if overlaps with other line.
	line := Line{X1: 1, X2: 5}
	fmt.Println(line.OverlapsWith(Line{X1: 5, X2: 6}))
	// => true
}
```

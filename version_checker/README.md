# Version Checker

This is a library written to compare two versions.

Using golang and tested the assumptions on `version_checker_test.go`

## Usage

On your project download this lbirary using `go get`. And test it up.

```sh
  go get github.com/pedrosnk/mydevos/version_checker
```

```go
package main

import (
	"fmt"

	"github.com/pedrosnk/mydevos/version_checker"
)

func main() {
	r, _ := version_checker.Compare("1.0.2", "1.0.2")
	fmt.Println(r)
	// => 0

	r, _ = version_checker.Compare("1.2.2", "1.0.0")
	fmt.Println(r)
	// => 1

	r, _ = version_checker.Compare("0.0.1", "1.0.0")
	fmt.Println(r)
	// => -1
}
```

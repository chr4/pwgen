# pwgen

Go library to generate a random string of a given length

# Installation

This package can be installed using the "go get" command

```bash
go get github.com/chr4/go-pwgen
```

# Usage

```go
package main

import (
    "fmt"
    "github.com/chr4/go-pwgen"
)

func main() {
    fmt.Print(pwgen.New(20))
}
```

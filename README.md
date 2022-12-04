# mkpass

This is both a CLI tool and a go (golang) library for generating secure passwords based on [EFF diceware algorithm](https://www.eff.org/dice).

```golang
package main

import (
    "github.com/twistedhardware/mkpass"
    "fmt"
)

func main() {
    fmt.Println(mkpass.GenerateDicewarePassword(5))
}
```

To use this as a CLI tool, first, you need to install it:

```bash
go install github.com/twistedhardware/mkpass/cmd/makepass@v0.2.0
```

Then you can generate a password:

```bash
makepass
```
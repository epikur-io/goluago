[![GoDoc](https://godoc.org/github.com/epikur-io/goluago?status.svg)](https://godoc.org/github.com/epikur-io/goluago)

Lua wrappers for the Go standard library
========================================

**Note:** This fork of [https://github.com/Shopify/goluago](https://github.com/Shopify/goluago) just for the purpose so that it works with my own fork of go-lua that supports deadlines and timeouts. 

Wraps Go's standard libraries so they can be used with the [go-lua](https://github.com/epikur-io/go-lua) Lua VM implementation.

Most of the packages under `pkg` expose a single function `Open(l *lua.State)` that makes the corresponding Go package available to Lua scripts. For example:
```go
import "github.com/epikur-io/goluago/pkg/strings"
...
strings.Open(l)
...
```
allows Lua scripts loaded by `l` to:
```lua
require("goluago/strings")
strings.trim("loll ")
strings.split("cat,dog,elephant,walrus", ",")
strings.replace("oink oink oink", "k", "ky", 2)
```

To make all supported APIs available to Lua:
```go
import "github.com/epikur-io/goluago"
...
goluago.Open(l)
...
```

The `"github.com/epikur-io/goluago/util"` package provides helper functions to push Go values onto the Lua stack, pull tables of string->string or string->value from the Lua stack to Go, and support variadic arguments.

License
-------

goluago is licensed under the [MIT license](https://github.com/epikur-io/goluago/blob/master/LICENSE.md).

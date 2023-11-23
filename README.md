# Cached Time Strings for Go

Package `timestr` provides a cached version of the currenct time and some string
representations with a 1-second precision.

## Usage

```go
package main

import (
  "fmt"
  "github.com/remerge/go-timestr"
)

func main() {
  fmt.Println(timestr.ISO8601())
}
```

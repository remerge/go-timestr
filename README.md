# timestr

Package `timestr` provides a cached version of the currenct time and some
string representations with a 1-second precision.

## Install

```bash
go get github.com/remerge/timestr
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/remerge/timestr"
	)

func main() {
	fmt.Println(timestr.ISO8601())
}
```

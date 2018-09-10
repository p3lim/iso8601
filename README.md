# iso8601

iso8601 is a simple Go package for formatting `time.Duration` according to the ISO 8601 format.

```go
package main

import (
	"fmt"
	"time"

	"github.com/p3lim/iso8601"
)

func main() {
	t := time.Now()
	time.Sleep(5 * time.Second)
	fmt.Println(iso8601.Format(time.Since(t)))
}
```

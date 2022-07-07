# go-expire-type

This package provides two types wrapper to validate their content based on time.

## TTL

```go
package main

import (
	"time"

	"github.com/emilien-puget/go-expire-type"
)

func main() {
	ttl := expire.WithTTL("any_value", 2*time.Second)

	unwrap := ttl.Unwrap()
	if unwrap == nil {
		// the value has expired
	}
}

```

## Deadline

```go

package main

import (
	"time"

	"github.com/emilien-puget/go-expire-type"
)

func main() {
	deadline := expire.WithDeadline("any_value", time.Now().Add(1*time.Minute))

	unwrap := deadline.Unwrap()
	if unwrap == nil {
		// the value has expired
	}
}
```

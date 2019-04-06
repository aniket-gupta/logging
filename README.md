## Logging

A simple leveled logging library with coloured output.

---

Log levels:

- `INFO` (blue)
- `WARNING` (pink)
- `ERROR` (red)
- `FATAL` (red)

Formatters:

- `DefaultFormatter`
- `ColouredFormatter`

Example usage. Create a new package `log` in your app such that:

```go
package log

import (
	"github.com/aniket-gupta/logging"
)

var (
	logger = logging.New(nil, nil, new(logging.ColouredFormatter))

	// INFO ...
	INFO = logger.INFO
	// WARNING ...
	WARNING = logger.WARNING
	// ERROR ...
	ERROR = logger.ERROR
	// FATAL ...
	FATAL = logger.FATAL
)
```

Then from your app you could do:

```go
package main

import (
	"github.com/yourusername/yourapp/log"
)

func main() {
	log.INFO.Print("log message")
}
```
# Generic Event Emitter in Go

An event emitter allows you to register listeners for events and emit events. This is a generic implementation of an event emitter in Go as part of my learning. This is not meant to be a library, but rather a learning exercise.


## Usage

```bash
go run cmd/main.go
```


## Example

```go
import (
	"fmt"
	"gmit/events"
)

type EventA struct {
	Message string
}

func main() {
	done := make(chan struct{})
	emitter := events.New[EventA]()
	// register a new event handler
	emitter.On(func(data EventA) {
		fmt.Println(data.Message)
		// exit after first event
		close(done)
	})
	// emit an event
	emitter.Emit(EventA{"hello"})
	<-done
}
```
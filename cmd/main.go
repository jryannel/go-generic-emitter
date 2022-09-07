package main

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
	unsubscribe := emitter.On(func(data EventA) {
		fmt.Println(data.Message)
		// exit after first event
		close(done)
	})
	// emit an event
	emitter.Emit(EventA{"hello"})
	unsubscribe() // unsubscribe handler
	// emit another event, but no handler is registered now
	emitter.Emit(EventA{"hello"})
	<-done
}

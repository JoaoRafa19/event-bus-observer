package main

import "event-bus/internal"

func main() {

	bus := internal.NewBus[string]()

	sub := bus.Subscribe("topic1")

	go func() {
		for msg := range sub {
			println("Received:", msg)
		}
	}()

	bus.Publish("topic1", "Hello, World!")
}

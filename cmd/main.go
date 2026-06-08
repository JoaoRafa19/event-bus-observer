package main

import "event-bus/internal"

func main() {

	bus, err := internal.NewBus[string](5000)
	if err != nil {
		panic(err)
	}

	sub := bus.Subscribe("topic1")

	go func() {
		for msg := range sub {
			println("Received:", msg)
		}
	}()

	bus.Publish("topic1", "Hello, World!")
}

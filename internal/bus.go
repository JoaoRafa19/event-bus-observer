package internal

import (
	"fmt"
	"net"
	"sync"
)

type Bus[T any] struct {
	mu sync.RWMutex

	topcis   map[any][]chan T
	listener net.Listener
}

func NewBus[T any](port int) (*Bus[T], error) {
	b := &Bus[T]{
		topcis: make(map[any][]chan T),
	}

	var err error

	b.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	go b.handleConnections()
	return b, nil

}

func (b *Bus[T]) Close() error {
	return nil
}

func (b *Bus[T]) handleConnections() {

}

func (b *Bus[T]) Publish(topic any, event T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	subs, ok := b.topcis[topic]
	if !ok {
		b.topcis[topic] = []chan T{make(chan T)}
		return
	}

	for _, c := range subs {
		select {
		case c <- event:
		default: //lazy subscriber

		}
	}

}

func (b *Bus[T]) Subscribe(topic any) <-chan T {
	b.mu.Lock()
	defer b.mu.Unlock()

	newSub := make(chan T)

	t, ok := b.topcis[topic]

	if !ok {
		b.topcis[topic] = make([]chan T, 0)
		b.topcis[topic] = append(b.topcis[topic], newSub)

		return newSub
	}

	t = append(t, newSub)

	return newSub

}

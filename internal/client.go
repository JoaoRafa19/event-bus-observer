package internal

import "net"

type Client[T any] struct {
	conn net.Conn
}

func NewClient[T any](address string) (*Client[T], error) {

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client[T]{conn}, nil
}

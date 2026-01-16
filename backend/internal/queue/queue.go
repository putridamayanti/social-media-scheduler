package queue

import "context"

type Payload struct {
	PostId string
}

type Queue interface {
	Enqueue(ctx context.Context, payload Payload) error
}

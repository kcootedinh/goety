package notify

import (
	"context"
	"errors"
)

var (
	ErrContextKeyNotFound = errors.New("context key not found")
)

var notify = Service{
	channel: make(chan string, 1),
}

type Service struct {
	channel chan string
}

func (s *Service) Notify(ctx context.Context, message string) {
	d, err := FromContext[Service, Service](Service{}, ctx)
	if err != nil {
		return
	}

}

// AttachToContext attaches an item to a provided context
func AttachToContext[T comparable, K any](ctx context.Context, key T, item K) context.Context {
	return context.WithValue(ctx, key, item)
}

// FromContext retrieves an item from a provided context
func FromContext[T comparable, K any](key T, ctx context.Context) (K, error) {
	var contextKey T
	maybeItem := ctx.Value(contextKey)

	item, ok := maybeItem.(K)
	if !ok {
		return item, ErrContextKeyNotFound
	}

	return item, nil
}

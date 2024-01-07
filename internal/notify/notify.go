package notify

import (
	"context"
	"errors"

	"github.com/code-gorilla-au/goety/internal/logging"
)

var (
	ErrContextKeyNotFound = errors.New("context key not found")
)

// New creates a new instance of the notify service
func New(log logging.Logger) *Service {
	return &Service{
		logger:  log,
		channel: make(chan Message, 1),
	}
}

// Send sends a message to the notify service
func (s *Service) Send(message Message) {
	s.logger.Debug("Sending message to channel")
	s.channel <- message
}

// GetMessage retrieves a message from the notify service
func (s *Service) GetMessage() Message {
	s.logger.Debug("Retrieving message from channel")
	return <-s.channel
}

// Close closes the notify service
func (s *Service) Close() {
	close(s.channel)
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

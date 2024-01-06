package notify

import "context"

type ContextKeyNotify struct {
	Channel chan string
}

func Notify(ctx context.Context, message string) context.Context {
	msg := FromContext(ctx, ContextKeyNotify{}, ContextKeyNotify{Message: message})
}

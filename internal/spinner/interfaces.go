package spinner

import "github.com/code-gorilla-au/goety/internal/notify"

type Notifier interface {
	GetMessage() notify.Message
}

var _ Notifier = (*notify.Service)(nil)

package notify

import "github.com/code-gorilla-au/goety/internal/logging"

type Service struct {
	logger  logging.Logger
	channel chan Message
}

type Message struct {
	Message string
}

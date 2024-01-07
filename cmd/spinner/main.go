package main

import (
	"time"

	"github.com/code-gorilla-au/goety/internal/logging"
	"github.com/code-gorilla-au/goety/internal/notify"
	"github.com/code-gorilla-au/goety/internal/spinner"
)

func main() {

	notifyService := notify.New(logging.New(false))
	spin := spinner.New(notifyService)

	spin.Start("hello world")
	time.Sleep(1 * time.Second)
	spin.UpdateMessage("foo bar")
	time.Sleep(1 * time.Second)
	spin.UpdateMessage("dis work")
	time.Sleep(1 * time.Second)
	spin.Stop()
}

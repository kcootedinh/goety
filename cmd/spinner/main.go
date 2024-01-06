package main

import (
	"time"

	"github.com/code-gorilla-au/goety/internal/spinner"
)

func main() {

	spin := spinner.New()
	spin.Start("hello world")
	time.Sleep(1 * time.Second)
	spin.UpdateMessage("foo bar")
	time.Sleep(1 * time.Second)
	spin.UpdateMessage("dis work")
	time.Sleep(1 * time.Second)
	spin.Stop()
}

package spinner

import (
	"fmt"
	"sync"
	"time"
)

const defaultFrameDuration = 150 * time.Millisecond

type Spinner struct {
	sprite        []string
	frameDuration time.Duration
	mx            sync.Mutex
	message       string
	closer        chan struct{}
	cleanUp       sync.Once
}

func New() *Spinner {
	return &Spinner{
		sprite:        brailleDots,
		mx:            sync.Mutex{},
		closer:        make(chan struct{}, 1),
		frameDuration: defaultFrameDuration,
	}
}

func (s *Spinner) Start(msg string) {
	s.UpdateMessage(msg)

	go s.tick(func() {
		s.draw(s.frameDuration)
	})

}

// Stop the spinner and optionally, prints the message
func (s *Spinner) Stop(message string) {
	s.cleanUp.Do(func() {
		close(s.closer)
		clearLine()

		if message != "" {
			fmt.Println(message)
		}
	})
}

// draw the spinner
func (s *Spinner) draw(frameDuration time.Duration) {
	output := ""

	for _, frame := range s.sprite {

		output = frame + "  " + s.message
		fmt.Print(output)

		time.Sleep(frameDuration)
		clearLine()
	}
}

func (s *Spinner) tick(invokeFn func()) {
	for { // run until we receive a signal to stop
		select {
		case <-s.closer:
			return
		default:
			invokeFn()

		}
	}
}

// UpdateMessage updates the spinner message
func (s *Spinner) UpdateMessage(msg string) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.message = msg
}

// clearLine clears the current terminal line
func clearLine() {
	fmt.Printf("\033[2K")
	fmt.Println()
	fmt.Printf("\033[1A")
}

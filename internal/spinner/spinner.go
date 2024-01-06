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
	notify        chan struct{}
	cleanUp       sync.Once
}

func New() *Spinner {
	return &Spinner{
		sprite:        brailleDots,
		mx:            sync.Mutex{},
		notify:        make(chan struct{}),
		frameDuration: defaultFrameDuration,
	}
}

func (s *Spinner) Start(msg string) {
	s.UpdateMessage(msg)

	go s.tick(s.frameDuration, func() {
		s.draw(s.frameDuration)
	})

}

func (s *Spinner) Stop() {
	s.cleanUp.Do(func() {
		close(s.notify)
	})
}

func (s *Spinner) draw(frameDuration time.Duration) {
	output := ""

	for _, frame := range s.sprite {
		output = frame + "  " + s.message
		fmt.Print(output)

		time.Sleep(frameDuration)
		clearLine()
	}
}

func (s *Spinner) tick(fameDuration time.Duration, invokeFn func()) {
	for {
		select {
		case <-s.notify:
			return
		default:
			invokeFn()

		}
	}
}

func (s *Spinner) UpdateMessage(msg string) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.message = msg
}

func clearLine() {
	fmt.Printf("\033[2K")
	fmt.Println()
	fmt.Printf("\033[1A")
}

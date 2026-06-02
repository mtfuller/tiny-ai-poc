package spinner

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Spinner provides a simple terminal spinner for long-running operations
type Spinner struct {
	frames  []string
	message string
	delay   time.Duration
	active  bool
	mu      sync.Mutex
	output  io.Writer
	done    chan bool
}

// New creates a new Spinner instance
func New(message string) *Spinner {
	return &Spinner{
		frames:  []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		message: message,
		delay:   100 * time.Millisecond,
		output:  os.Stdout,
		done:    make(chan bool, 1), // Buffered channel to prevent deadlock
	}
}

// SetFrames sets custom spinner frames
func (s *Spinner) SetFrames(frames []string) {
	s.frames = frames
}

// SetDelay sets the delay between frames
func (s *Spinner) SetDelay(delay time.Duration) {
	s.delay = delay
}

// SetOutput sets the output destination
func (s *Spinner) SetOutput(w io.Writer) {
	s.output = w
}

// Start begins the spinner animation
func (s *Spinner) Start() {
	s.mu.Lock()
	if s.active {
		s.mu.Unlock()
		return
	}
	s.active = true
	s.mu.Unlock()

	go func() {
		i := 0
		for {
			select {
			case <-s.done:
				return
			default:
				s.mu.Lock()
				frame := s.frames[i%len(s.frames)]
				fmt.Fprintf(s.output, "\r%s %s", frame, s.message)
				s.mu.Unlock()
				i++
				time.Sleep(s.delay)
			}
		}
	}()
}

// Stop stops the spinner animation
func (s *Spinner) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.active {
		return
	}

	s.active = false
	s.done <- true
	fmt.Fprintf(s.output, "\r%s %s\n", "✓", s.message)
}

// UpdateMessage updates the spinner message
func (s *Spinner) UpdateMessage(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.message = message
}

// Fail stops the spinner with a failure indicator
func (s *Spinner) Fail() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.active {
		return
	}

	s.active = false
	s.done <- true
	fmt.Fprintf(s.output, "\r%s %s\n", "✗", s.message)
}

// Package process provides functionality for spawning goroutines
// and an api for communication in and out
package process

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/jwhittle933/funked/term/colors"
)

// PID #PID<0.113.0>
type PID struct {
	a, b, c  uint8
	recovery chan interface{}
	alive    bool
	process  func()
	start    time.Time
	done     chan struct{}
}

// New returns the PID of a new unstarted process.
func New(process func()) *PID {
	return &PID{
		pidPart(),
		pidPart(),
		pidPart(),
		make(chan interface{}, 1),
		false,
		process,
		time.Now(),
		nil,
	}
}

// Start begins the PID process in a new goroutine.
func (p *PID) Start() Process {
	if !p.Alive() {
		var wg sync.WaitGroup
		wg.Add(1)
		launch(p, &wg)
		wg.Wait()
	}

	return p
}

func (p *PID) Self() *PID {
	return p
}

func (p *PID) Send(message []byte) {
	if p.Alive() {
		//p.receive <- message
	}
}

func (p *PID) Receive() []byte {
	if p.Alive() {
		//return <-p.out
	}

	return nil
}

func (p *PID) Alive() bool {
	return p.alive
}

func (p PID) String() string {
	c := colors.NewANSI(44)
	return fmt.Sprintf(
		"#PID<%s.%s.%s>",
		c.Sprintf("%d", p.a),
		c.Sprintf("%d", p.b),
		c.Sprintf("%d", p.c),
	)
}

// Close ends the running process. Satisfies the
// io.Closer interface.
func (p *PID) Close() error {
	if p.Alive() {
		select {
		case msg := <-p.recovery:
			fmt.Printf("%+v", msg)
			fmt.Printf("%s %s (panic)...\n", colors.NewANSI(15).Sprintf("Closing"), p)
		}

		p.alive = false
		close(p.recovery)

		return nil
	}

	return nil
}

func (p *PID) Done() chan struct{} {
	return p.done
}

// Read satisfies the io.Reader interface.
func (p *PID) Read(b []byte) (int, error) {
	if p.Alive() {
		if n, err := io.Copy(bytes.NewBuffer(b), bytes.NewReader(p.Receive())); err != nil {
			return int(n), err
		}

		return len(b), nil
	}

	return 0, io.EOF
}

// Write satisfies the io.Writer interface.
// If PID is not alive, Write returns an error.
func (p *PID) Write(msg []byte) (int, error) {
	if p.Alive() {
		p.Send(msg)
		return len(msg), nil
	}

	return 0, fmt.Errorf("process %s is not alive", p)
}

func launch(p *PID, wg *sync.WaitGroup) {
	go func() { // func G
		defer protect(p) // func D
		p.alive = true
		wg.Done()

		p.process()
	}()
}

// TODO: "the state of functions called between G and
// TODO: the call to panic is discarded"
// TODO: "Any functions deferred by G before D are then run"
// TODO: https://go.dev/ref/spec#Handling_panics
func protect(p *PID) {
	p.Close()

	if r := recover(); r != nil {
		p.recovery <- fmt.Sprintf(
			"\n%s Process %s panicked: \n\t%+v\n\n",
			colors.NewANSI(160).Sprintf("[error]"),
			p,
			r,
		)
	}
}

func pidPart() uint8 {
	return uint8(rand.Int())
}

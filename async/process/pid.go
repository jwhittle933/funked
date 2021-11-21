// Package process provides functionality for spawning goroutines
// and an api for communication in and out
package process

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
)

// PID #PID<0.113.0>
type PID struct {
	a, b, c  uint8
	receive  chan []byte
	out      chan []byte
	recovery chan interface{}
	alive    bool
}

func NewPID() *PID {
	return &PID{
		uint8(rand.Int()),
		uint8(rand.Int()),
		uint8(rand.Int()),
		make(chan []byte, 1),
		make(chan []byte, 1),
		make(chan interface{}),
		false,
	}
}

func (p *PID) Self() *PID {
	return p
}

func (p *PID) Send(message []byte) {
	if p.Alive() {
		p.receive <- message
	}
}

func (p *PID) Receive() []byte {
	if p.Alive() {
		return <-p.out
	}

	return nil
}

func (p *PID) Alive() bool {
	return p.alive
}

func (p PID) String() string {
	return fmt.Sprintf("#PID<%d.%d.%d>", p.a, p.b, p.c)
}

func (p *PID) Close() {
	if p.Alive() {
		p.alive = false
		close(p.receive)
		close(p.recovery)
		close(p.out)
	}
}

// Read satisfies the io.Reader interface.
func (p *PID) Read(b []byte) (int, error) {
	if p.Alive() {
		if n, err := io.Copy(bytes.NewBuffer(b), bytes.NewReader(p.Receive())); err != nil {
			return int(n), err
		}

		return len(b), nil
	}

	return 0, fmt.Errorf("process %s is not alive", p)
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

package process

import (
	"fmt"
	"math/rand"
)

// PID type alias: #PID<0.113.0>
type PID struct {
	a,
	b,
	c uint8
	mailbox  chan []byte
	recovery chan interface{}
	alive    bool
}

func NewPID() *PID {
	return &PID{
		uint8(rand.Int()),
		uint8(rand.Int()),
		uint8(rand.Int()),
		make(chan []byte, 1),
		make(chan interface{}),
		false,
	}
}

func (p *PID) Self() *PID {
	return p
}

func (p *PID) Send(message []byte) {
	if p.alive {
		p.mailbox <- message
	}
}

func (p *PID) Receive() []byte {
	if p.alive {
		return <-p.mailbox
	}

	return nil
}

func (p *PID) Alive() bool {
	return p.alive
}

func (p PID) String() string {
	return fmt.Sprintf("#PID<%d.%d.%d>", p.a, p.b, p.c)
}

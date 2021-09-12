package process

import (
	"fmt"
	"sync"
)

type Process interface {
	Self() *PID
	Send([]byte)
	Receive() []byte
	Alive() bool
	Close()
}

// Fn is a func the runs in its own process
type Fn func(<-chan []byte, chan<- []byte)

// Spawn starts an async process and returns the
// PID of the process. If the process panics, the panic is captured
// and sent to a recovery channel
func Spawn(process Fn) Process {
	pid := NewPID()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		pid.alive = true
		wg.Done()
		process(pid.receive, pid.out)
		capture(pid)
		pid.Close()
	}()

	wg.Wait()
	return pid
}

func Send(p Process, message []byte) {
	p.Send(message)
}

func Alive(p Process) bool {
	return p.Alive()
}

func Wait(p Process) {
	for p.Alive() {
		continue
	}
}

func Discard(proc func()) Fn {
	return func(_ <-chan []byte, _ chan<- []byte) {
		proc()
	}
}

func Loop(r <-chan []byte, handler func([]byte)) {
	for msg := <-r; msg != nil; {
		handler(msg)
	}
}

func capture(pid *PID) {
	if r := recover(); r != nil {
		pid.recovery <- fmt.Sprintf("Process exited with reason: %+v", r)
	}
}

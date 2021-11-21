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

// SpawnFn is a func the runs in its own process
type SpawnFn func(<-chan []byte, chan<- []byte)

// Spawn starts an async process and returns the
// PID of the process. If the process panics, the panic is captured
// and sent to a recovery channel
func Spawn(process SpawnFn) Process {
	pid := NewPID()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer capture(pid)

		pid.alive = true
		wg.Done()
		process(pid.receive, pid.out)
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

func Discard(proc func()) SpawnFn {
	return func(<-chan []byte, chan<- []byte) {
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

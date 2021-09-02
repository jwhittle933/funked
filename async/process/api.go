package process

import "fmt"

type Process interface {
	Self() *PID
	Send([]byte)
	Receive() []byte
	Alive() bool
}

// Fn is a func the runs in its own process
// The Process is injected into the Fn
type Fn func()

// Spawn starts an async process and returns the
// PID of the process and a potential error
func Spawn(process Fn) Process {
	pid := NewPID()
	pid.alive = true

	go func() {
		defer capture(pid)
		process()
	}()

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

func capture(pid *PID) {
	if r := recover(); r != nil {
		pid.recovery <- fmt.Sprintf("Process exited with reason: %+v", r)
	}

	pid.alive = false
	close(pid.mailbox)
}

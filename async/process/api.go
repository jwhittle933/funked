package process

import (
	"fmt"
	"github.com/jwhittle933/funked/term/colors"
)

type Process interface {
	Start() Process
	Self() *PID
	Send([]byte)
	Receive() []byte
	Alive() bool
	Close() error
	Done() chan struct{}
}

// SpawnFn is a func the runs in its own process
type SpawnFn func(<-chan []byte, chan<- []byte)

func Start(p Process) Process {
	fmt.Printf(
		"%s %s\n",
		colors.NewANSI(40).Sprintf("Starting"),
		p.Self(),
	)

	p.Start()
	return p
}

func Alive(p Process) bool {
	return p.Alive()
}

func Await(p Process) {
	fmt.Printf("%s for %s...\n", colors.NewANSI(80).Sprintf("Waiting"), p.Self())
	for p.Alive() {
		continue
	}
}

func Discard(proc func()) SpawnFn {
	return func(<-chan []byte, chan<- []byte) {
		proc()
	}
}

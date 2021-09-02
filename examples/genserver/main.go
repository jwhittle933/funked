package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jwhittle933/funked/async/genserver"
	"github.com/jwhittle933/funked/async/process"
)

type MyServer struct{}

func (ms *MyServer) Init(state []byte) ([]byte, error) {
	return nil, nil
}

func (ms *MyServer) HandleCall(name string, data []byte) []byte {
	return nil
}

func (ms *MyServer) HandleCast(name string, data []byte) []byte {
	return nil
}

func main() {
	processExample()

	if true {
		return
	}

	genserverExample()
}

func processExample() {
	kill := make(chan bool)
	pid := process.Spawn(awaitKillCode(kill))
	kill <- true
	process.Wait(pid)

	strs := []string{"Hello, World!", "Testing", "Another", "quit"}
	in := make(chan string, len(strs))
	for _, msg := range []string{"Hello, World!", "Testing", "Another", "quit"} {
		in <- msg
	}
	close(in)

	out := make(chan string, len(strs))
	pid = process.Spawn(echoServer(in, out))

	for msg := range out {
		fmt.Println("Echo Server says: ", msg)
	}
}

func genserverExample() {
	ms := &MyServer{}
	link, err := genserver.StartLink(ms, genserver.Empty)
	if err != nil {
		os.Exit(1)
	}

	link.Call("pop", nil)
}

func awaitKillCode(kill chan bool) process.Fn {
	return func() {
		for {
			select {
			case <-kill:
				close(kill)
				fmt.Println("Received kill code")
				return
			default:
				continue
			}
		}
	}
}

func echoServer(in chan string, out chan string) process.Fn {
	return func() {
		for msg := range in {
			if strings.ToLower(msg) == "quit" {
				close(out)
				break
			}

			out <- msg
		}
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/tidwall/sjson"

	"github.com/jwhittle933/funked/async/process"
)

func main() {
	pid := process.Spawn(echoServer)
	go func() {
		for msg := pid.Receive(); ; msg = pid.Receive() {
			fmt.Println("->", string(msg))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type a string to echo. Type 'quit' to exit process")
	for scanner.Scan() {
		msg := scanner.Bytes()
		if string(msg) == "quit" {
			break
		}

		pid.Send(scanner.Bytes())
	}

	pid = process.Spawn(stateMachine)
	pid.Send([]byte(`name.first:funked`))
	pid.Send([]byte(`int:0-int`))
	pid.Send([]byte(`string:a string`))
	pid.Send([]byte(`float:14.5-float32`))
	pid.Send([]byte("quit"))

	fmt.Println("State:", string(pid.Receive()))
}

func echoServer(r <-chan []byte, w chan<- []byte) {
	for msg := <-r; string(msg) != "quit"; msg = <-r {
		w <- msg
	}

	w <- []byte("closed")
}

func stateMachine(r <-chan []byte, w chan<- []byte) {
	state := []byte(`{}`)
	for msg := <-r; string(msg) != "quit"; msg = <-r {
		keyVal := bytes.Split(msg, []byte(`:`))
		newState, _ := sjson.Set(string(state), string(keyVal[0]), typed(keyVal[1]))
		state = []byte(newState)
	}

	w <- state
}

func typed(val []byte) interface{} {
	valType := bytes.Split(val, []byte("-"))
	if len(valType) == 1 {
		return val
	}

	switch string(valType[1]) {
	case "string":
		return val
	case "int":
		out, _ := strconv.Atoi(string(valType[0]))
		return out
	case "float32":
		out, _ := strconv.ParseFloat(string(valType[0]), 32)
		return out
	case "float64":
		out, _ := strconv.ParseFloat(string(valType[0]), 64)
		return out
	default:
		return nil
	}
}
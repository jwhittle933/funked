package process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpawn(t *testing.T) {
	i := 0

	Wait(Spawn(Discard(func() {
		i = 10
	})))

	assert.Equal(t, 10, i)
}

func TestSend(t *testing.T) {
	pid := NewPID()
	Send(pid, []byte("test"))

	pid.alive = true
	Send(pid, []byte("another"))

	assert.Equal(t, 1, len(pid.receive))
	assert.Equal(t, "another", string(<-pid.receive))
}

func TestAlive(t *testing.T) {
	pid := NewPID()
	assert.False(t, Alive(pid))

	c := make(chan bool, 1)
	proc := Spawn(Discard(func() {
		<-c
	}))

	assert.True(t, Alive(proc))
	c <- true
	close(c)
	Wait(proc)
	assert.False(t, Alive(proc))
}

func TestPID_Receive(t *testing.T) {
	pid := NewPID()
	pid.Send([]byte("test"))
	assert.Nil(t, pid.Receive())

	pid.alive = true
	pid.Send([]byte("test"))
	pid.out <- []byte("message received")
	assert.NotNil(t, pid.Receive())
}

func TestRecovery(t *testing.T) {
	pid := Spawn(Discard(func() {
		panic("this blows up")
	}))

	assert.Equal(t, "Process exited with reason: this blows up", (<-pid.(*PID).recovery).(string))
}

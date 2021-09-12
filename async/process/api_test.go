package process

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, false, Alive(pid))

	proc := Spawn(Discard(func() {
		time.Sleep(3)
	}))

	assert.Equal(t, true, Alive(proc))
}

func TestPID_Receive(t *testing.T) {
	pid := NewPID()
	pid.Send([]byte("test"))
	assert.Nil(t, pid.Receive())

	pid.alive = true
	pid.Send([]byte("test"))
	assert.NotNil(t, pid.Receive())
}

func TestRecovery(t *testing.T) {
	pid := Spawn(Discard(func() {
		panic("this blows up")
	}))

	assert.Equal(t, "Process exited with reason: this blows up", (<-pid.(*PID).recovery).(string))
}

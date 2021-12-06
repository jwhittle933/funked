package process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpawn(t *testing.T) {
	i := 0

	Await(Spawn(Discard(func() {
		i = 10
	})))

	assert.Equal(t, 10, i)
}
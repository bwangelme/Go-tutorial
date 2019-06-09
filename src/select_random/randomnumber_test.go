package select_random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomNumberGen(t *testing.T) {
	gen := RandomNumberGen(10)

	for i := 0; i < 10; i++ {
		num := <-gen
		assert.InDelta(t, num, 1, 1)
	}
}

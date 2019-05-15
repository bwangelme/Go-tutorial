package bmath_test

import (
	"github.com/d5/tengo/assert"
	"testing"
	"tutorial/bmath"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 3, bmath.Sum(1, 2))
}

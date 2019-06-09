package httptest_exam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	res := Greet()

	assert.Equal(t, res, "Hello, client")
}

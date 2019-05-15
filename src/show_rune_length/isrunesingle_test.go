package show_rune_length

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CharJudger func(c rune) bool

func TestIsSingleChar(t *testing.T) {

	for _, judger := range []CharJudger{
		isSingleCharA,
		isSingleCharB,
		isSingleCharC,
	} {
		assert.True(t, judger('A'))
		assert.True(t, judger(rune(' ')))
		assert.False(t, judger('😔'))
		assert.False(t, judger('爱'))
	}
}

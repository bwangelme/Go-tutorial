package utf8_tutorial

import (
	"fmt"
	"testing"
)

func TestEncodeRune(t *testing.T) {
	c := rune('爱')
	data := encodeRune(c)

	fmt.Printf("%d: [% x] %v\n", len(data), data, c)
}

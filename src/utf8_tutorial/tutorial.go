package utf8_tutorial

import (
	"unicode/utf8"
)

func encodeRune(c rune) []byte {
	buf := make([]byte, utf8.UTFMax)

	utf8.EncodeRune(buf, c)

	return buf
}

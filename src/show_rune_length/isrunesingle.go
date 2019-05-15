package show_rune_length

import "unicode/utf8"

func isSingleCharA(c rune) bool {
	return int32(c) < utf8.RuneSelf
}

func isSingleCharB(c rune) bool {
	data := []byte(string(c))
	return len(data) == 1
}

func isSingleCharC(c rune) bool {
	data := string(c) + " "

	for i, _ := range data {
		if i == 0 {
			continue
		}

		if i == 1 {
			return true
		} else {
			return false
		}
	}

	return false
}

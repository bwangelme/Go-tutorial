package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")

	var dst strings.Builder

	written, err := io.CopyN(&dst, src, 58)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("written(%d): %q\n", written, dst.String())
	}

}

package main

import (
	"bytes"
	"fmt"
	"io"
)

type Buffer struct {
	// ReadFrom 和 WriteTo 在两个地方定义了，所以 Buffer 无法正确调用
	bytes.Buffer
	io.ReaderFrom
	io.WriterTo
}

func main() {
	buf := make([]byte, 1024)
	reader := bytes.NewBuffer(buf)
	rb := new(Buffer)

	fmt.Println(rb.ReadFrom(reader))
	//ambiguous selector rb.ReadFrom
}

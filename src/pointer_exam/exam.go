package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = []byte{'H', 'E', 'L', 'L', 'O'}

	s := *(*string)(unsafe.Pointer(&b))

	fmt.Println("b =", b)
	fmt.Println("s =", s)

	b[1] = 'B'
	fmt.Println("s =", s)

	s = "WORLD"
	fmt.Println("b =", b)
	fmt.Println("s =", s)

	//b = [72 69 76 76 79]
	//s = HELLO
	//s = HBLLO
	//b = [72 66 76 76 79]
	//s = WORLD
}

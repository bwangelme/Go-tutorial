package main

import (
	"io/ioutil"
	"log"
	"tutorial/bmath"
)

func main() {
	data := []byte{0x61, 0x62}
	bmath.Sum(1, 2)

	// 写如到空文件名的时候会报错
	// 2019/04/26 07:37:42 open : no such file or directory
	if err := ioutil.WriteFile("/tmp/abc", data, 0644); err != nil {
		log.Fatal(err)
	}
}

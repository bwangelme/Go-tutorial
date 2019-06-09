package main

import "log"

func MyRecover() {
	if r := recover(); r != nil {
		log.Println(r)
	}
}

func main() {
	// recover 函数捕获的是祖父一级的函数栈帧的异常信息
	// 例如下面的例子中，捕获的是main函数栈帧中的异常
	// recover
	// |-- MyRecover
	// |-- main
	defer MyRecover() // 可以捕获到异常
	//defer recover() // 无法捕获到异常
	panic(1)
}

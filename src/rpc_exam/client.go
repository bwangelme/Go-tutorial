package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type HelloServiceClient struct {
	c *rpc.Client
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.c.Call(HelloServiceName+".Hello", request, reply)
}

// 这条语句是保护语句，确保 HelloServiceClient 实现了 HelloServiceInterface 接口
var _ HelloServiceInterface = (*HelloServiceClient)(nil) // 等号右边就是普通的类型转换

func DialHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	c := &HelloServiceClient{c: client}
	return c, nil
}

func main() {
	client, err := DialHelloServiceClient("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var arg string
	if len(os.Args) != 2 {
		arg = "xff"
	} else {
		arg = os.Args[1]
	}

	var reply string
	err = client.Hello(arg, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

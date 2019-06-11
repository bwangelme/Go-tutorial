package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func HelloServiceRegist(svc HelloServiceInterface) {
	rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	HelloServiceRegist(&HelloService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go rpc.ServeConn(conn)
	}

}

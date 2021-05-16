package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aiaoyang/directorystructrue/api/say"
	"github.com/aiaoyang/directorystructrue/configs"
	"github.com/aiaoyang/directorystructrue/internal/client"
)

func main() {
	addr := fmt.Sprintf("%s:%d", configs.GCONF.Addr, configs.GCONF.Port)
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	msg := &say.HelloRequest{
		Name: "dong",
	}
	client.SendMsg(c, msg)
}

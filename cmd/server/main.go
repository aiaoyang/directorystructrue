package main

import (
	"fmt"
	"log"
	"net"

	"github.com/aiaoyang/directorystructrue/configs"
	"github.com/aiaoyang/directorystructrue/internal/server"
)

func main() {
	addr := fmt.Sprintf("%s:%d", configs.GCONF.Addr, configs.GCONF.Port)
	fmt.Printf("server started at %s\n", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		err = server.Handle(c)
		if err != nil {
			log.Printf("got err : %s\n", err)
		}
	}
}

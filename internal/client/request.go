package client

import (
	"fmt"
	"net"

	"github.com/aiaoyang/directorystructrue/api/say"
	"google.golang.org/protobuf/proto"
)

func SendMsg(c net.Conn, msg *say.HelloRequest) (err error) {
	defer c.Close()
	b, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	n, err := c.Write(b)
	if err != nil {
		return
	}
	fmt.Printf("send %d bytes\n", n)
	return
}

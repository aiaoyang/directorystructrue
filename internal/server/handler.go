package server

import (
	"fmt"
	"io"
	"net"

	"github.com/aiaoyang/directorystructrue/api/say"
	"google.golang.org/protobuf/proto"
)

func Handle(c net.Conn) (err error) {
	res, err := getConnContent(c)
	if err != nil {
		return
	}
	fmt.Printf("got %v\n", res)
	return
}

func getConnContent(c net.Conn) (content *say.HelloRequest, err error) {
	content = &say.HelloRequest{}
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil && err != io.EOF {
		return
	}
	err = proto.Unmarshal(buf[:n], content)
	return

}

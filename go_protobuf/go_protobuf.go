package main

import (
	"fmt"
	stdProto "goDemo/go_protobuf/proto"
	"net"

	"github.com/golang/protobuf/proto"
)

func main() {
	var p stdProto.Person
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)

	}
	fmt.Println(p)
}

func handleConn(conn net.Conn) {
	buffer := make([]byte, 256)
	person := &stdProto.Person{}
	cnt, _ := conn.Read(buffer)
	proto.Unmarshal(buffer[:cnt], person)
	fmt.Println(person.Name)
}

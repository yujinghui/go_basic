package main 

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("start to send")
	tcpAddr , _:= net.ResolveTCPAddr("tcp4", ":8887")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	_, err := conn.Write([]byte("hello world\n"))
	conn.Close()
	fmt.Println(err)
}
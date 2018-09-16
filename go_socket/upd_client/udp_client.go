package main 

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("start to send")
	udpAddr , _:= net.ResolveUDPAddr("udp4", ":8887")
	conn, _ := net.DialUDP("udp", nil, udpAddr)
	_, err := conn.Write([]byte("hello world\n"))
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	fmt.Println(string(buf[:]))
	fmt.Println(n)
	
	conn.Close()
	fmt.Println(err)
}
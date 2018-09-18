package main

import (
	"net"
	"fmt"
	"time"
)
func main(){
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8887")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil  {
			fmt.Println(err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	buffer := make([]byte, 512)
	defer conn.Close()
	read_line, err := conn.Read(buffer)
	fmt.Printf("read line %v" , read_line)
	if(err != nil) {
		return
	}

	if read_line == 0 {
		return 
	}
	fmt.Println(buffer)
	fmt.Println(string(buffer[:]))
}

func checkError(err error) {
	fmt.Println("------> err", err)
}
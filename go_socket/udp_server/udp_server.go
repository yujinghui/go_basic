
package main

import (
	"net"
	"fmt"
	"time"
)
func main(){
	udpAddr, err := net.ResolveUDPAddr("udp4", ":8887")
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for{
	    handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return 
	}
	conn.WriteToUDP([]byte(time.Now().String()), addr)
	fmt.Println(string(buf[:]))
}

func checkError(err error) {
	fmt.Println("------> err", err)
}
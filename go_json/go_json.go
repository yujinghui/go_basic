package main

import (
	"fmt"
	"encoding/json"
)


type Server struct {
	ServerName string `json:"serverName"`
	ServerIP string 
	// serverName string // 小写开头的不会解析
}

type ServerSlice struct {
	Servers []Server
}
func main(){
	var s ServerSlice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	byte1, _ := json.Marshal(s)
	fmt.Println(string(byte1))
}
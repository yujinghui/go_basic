package main


import(
	"github.com/golang/protobuf/proto"
	"fmt"
	stdProto "goDemo/go_protobuf/proto"
	"net"
	"time"
)

func main() {
	var conn net.Conn
	var err error
	for conn, err = net.Dial("tcp", ":8888"); err != nil; conn, err = net.Dial("tcp", ":8888") {
		fmt.Println("tcp dial fail and ready to retry")
		time.Sleep(time.Second)
		fmt.Println("dialing...")
	}
	defer conn.Close()
	p := &stdProto.Person{
		Name: "yujinghui",
		Age:11,
		Email:"jinghui@111.com",
	}
	personData, _ := proto.Marshal(p)
	conn.Write(personData)
}

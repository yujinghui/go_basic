
package main

import (
	"fmt"
)

func main(){
	str := "helloworld"
	fmt.Println(len(str))
	map1 := make(map[string][]string, 10)
	map1["hellow"] = []string{"world"}
	fmt.Println(map1)
}